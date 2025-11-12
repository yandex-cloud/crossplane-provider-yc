/*
Copyright 2022 YANDEX LLC
This is modified version of the software, made by the Crossplane Authors
and available at: https://github.com/crossplane-contrib/provider-jet-template

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"context"
	"os"
	"path/filepath"
	"time"

	authv1 "k8s.io/api/authorization/v1"
	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/tools/leaderelection/resourcelock"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/cache"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"
	"sigs.k8s.io/controller-runtime/pkg/manager"

	xpcontroller "github.com/crossplane/crossplane-runtime/v2/pkg/controller"
	"github.com/crossplane/crossplane-runtime/v2/pkg/errors"
	"github.com/crossplane/crossplane-runtime/v2/pkg/feature"
	"github.com/crossplane/crossplane-runtime/v2/pkg/gate"
	"github.com/crossplane/crossplane-runtime/v2/pkg/logging"
	"github.com/crossplane/crossplane-runtime/v2/pkg/ratelimiter"
	"github.com/crossplane/crossplane-runtime/v2/pkg/reconciler/customresourcesgate"
	tjcontroller "github.com/crossplane/upjet/v2/pkg/controller"
	"github.com/crossplane/upjet/v2/pkg/terraform"
	"gopkg.in/alecthomas/kingpin.v2"

	apisCluster "github.com/yandex-cloud/crossplane-provider-yc/apis/cluster"
	apisNamespaced "github.com/yandex-cloud/crossplane-provider-yc/apis/namespaced"
	"github.com/yandex-cloud/crossplane-provider-yc/config"
	"github.com/yandex-cloud/crossplane-provider-yc/internal/clients"
	controllerCluster "github.com/yandex-cloud/crossplane-provider-yc/internal/controller/cluster"
	controllerNamespaced "github.com/yandex-cloud/crossplane-provider-yc/internal/controller/namespaced"
	"github.com/yandex-cloud/crossplane-provider-yc/internal/features"
)

func main() {
	var (
		app              = kingpin.New(filepath.Base(os.Args[0]), "YC support for Crossplane.").DefaultEnvars()
		debug            = app.Flag("debug", "Run with debug logging.").Short('d').Bool()
		syncPeriod       = app.Flag("sync", "Controller manager sync period such as 300ms, 1.5h, or 2h45m").Short('s').Default("1h").Duration()
		pollInterval     = app.Flag("poll", "Poll interval controls how often an individual resource should be checked for drift.").Default("10m").Duration()
		leaderElection   = app.Flag("leader-election", "Use leader election for the controller manager.").Short('l').Default("false").OverrideDefaultFromEnvar("LEADER_ELECTION").Bool()
		maxReconcileRate = app.Flag("max-reconcile-rate", "The global maximum rate per second at which resources may be checked for drift from the desired state.").Default("10").Int()

		terraformVersion = app.Flag("terraform-version", "Terraform version.").Required().Envar("TERRAFORM_VERSION").String()
		providerSource   = app.Flag("terraform-provider-source", "Terraform provider source.").Required().Envar("TERRAFORM_PROVIDER_SOURCE").String()
		providerVersion  = app.Flag("terraform-provider-version", "Terraform provider version.").Required().Envar("TERRAFORM_PROVIDER_VERSION").String()

		enableManagementPolicies = app.Flag("enable-management-policies", "Enable support for Management Policies.").Default("true").Envar("ENABLE_MANAGEMENT_POLICIES").Bool()
	)
	kingpin.MustParse(app.Parse(os.Args[1:]))

	zl := zap.New(zap.UseDevMode(*debug))
	log := logging.NewLogrLogger(zl.WithName("crossplane-provider-yc"))
	if *debug {
		// The controller-runtime runs with a no-op logger by default. It is
		// *very* verbose even at info level, so we only provide it a real
		// logger when we're running in debug mode.
		ctrl.SetLogger(zl)
	}

	log.Debug("Starting", "sync-period", syncPeriod.String(), "poll-interval", pollInterval.String(), "max-reconcile-rate", *maxReconcileRate)

	cfg, err := ctrl.GetConfig()
	kingpin.FatalIfError(err, "Cannot get API server rest config")

	mgr, err := ctrl.NewManager(cfg, ctrl.Options{
		LeaderElection:   *leaderElection,
		LeaderElectionID: "crossplane-leader-election-upjet-provider-template",
		Cache: cache.Options{
			SyncPeriod: syncPeriod,
		},
		LeaderElectionResourceLock: resourcelock.LeasesResourceLock,
		LeaseDuration:              func() *time.Duration { d := 60 * time.Second; return &d }(),
		RenewDeadline:              func() *time.Duration { d := 50 * time.Second; return &d }(),
	})
	kingpin.FatalIfError(err, "Cannot create controller manager")

	// Add both cluster and namespaced APIs to scheme
	kingpin.FatalIfError(apisCluster.AddToScheme(mgr.GetScheme()), "Cannot add cluster APIs to scheme")
	kingpin.FatalIfError(apisNamespaced.AddToScheme(mgr.GetScheme()), "Cannot add namespaced APIs to scheme")
	kingpin.FatalIfError(apiextensionsv1.AddToScheme(mgr.GetScheme()), "Cannot add api-extensions APIs to scheme")
	kingpin.FatalIfError(authv1.AddToScheme(mgr.GetScheme()), "Cannot add k8s authorization APIs to scheme")

	// Get both cluster and namespaced providers
	clusterProvider := config.GetProvider()
	namespacedProvider := config.GetProviderNamespaced()

	// Create controller options for cluster-scoped resources
	clusterOpts := tjcontroller.Options{
		Options: xpcontroller.Options{
			Logger:                  log,
			GlobalRateLimiter:       ratelimiter.NewGlobal(*maxReconcileRate),
			PollInterval:            *pollInterval,
			MaxConcurrentReconciles: *maxReconcileRate,
			Features:                &feature.Flags{},
		},
		Provider:              clusterProvider,
		WorkspaceStore:        terraform.NewWorkspaceStore(log),
		SetupFn:               clients.TerraformSetupBuilder(*terraformVersion, *providerSource, *providerVersion, clusterProvider),
		OperationTrackerStore: tjcontroller.NewOperationStore(log),
	}

	// Create controller options for namespaced resources
	namespacedOpts := tjcontroller.Options{
		Options: xpcontroller.Options{
			Logger:                  log,
			GlobalRateLimiter:       ratelimiter.NewGlobal(*maxReconcileRate),
			PollInterval:            *pollInterval,
			MaxConcurrentReconciles: *maxReconcileRate,
			Features:                &feature.Flags{},
		},
		Provider:              namespacedProvider,
		WorkspaceStore:        terraform.NewWorkspaceStore(log),
		SetupFn:               clients.TerraformSetupBuilder(*terraformVersion, *providerSource, *providerVersion, namespacedProvider),
		OperationTrackerStore: tjcontroller.NewOperationStore(log),
	}

	// Apply feature flags to both controller options
	if *enableManagementPolicies {
		clusterOpts.Features.Enable(features.EnableBetaManagementPolicies)
		namespacedOpts.Features.Enable(features.EnableBetaManagementPolicies)
		log.Info("Beta feature enabled", "flag", features.EnableBetaManagementPolicies)
	}

	// Check if SafeStart capability is available
	canSafeStart, err := canWatchCRD(context.TODO(), mgr)
	kingpin.FatalIfError(err, "SafeStart precheck failed")

	if canSafeStart {
		log.Info("SafeStart capability enabled - controllers will start after CRDs are available")

		// Create CRD gate for SafeStart
		crdGate := new(gate.Gate[schema.GroupVersionKind])
		clusterOpts.Gate = crdGate
		namespacedOpts.Gate = crdGate

		// Setup CRD gate controller
		kingpin.FatalIfError(customresourcesgate.Setup(mgr, xpcontroller.Options{
			Logger:                  log,
			Gate:                    crdGate,
			MaxConcurrentReconciles: 1,
		}), "Cannot setup CRD gate")

		// Setup gated controllers
		kingpin.FatalIfError(controllerCluster.SetupGated(mgr, clusterOpts),
			"Cannot setup cluster-scoped controllers")
		kingpin.FatalIfError(controllerNamespaced.SetupGated(mgr, namespacedOpts),
			"Cannot setup namespaced controllers")
	} else {
		log.Info("Provider has missing RBAC permissions for watching CRDs, controller SafeStart capability will be disabled")

		// Setup regular controllers without SafeStart
		kingpin.FatalIfError(controllerCluster.Setup(mgr, clusterOpts),
			"Cannot setup cluster-scoped controllers")
		kingpin.FatalIfError(controllerNamespaced.Setup(mgr, namespacedOpts),
			"Cannot setup namespaced controllers")
	}

	kingpin.FatalIfError(mgr.Start(ctrl.SetupSignalHandler()), "Cannot start controller manager")
}

// canWatchCRD checks if the provider has RBAC permissions to watch CRDs
// This is required for SafeStart capability
func canWatchCRD(ctx context.Context, mgr manager.Manager) (bool, error) {
	verbs := []string{"get", "list", "watch"}
	for _, verb := range verbs {
		sar := &authv1.SelfSubjectAccessReview{
			Spec: authv1.SelfSubjectAccessReviewSpec{
				ResourceAttributes: &authv1.ResourceAttributes{
					Group:    "apiextensions.k8s.io",
					Resource: "customresourcedefinitions",
					Verb:     verb,
				},
			},
		}
		if err := mgr.GetClient().Create(ctx, sar); err != nil {
			return false, errors.Wrapf(err, "unable to perform RBAC check for verb %s on CustomResourceDefinitions", verb)
		}
		if !sar.Status.Allowed {
			return false, nil
		}
	}
	return true, nil
}
