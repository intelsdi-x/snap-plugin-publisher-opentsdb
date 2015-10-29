/*
http://www.apache.org/licenses/LICENSE-2.0.txt


Copyright 2015 Intel Corporation

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

package opentsdb

import (
	"fmt"
	"strconv"
)

const (
	EmptyString = ""
)

type DataPoint struct {
	Metric    StringValue            `json:"metric"`
	Timestamp int64                  `json:"timestamp"`
	Value     interface{}            `json:"value"`
	Tags      map[string]StringValue `json:"tags"`
}

// Valid verifies the mandatory fields of the Datapoint.
func (d *DataPoint) Valid() bool {
	if d.Metric == EmptyString || d.Value == nil || d.Timestamp == 0 || len(d.Tags) == 0 {
		return false
	}

	if _, err := strconv.ParseFloat(fmt.Sprint(d.Value), 64); err != nil {
		return false
	}
	return true
}
