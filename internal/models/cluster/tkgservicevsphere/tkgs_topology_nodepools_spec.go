/*
Copyright © 2021 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package tkgservicevspheremodel

import "github.com/go-openapi/swag"

// VmwareTanzuManageV1alpha1ClusterNodepoolSpec Spec for the cluster nodepool.
//
// swagger:model vmware.tanzu.manage.v1alpha1.cluster.nodepool.Spec
type VmwareTanzuManageV1alpha1ClusterNodepoolSpec struct {

	// Cloud labels.
	CloudLabels map[string]string `json:"cloudLabels,omitempty"`

	// Node labels.
	NodeLabels map[string]string `json:"nodeLabels,omitempty"`

	// Nodepool config for tkg service vsphere.
	TkgServiceVsphere *VmwareTanzuManageV1alpha1ClusterNodepoolTKGServiceVsphereNodepool `json:"tkgServiceVsphere,omitempty"`

	// Count is the number of nodes.
	WorkerNodeCount string `json:"workerNodeCount,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterNodepoolSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterNodepoolSpec) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClusterNodepoolSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
