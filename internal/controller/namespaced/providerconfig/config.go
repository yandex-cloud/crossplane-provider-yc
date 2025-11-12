/*
Copyright 2021 The Crossplane Authors.

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

package providerconfig

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/crossplane-runtime/v2/pkg/event"
	"github.com/crossplane/crossplane-runtime/v2/pkg/reconciler/providerconfig"
	"github.com/crossplane/crossplane-runtime/v2/pkg/resource"
	"github.com/crossplane/upjet/v2/pkg/controller"

	"github.com/yandex-cloud/crossplane-provider-yc/apis/namespaced/v1beta1"
)

// Setup adds controllers that reconcile both ProviderConfig and ClusterProviderConfig
// by accounting for their current usage.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	// Setup for namespaced ProviderConfig
	pcName := providerconfig.ControllerName(v1beta1.ProviderConfigGroupKind)
	pcOf := resource.ProviderConfigKinds{
		Config:    v1beta1.ProviderConfigGroupVersionKind,
		Usage:     v1beta1.ProviderConfigUsageGroupVersionKind,
		UsageList: v1beta1.ProviderConfigUsageListGroupVersionKind,
	}

	if err := ctrl.NewControllerManagedBy(mgr).
		Named(pcName).
		WithOptions(o.ForControllerRuntime()).
		For(&v1beta1.ProviderConfig{}).
		Watches(&v1beta1.ProviderConfigUsage{}, &resource.EnqueueRequestForProviderConfig{}).
		Complete(providerconfig.NewReconciler(mgr, pcOf,
			providerconfig.WithLogger(o.Logger.WithValues("controller", pcName)),
			providerconfig.WithRecorder(event.NewAPIRecorder(mgr.GetEventRecorderFor(pcName))))); err != nil {
		return err
	}

	// Setup for ClusterProviderConfig
	cpcName := providerconfig.ControllerName(v1beta1.ClusterProviderConfigGroupKind)
	cpcOf := resource.ProviderConfigKinds{
		Config:    v1beta1.ClusterProviderConfigGroupVersionKind,
		Usage:     v1beta1.ProviderConfigUsageGroupVersionKind,
		UsageList: v1beta1.ProviderConfigUsageListGroupVersionKind,
	}

	return ctrl.NewControllerManagedBy(mgr).
		Named(cpcName).
		WithOptions(o.ForControllerRuntime()).
		For(&v1beta1.ClusterProviderConfig{}).
		Watches(&v1beta1.ProviderConfigUsage{}, &resource.EnqueueRequestForProviderConfig{}).
		Complete(providerconfig.NewReconciler(mgr, cpcOf,
			providerconfig.WithLogger(o.Logger.WithValues("controller", cpcName)),
			providerconfig.WithRecorder(event.NewAPIRecorder(mgr.GetEventRecorderFor(cpcName)))))
}

// SetupGated adds controllers with gated startup that reconcile both ProviderConfig
// and ClusterProviderConfig
func SetupGated(mgr ctrl.Manager, o controller.Options) error {
	o.Gate.Register(func() {
		if err := Setup(mgr, o); err != nil {
			mgr.GetLogger().Error(err, "unable to setup reconcilers")
		}
	}, v1beta1.ClusterProviderConfigGroupVersionKind,
		v1beta1.ProviderConfigGroupVersionKind,
		v1beta1.ProviderConfigUsageGroupVersionKind)
	return nil
}
