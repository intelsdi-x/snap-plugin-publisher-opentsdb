// +build unit

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
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestValid(t *testing.T) {
	dataPoint := DataPoint{"Temperature", 1442028635, 23.1, map[string]StringValue{"host": "abc"}}

	Convey("Valid assertions", t, func() {
		So(dataPoint.Valid(), ShouldBeTrue)
	})
}

func TestInvalid_1(t *testing.T) {
	dataPoint := DataPoint{}

	Convey("Invalid assertions", t, func() {
		So(dataPoint.Valid(), ShouldBeFalse)

		dataPoint = DataPoint{
			Metric: "test",
			Value:  123,
		}
		So(dataPoint.Valid(), ShouldBeFalse)

		dataPoint = DataPoint{
			Metric:    "test",
			Value:     123,
			Timestamp: 12345,
		}
		So(dataPoint.Valid(), ShouldBeFalse)

		dataPoint := DataPoint{
			Metric: "test",
			Value:  123,
			Tags: map[string]StringValue{
				"host": "abc",
			},
		}
		So(dataPoint.Valid(), ShouldBeFalse)

		dataPoint = DataPoint{
			Value:     123,
			Timestamp: 12345,
			Tags: map[string]StringValue{
				"host": "abc",
			},
		}
		So(dataPoint.Valid(), ShouldBeFalse)
	})
}
