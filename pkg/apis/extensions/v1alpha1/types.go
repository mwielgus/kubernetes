/*
Copyright 2015 The Kubernetes Authors All rights reserved.

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

package v1alpha1

import (
	"k8s.io/kubernetes/pkg/api/resource"
)

// Alpha-level support for Custom Metrics in HPA (as annotations).
type CustomMetricUtilization struct {
	// Custom Metric name.
	Name string `json:"name"`
	// Custom Metric value (average).
	Value resource.Quantity `json:"value"`
}

type CustomMetricUtiliztionList struct {
	Items []CustomMetricUtilization `json:"items"`
}

// An APIVersion represents a single concrete version of an object model.
type APIVersion struct {
        // Name of this version (e.g. 'v1').
        Name string `json:"name,omitempty"`

        // The API group to add this object into, default 'experimental'.
        APIGroup string `json:"apiGroup,omitempty"`
}