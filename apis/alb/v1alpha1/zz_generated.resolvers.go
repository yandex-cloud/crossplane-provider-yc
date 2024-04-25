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
// Code generated by angryjet. DO NOT EDIT.

package v1alpha1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	errors "github.com/pkg/errors"
	v1alpha1 "github.com/yandex-cloud/provider-jet-yc/apis/resourcemanager/v1alpha1"
	v1alpha11 "github.com/yandex-cloud/provider-jet-yc/apis/vpc/v1alpha1"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this BackendGroup.
func (mg *BackendGroup) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var mrsp reference.MultiResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.FolderID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.FolderIDRef,
		Selector:     mg.Spec.ForProvider.FolderIDSelector,
		To: reference.To{
			List:    &v1alpha1.FolderList{},
			Managed: &v1alpha1.Folder{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.FolderID")
	}
	mg.Spec.ForProvider.FolderID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.FolderIDRef = rsp.ResolvedReference

	for i3 := 0; i3 < len(mg.Spec.ForProvider.GRPCBackend); i3++ {
		mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
			CurrentValues: reference.FromPtrValues(mg.Spec.ForProvider.GRPCBackend[i3].TargetGroupIds),
			Extract:       reference.ExternalName(),
			References:    mg.Spec.ForProvider.GRPCBackend[i3].TargetGroupIdsRefs,
			Selector:      mg.Spec.ForProvider.GRPCBackend[i3].TargetGroupIdsSelector,
			To: reference.To{
				List:    &TargetGroupList{},
				Managed: &TargetGroup{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.GRPCBackend[i3].TargetGroupIds")
		}
		mg.Spec.ForProvider.GRPCBackend[i3].TargetGroupIds = reference.ToPtrValues(mrsp.ResolvedValues)
		mg.Spec.ForProvider.GRPCBackend[i3].TargetGroupIdsRefs = mrsp.ResolvedReferences

	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.HTTPBackend); i3++ {
		mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
			CurrentValues: reference.FromPtrValues(mg.Spec.ForProvider.HTTPBackend[i3].TargetGroupIds),
			Extract:       reference.ExternalName(),
			References:    mg.Spec.ForProvider.HTTPBackend[i3].TargetGroupIdsRefs,
			Selector:      mg.Spec.ForProvider.HTTPBackend[i3].TargetGroupIdsSelector,
			To: reference.To{
				List:    &TargetGroupList{},
				Managed: &TargetGroup{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.HTTPBackend[i3].TargetGroupIds")
		}
		mg.Spec.ForProvider.HTTPBackend[i3].TargetGroupIds = reference.ToPtrValues(mrsp.ResolvedValues)
		mg.Spec.ForProvider.HTTPBackend[i3].TargetGroupIdsRefs = mrsp.ResolvedReferences

	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.StreamBackend); i3++ {
		mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
			CurrentValues: reference.FromPtrValues(mg.Spec.ForProvider.StreamBackend[i3].TargetGroupIds),
			Extract:       reference.ExternalName(),
			References:    mg.Spec.ForProvider.StreamBackend[i3].TargetGroupIdsRefs,
			Selector:      mg.Spec.ForProvider.StreamBackend[i3].TargetGroupIdsSelector,
			To: reference.To{
				List:    &TargetGroupList{},
				Managed: &TargetGroup{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.StreamBackend[i3].TargetGroupIds")
		}
		mg.Spec.ForProvider.StreamBackend[i3].TargetGroupIds = reference.ToPtrValues(mrsp.ResolvedValues)
		mg.Spec.ForProvider.StreamBackend[i3].TargetGroupIdsRefs = mrsp.ResolvedReferences

	}
	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.FolderID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.FolderIDRef,
		Selector:     mg.Spec.InitProvider.FolderIDSelector,
		To: reference.To{
			List:    &v1alpha1.FolderList{},
			Managed: &v1alpha1.Folder{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.FolderID")
	}
	mg.Spec.InitProvider.FolderID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.FolderIDRef = rsp.ResolvedReference

	for i3 := 0; i3 < len(mg.Spec.InitProvider.GRPCBackend); i3++ {
		mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
			CurrentValues: reference.FromPtrValues(mg.Spec.InitProvider.GRPCBackend[i3].TargetGroupIds),
			Extract:       reference.ExternalName(),
			References:    mg.Spec.InitProvider.GRPCBackend[i3].TargetGroupIdsRefs,
			Selector:      mg.Spec.InitProvider.GRPCBackend[i3].TargetGroupIdsSelector,
			To: reference.To{
				List:    &TargetGroupList{},
				Managed: &TargetGroup{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.InitProvider.GRPCBackend[i3].TargetGroupIds")
		}
		mg.Spec.InitProvider.GRPCBackend[i3].TargetGroupIds = reference.ToPtrValues(mrsp.ResolvedValues)
		mg.Spec.InitProvider.GRPCBackend[i3].TargetGroupIdsRefs = mrsp.ResolvedReferences

	}
	for i3 := 0; i3 < len(mg.Spec.InitProvider.HTTPBackend); i3++ {
		mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
			CurrentValues: reference.FromPtrValues(mg.Spec.InitProvider.HTTPBackend[i3].TargetGroupIds),
			Extract:       reference.ExternalName(),
			References:    mg.Spec.InitProvider.HTTPBackend[i3].TargetGroupIdsRefs,
			Selector:      mg.Spec.InitProvider.HTTPBackend[i3].TargetGroupIdsSelector,
			To: reference.To{
				List:    &TargetGroupList{},
				Managed: &TargetGroup{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.InitProvider.HTTPBackend[i3].TargetGroupIds")
		}
		mg.Spec.InitProvider.HTTPBackend[i3].TargetGroupIds = reference.ToPtrValues(mrsp.ResolvedValues)
		mg.Spec.InitProvider.HTTPBackend[i3].TargetGroupIdsRefs = mrsp.ResolvedReferences

	}
	for i3 := 0; i3 < len(mg.Spec.InitProvider.StreamBackend); i3++ {
		mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
			CurrentValues: reference.FromPtrValues(mg.Spec.InitProvider.StreamBackend[i3].TargetGroupIds),
			Extract:       reference.ExternalName(),
			References:    mg.Spec.InitProvider.StreamBackend[i3].TargetGroupIdsRefs,
			Selector:      mg.Spec.InitProvider.StreamBackend[i3].TargetGroupIdsSelector,
			To: reference.To{
				List:    &TargetGroupList{},
				Managed: &TargetGroup{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.InitProvider.StreamBackend[i3].TargetGroupIds")
		}
		mg.Spec.InitProvider.StreamBackend[i3].TargetGroupIds = reference.ToPtrValues(mrsp.ResolvedValues)
		mg.Spec.InitProvider.StreamBackend[i3].TargetGroupIdsRefs = mrsp.ResolvedReferences

	}

	return nil
}

// ResolveReferences of this HTTPRouter.
func (mg *HTTPRouter) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.FolderID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.FolderIDRef,
		Selector:     mg.Spec.ForProvider.FolderIDSelector,
		To: reference.To{
			List:    &v1alpha1.FolderList{},
			Managed: &v1alpha1.Folder{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.FolderID")
	}
	mg.Spec.ForProvider.FolderID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.FolderIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.FolderID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.FolderIDRef,
		Selector:     mg.Spec.InitProvider.FolderIDSelector,
		To: reference.To{
			List:    &v1alpha1.FolderList{},
			Managed: &v1alpha1.Folder{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.FolderID")
	}
	mg.Spec.InitProvider.FolderID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.FolderIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this LoadBalancer.
func (mg *LoadBalancer) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var mrsp reference.MultiResolutionResponse
	var err error

	for i3 := 0; i3 < len(mg.Spec.ForProvider.AllocationPolicy); i3++ {
		for i4 := 0; i4 < len(mg.Spec.ForProvider.AllocationPolicy[i3].Location); i4++ {
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.AllocationPolicy[i3].Location[i4].SubnetID),
				Extract:      reference.ExternalName(),
				Reference:    mg.Spec.ForProvider.AllocationPolicy[i3].Location[i4].SubnetIDRef,
				Selector:     mg.Spec.ForProvider.AllocationPolicy[i3].Location[i4].SubnetIDSelector,
				To: reference.To{
					List:    &v1alpha11.SubnetList{},
					Managed: &v1alpha11.Subnet{},
				},
			})
			if err != nil {
				return errors.Wrap(err, "mg.Spec.ForProvider.AllocationPolicy[i3].Location[i4].SubnetID")
			}
			mg.Spec.ForProvider.AllocationPolicy[i3].Location[i4].SubnetID = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.ForProvider.AllocationPolicy[i3].Location[i4].SubnetIDRef = rsp.ResolvedReference

		}
	}
	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.FolderID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.FolderIDRef,
		Selector:     mg.Spec.ForProvider.FolderIDSelector,
		To: reference.To{
			List:    &v1alpha1.FolderList{},
			Managed: &v1alpha1.Folder{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.FolderID")
	}
	mg.Spec.ForProvider.FolderID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.FolderIDRef = rsp.ResolvedReference

	for i3 := 0; i3 < len(mg.Spec.ForProvider.Listener); i3++ {
		for i4 := 0; i4 < len(mg.Spec.ForProvider.Listener[i3].Endpoint); i4++ {
			for i5 := 0; i5 < len(mg.Spec.ForProvider.Listener[i3].Endpoint[i4].Address); i5++ {
				for i6 := 0; i6 < len(mg.Spec.ForProvider.Listener[i3].Endpoint[i4].Address[i5].InternalIPv4Address); i6++ {
					rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
						CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Listener[i3].Endpoint[i4].Address[i5].InternalIPv4Address[i6].SubnetID),
						Extract:      reference.ExternalName(),
						Reference:    mg.Spec.ForProvider.Listener[i3].Endpoint[i4].Address[i5].InternalIPv4Address[i6].SubnetIDRef,
						Selector:     mg.Spec.ForProvider.Listener[i3].Endpoint[i4].Address[i5].InternalIPv4Address[i6].SubnetIDSelector,
						To: reference.To{
							List:    &v1alpha11.SubnetList{},
							Managed: &v1alpha11.Subnet{},
						},
					})
					if err != nil {
						return errors.Wrap(err, "mg.Spec.ForProvider.Listener[i3].Endpoint[i4].Address[i5].InternalIPv4Address[i6].SubnetID")
					}
					mg.Spec.ForProvider.Listener[i3].Endpoint[i4].Address[i5].InternalIPv4Address[i6].SubnetID = reference.ToPtrValue(rsp.ResolvedValue)
					mg.Spec.ForProvider.Listener[i3].Endpoint[i4].Address[i5].InternalIPv4Address[i6].SubnetIDRef = rsp.ResolvedReference

				}
			}
		}
	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.Listener); i3++ {
		for i4 := 0; i4 < len(mg.Spec.ForProvider.Listener[i3].HTTP); i4++ {
			for i5 := 0; i5 < len(mg.Spec.ForProvider.Listener[i3].HTTP[i4].Handler); i5++ {
				rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
					CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Listener[i3].HTTP[i4].Handler[i5].HTTPRouterID),
					Extract:      reference.ExternalName(),
					Reference:    mg.Spec.ForProvider.Listener[i3].HTTP[i4].Handler[i5].HTTPRouterIDRef,
					Selector:     mg.Spec.ForProvider.Listener[i3].HTTP[i4].Handler[i5].HTTPRouterIDSelector,
					To: reference.To{
						List:    &HTTPRouterList{},
						Managed: &HTTPRouter{},
					},
				})
				if err != nil {
					return errors.Wrap(err, "mg.Spec.ForProvider.Listener[i3].HTTP[i4].Handler[i5].HTTPRouterID")
				}
				mg.Spec.ForProvider.Listener[i3].HTTP[i4].Handler[i5].HTTPRouterID = reference.ToPtrValue(rsp.ResolvedValue)
				mg.Spec.ForProvider.Listener[i3].HTTP[i4].Handler[i5].HTTPRouterIDRef = rsp.ResolvedReference

			}
		}
	}
	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.NetworkID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.NetworkIDRef,
		Selector:     mg.Spec.ForProvider.NetworkIDSelector,
		To: reference.To{
			List:    &v1alpha11.NetworkList{},
			Managed: &v1alpha11.Network{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.NetworkID")
	}
	mg.Spec.ForProvider.NetworkID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.NetworkIDRef = rsp.ResolvedReference

	mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
		CurrentValues: reference.FromPtrValues(mg.Spec.ForProvider.SecurityGroupIds),
		Extract:       reference.ExternalName(),
		References:    mg.Spec.ForProvider.SecurityGroupIdsRefs,
		Selector:      mg.Spec.ForProvider.SecurityGroupIdsSelector,
		To: reference.To{
			List:    &v1alpha11.SecurityGroupList{},
			Managed: &v1alpha11.SecurityGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.SecurityGroupIds")
	}
	mg.Spec.ForProvider.SecurityGroupIds = reference.ToPtrValues(mrsp.ResolvedValues)
	mg.Spec.ForProvider.SecurityGroupIdsRefs = mrsp.ResolvedReferences

	for i3 := 0; i3 < len(mg.Spec.InitProvider.AllocationPolicy); i3++ {
		for i4 := 0; i4 < len(mg.Spec.InitProvider.AllocationPolicy[i3].Location); i4++ {
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.AllocationPolicy[i3].Location[i4].SubnetID),
				Extract:      reference.ExternalName(),
				Reference:    mg.Spec.InitProvider.AllocationPolicy[i3].Location[i4].SubnetIDRef,
				Selector:     mg.Spec.InitProvider.AllocationPolicy[i3].Location[i4].SubnetIDSelector,
				To: reference.To{
					List:    &v1alpha11.SubnetList{},
					Managed: &v1alpha11.Subnet{},
				},
			})
			if err != nil {
				return errors.Wrap(err, "mg.Spec.InitProvider.AllocationPolicy[i3].Location[i4].SubnetID")
			}
			mg.Spec.InitProvider.AllocationPolicy[i3].Location[i4].SubnetID = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.InitProvider.AllocationPolicy[i3].Location[i4].SubnetIDRef = rsp.ResolvedReference

		}
	}
	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.FolderID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.FolderIDRef,
		Selector:     mg.Spec.InitProvider.FolderIDSelector,
		To: reference.To{
			List:    &v1alpha1.FolderList{},
			Managed: &v1alpha1.Folder{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.FolderID")
	}
	mg.Spec.InitProvider.FolderID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.FolderIDRef = rsp.ResolvedReference

	for i3 := 0; i3 < len(mg.Spec.InitProvider.Listener); i3++ {
		for i4 := 0; i4 < len(mg.Spec.InitProvider.Listener[i3].Endpoint); i4++ {
			for i5 := 0; i5 < len(mg.Spec.InitProvider.Listener[i3].Endpoint[i4].Address); i5++ {
				for i6 := 0; i6 < len(mg.Spec.InitProvider.Listener[i3].Endpoint[i4].Address[i5].InternalIPv4Address); i6++ {
					rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
						CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Listener[i3].Endpoint[i4].Address[i5].InternalIPv4Address[i6].SubnetID),
						Extract:      reference.ExternalName(),
						Reference:    mg.Spec.InitProvider.Listener[i3].Endpoint[i4].Address[i5].InternalIPv4Address[i6].SubnetIDRef,
						Selector:     mg.Spec.InitProvider.Listener[i3].Endpoint[i4].Address[i5].InternalIPv4Address[i6].SubnetIDSelector,
						To: reference.To{
							List:    &v1alpha11.SubnetList{},
							Managed: &v1alpha11.Subnet{},
						},
					})
					if err != nil {
						return errors.Wrap(err, "mg.Spec.InitProvider.Listener[i3].Endpoint[i4].Address[i5].InternalIPv4Address[i6].SubnetID")
					}
					mg.Spec.InitProvider.Listener[i3].Endpoint[i4].Address[i5].InternalIPv4Address[i6].SubnetID = reference.ToPtrValue(rsp.ResolvedValue)
					mg.Spec.InitProvider.Listener[i3].Endpoint[i4].Address[i5].InternalIPv4Address[i6].SubnetIDRef = rsp.ResolvedReference

				}
			}
		}
	}
	for i3 := 0; i3 < len(mg.Spec.InitProvider.Listener); i3++ {
		for i4 := 0; i4 < len(mg.Spec.InitProvider.Listener[i3].HTTP); i4++ {
			for i5 := 0; i5 < len(mg.Spec.InitProvider.Listener[i3].HTTP[i4].Handler); i5++ {
				rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
					CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Listener[i3].HTTP[i4].Handler[i5].HTTPRouterID),
					Extract:      reference.ExternalName(),
					Reference:    mg.Spec.InitProvider.Listener[i3].HTTP[i4].Handler[i5].HTTPRouterIDRef,
					Selector:     mg.Spec.InitProvider.Listener[i3].HTTP[i4].Handler[i5].HTTPRouterIDSelector,
					To: reference.To{
						List:    &HTTPRouterList{},
						Managed: &HTTPRouter{},
					},
				})
				if err != nil {
					return errors.Wrap(err, "mg.Spec.InitProvider.Listener[i3].HTTP[i4].Handler[i5].HTTPRouterID")
				}
				mg.Spec.InitProvider.Listener[i3].HTTP[i4].Handler[i5].HTTPRouterID = reference.ToPtrValue(rsp.ResolvedValue)
				mg.Spec.InitProvider.Listener[i3].HTTP[i4].Handler[i5].HTTPRouterIDRef = rsp.ResolvedReference

			}
		}
	}
	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.NetworkID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.NetworkIDRef,
		Selector:     mg.Spec.InitProvider.NetworkIDSelector,
		To: reference.To{
			List:    &v1alpha11.NetworkList{},
			Managed: &v1alpha11.Network{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.NetworkID")
	}
	mg.Spec.InitProvider.NetworkID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.NetworkIDRef = rsp.ResolvedReference

	mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
		CurrentValues: reference.FromPtrValues(mg.Spec.InitProvider.SecurityGroupIds),
		Extract:       reference.ExternalName(),
		References:    mg.Spec.InitProvider.SecurityGroupIdsRefs,
		Selector:      mg.Spec.InitProvider.SecurityGroupIdsSelector,
		To: reference.To{
			List:    &v1alpha11.SecurityGroupList{},
			Managed: &v1alpha11.SecurityGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.SecurityGroupIds")
	}
	mg.Spec.InitProvider.SecurityGroupIds = reference.ToPtrValues(mrsp.ResolvedValues)
	mg.Spec.InitProvider.SecurityGroupIdsRefs = mrsp.ResolvedReferences

	return nil
}

