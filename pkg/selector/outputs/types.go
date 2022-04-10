// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

package outputs

import (
	"reflect"
	"strings"
)

// SortOptions can be set to customize the sorting on Output functions
type SortOptions struct {
	SortColumns []SortColumn
}

// SortColumns control the column and sorting mode (ascending | descending)
type SortColumn struct {
	Column string
	Mode   string
}

// SortOptionsFrom parses a string representation of sort options into a SortOption struct
// i.e. "vcpus:asc,memory:desc"
func SortOptionsFrom(sortOptions *string) *SortOptions {
	if sortOptions == nil {
		return nil
	}
	sortOpts := &SortOptions{}
	for _, column := range strings.Split(*sortOptions, ",") {
		tokens := strings.Split(column, ":")
		mode := "asc"
		if len(tokens) == 2 {
			mode = tokens[1]
		}
		sortOpts.SortColumns = append(sortOpts.SortColumns, SortColumn{
			Column: tokens[0],
			Mode:   mode,
		})
	}
	return sortOpts
}

type TableWide struct {
	InstanceType         string  `header:Instance Type`
	VCPUs                int     `header:VCPUs`
	Memory               float64 `header:Mem (GIB)`
	Hypervisor           string  `header:Hypervisor`
	CurrentGeneration    bool    `header:Current Gen`
	HibernationSupport   bool    `header:Hibernation Support`
	CPUArchitecture      string  `heaader:CPU Arch`
	NetworkPerformance   string  `header:Network Performance`
	ENIs                 int     `header:ENIs`
	GPUs                 int     `header:GPUs`
	GPUMemory            float64 `header: GPU Mem (GIB)`
	GPUInfo              string  `header:GPU INFO`
	OndemandPricePerHour string  `header:On-Demand Price/Hr`
	SpotPricePerHour     string  `header:Spot Price/Hr (30D AVG)`
}

func (tw TableWide) Headers() []interface{} {
	v := reflect.ValueOf(tw)
	for i := 0; i < v.NumField(); i++ {

	}
}
