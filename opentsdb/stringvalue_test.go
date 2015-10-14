// +build unit

/*
http://www.apache.org/licenses/LICENSE-2.0.txt


Copyright 2015 Intel Coporation

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
	"bytes"
	"encoding/json"
	"testing"
)

var stringtests = []struct {
	sv   StringValue
	json []byte
}{
	{StringValue("dog-cat-22"), []byte(`"dog-cat-22"`)},
	{StringValue("dog_cat_22"), []byte(`"dog__cat__22"`)},
	{StringValue("http://google.com:8080"), []byte(`"http_.//google.com_.8080"`)},
	{StringValue("/psutil/load/load15"), []byte(`"/psutil/load/load15"`)},
	{StringValue("/psutil/vm/free"), []byte(`"/psutil/vm/free"`)},
}

func TestStringValueMarshaling(t *testing.T) {
	for i, tt := range stringtests {
		json, err := json.Marshal(tt.sv)
		if err != nil {
			t.Errorf("%d. Marshal(%q) returned err: %s", i, tt.sv, err)
		} else {
			if !bytes.Equal(json, tt.json) {
				t.Errorf(
					"%d. Marshal(%q) => %q, want %q",
					i, tt.sv, json, tt.json,
				)
			}
		}
	}
}

func TestStringValueUnMarshaling(t *testing.T) {
	for i, tt := range stringtests {
		var sv StringValue
		err := json.Unmarshal(tt.json, &sv)
		if err != nil {
			t.Errorf("%d. Unmarshal(%q, &str) returned err: %s", i, tt.json, err)
		} else {
			if sv != tt.sv {
				t.Errorf(
					"%d. Unmarshal(%q, &str) => str==%q, want %q",
					i, tt.json, sv, tt.sv,
				)
			}
		}
	}
}