// ResolveReferences of this TargetGroup.
func (mg *TargetGroup) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.FolderID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.FolderIDRef,
		Selector:     mg.Spec.ForProvider.FolderIDSelector,
		To: reference.To{
			List:    &v1alpha1.FolderList{},
			Managed: &v1alpha1.Folder{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.FolderID")
	}
	mg.Spec.ForProvider.FolderID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.FolderIDRef = rsp.ResolvedReference

	for i3 := 0; i3 < len(mg.Spec.ForProvider.Target); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Target[i3].SubnetID),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.Target[i3].SubnetIDRef,
			Selector:     mg.Spec.ForProvider.Target[i3].SubnetIDSelector,
			To: reference.To{
				List:    &v1alpha11.SubnetList{},
				Managed: &v1alpha11.Subnet{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.Target[i3].SubnetID")
		}
		mg.Spec.ForProvider.Target[i3].SubnetID = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.Target[i3].SubnetIDRef = rsp.ResolvedReference

	}
	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.FolderID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.FolderIDRef,
		Selector:     mg.Spec.InitProvider.FolderIDSelector,
		To: reference.To{
			List:    &v1alpha1.FolderList{},
			Managed: &v1alpha1.Folder{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.FolderID")
	}
	mg.Spec.InitProvider.FolderID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.FolderIDRef = rsp.ResolvedReference

	for i3 := 0; i3 < len(mg.Spec.InitProvider.Target); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Target[i3].SubnetID),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.InitProvider.Target[i3].SubnetIDRef,
			Selector:     mg.Spec.InitProvider.Target[i3].SubnetIDSelector,
			To: reference.To{
				List:    &v1alpha11.SubnetList{},
				Managed: &v1alpha11.Subnet{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.InitProvider.Target[i3].SubnetID")
		}
		mg.Spec.InitProvider.Target[i3].SubnetID = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.InitProvider.Target[i3].SubnetIDRef = rsp.ResolvedReference

	}

	return nil
}

// ResolveReferences of this VirtualHost.
func (mg *VirtualHost) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.HTTPRouterID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.HTTPRouterIDRef,
		Selector:     mg.Spec.ForProvider.HTTPRouterIDSelector,
		To: reference.To{
			List:    &HTTPRouterList{},
			Managed: &HTTPRouter{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.HTTPRouterID")
	}
	mg.Spec.ForProvider.HTTPRouterID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.HTTPRouterIDRef = rsp.ResolvedReference

	for i3 := 0; i3 < len(mg.Spec.ForProvider.Route); i3++ {
		for i4 := 0; i4 < len(mg.Spec.ForProvider.Route[i3].GRPCRoute); i4++ {
			for i5 := 0; i5 < len(mg.Spec.ForProvider.Route[i3].GRPCRoute[i4].GRPCRouteAction); i5++ {
				rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
					CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Route[i3].GRPCRoute[i4].GRPCRouteAction[i5].BackendGroupID),
					Extract:      reference.ExternalName(),
					Reference:    mg.Spec.ForProvider.Route[i3].GRPCRoute[i4].GRPCRouteAction[i5].BackendGroupIDRef,
					Selector:     mg.Spec.ForProvider.Route[i3].GRPCRoute[i4].GRPCRouteAction[i5].BackendGroupIDSelector,
					To: reference.To{
						List:    &BackendGroupList{},
						Managed: &BackendGroup{},
					},
				})
				if err != nil {
					return errors.Wrap(err, "mg.Spec.ForProvider.Route[i3].GRPCRoute[i4].GRPCRouteAction[i5].BackendGroupID")
				}
				mg.Spec.ForProvider.Route[i3].GRPCRoute[i4].GRPCRouteAction[i5].BackendGroupID = reference.ToPtrValue(rsp.ResolvedValue)
				mg.Spec.ForProvider.Route[i3].GRPCRoute[i4].GRPCRouteAction[i5].BackendGroupIDRef = rsp.ResolvedReference

			}
		}
	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.Route); i3++ {
		for i4 := 0; i4 < len(mg.Spec.ForProvider.Route[i3].HTTPRoute); i4++ {
			for i5 := 0; i5 < len(mg.Spec.ForProvider.Route[i3].HTTPRoute[i4].HTTPRouteAction); i5++ {
				rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
					CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Route[i3].HTTPRoute[i4].HTTPRouteAction[i5].BackendGroupID),
					Extract:      reference.ExternalName(),
					Reference:    mg.Spec.ForProvider.Route[i3].HTTPRoute[i4].HTTPRouteAction[i5].BackendGroupIDRef,
					Selector:     mg.Spec.ForProvider.Route[i3].HTTPRoute[i4].HTTPRouteAction[i5].BackendGroupIDSelector,
					To: reference.To{
						List:    &BackendGroupList{},
						Managed: &BackendGroup{},
					},
				})
				if err != nil {
					return errors.Wrap(err, "mg.Spec.ForProvider.Route[i3].HTTPRoute[i4].HTTPRouteAction[i5].BackendGroupID")
				}
				mg.Spec.ForProvider.Route[i3].HTTPRoute[i4].HTTPRouteAction[i5].BackendGroupID = reference.ToPtrValue(rsp.ResolvedValue)
				mg.Spec.ForProvider.Route[i3].HTTPRoute[i4].HTTPRouteAction[i5].BackendGroupIDRef = rsp.ResolvedReference

			}
		}
	}
	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.HTTPRouterID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.HTTPRouterIDRef,
		Selector:     mg.Spec.InitProvider.HTTPRouterIDSelector,
		To: reference.To{
			List:    &HTTPRouterList{},
			Managed: &HTTPRouter{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.HTTPRouterID")
	}
	mg.Spec.InitProvider.HTTPRouterID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.HTTPRouterIDRef = rsp.ResolvedReference

	for i3 := 0; i3 < len(mg.Spec.InitProvider.Route); i3++ {
		for i4 := 0; i4 < len(mg.Spec.InitProvider.Route[i3].GRPCRoute); i4++ {
			for i5 := 0; i5 < len(mg.Spec.InitProvider.Route[i3].GRPCRoute[i4].GRPCRouteAction); i5++ {
				rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
					CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Route[i3].GRPCRoute[i4].GRPCRouteAction[i5].BackendGroupID),
					Extract:      reference.ExternalName(),
					Reference:    mg.Spec.InitProvider.Route[i3].GRPCRoute[i4].GRPCRouteAction[i5].BackendGroupIDRef,
					Selector:     mg.Spec.InitProvider.Route[i3].GRPCRoute[i4].GRPCRouteAction[i5].BackendGroupIDSelector,
					To: reference.To{
						List:    &BackendGroupList{},
						Managed: &BackendGroup{},
					},
				})
				if err != nil {
					return errors.Wrap(err, "mg.Spec.InitProvider.Route[i3].GRPCRoute[i4].GRPCRouteAction[i5].BackendGroupID")
				}
				mg.Spec.InitProvider.Route[i3].GRPCRoute[i4].GRPCRouteAction[i5].BackendGroupID = reference.ToPtrValue(rsp.ResolvedValue)
				mg.Spec.InitProvider.Route[i3].GRPCRoute[i4].GRPCRouteAction[i5].BackendGroupIDRef = rsp.ResolvedReference

			}
		}
	}
	for i3 := 0; i3 < len(mg.Spec.InitProvider.Route); i3++ {
		for i4 := 0; i4 < len(mg.Spec.InitProvider.Route[i3].HTTPRoute); i4++ {
			for i5 := 0; i5 < len(mg.Spec.InitProvider.Route[i3].HTTPRoute[i4].HTTPRouteAction); i5++ {
				rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
					CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Route[i3].HTTPRoute[i4].HTTPRouteAction[i5].BackendGroupID),
					Extract:      reference.ExternalName(),
					Reference:    mg.Spec.InitProvider.Route[i3].HTTPRoute[i4].HTTPRouteAction[i5].BackendGroupIDRef,
					Selector:     mg.Spec.InitProvider.Route[i3].HTTPRoute[i4].HTTPRouteAction[i5].BackendGroupIDSelector,
					To: reference.To{
						List:    &BackendGroupList{},
						Managed: &BackendGroup{},
					},
				})
				if err != nil {
					return errors.Wrap(err, "mg.Spec.InitProvider.Route[i3].HTTPRoute[i4].HTTPRouteAction[i5].BackendGroupID")
				}
				mg.Spec.InitProvider.Route[i3].HTTPRoute[i4].HTTPRouteAction[i5].BackendGroupID = reference.ToPtrValue(rsp.ResolvedValue)
				mg.Spec.InitProvider.Route[i3].HTTPRoute[i4].HTTPRouteAction[i5].BackendGroupIDRef = rsp.ResolvedReference

			}
		}
	}

	return nil
}
