// +build medium

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
	"bytes"
	"encoding/gob"
	"os"
	"testing"
	"time"

	"github.com/intelsdi-x/snap/control/plugin"
	"github.com/intelsdi-x/snap/core"
	"github.com/intelsdi-x/snap/core/ctypes"

	. "github.com/smartystreets/goconvey/convey"
)

func TestOpentsdbPublish(t *testing.T) {
	config := make(map[string]ctypes.ConfigValue)

	Convey("Snap Plugin integration testing with OpenTSDB", t, func() {
		var buf bytes.Buffer
		buf.Reset()
		enc := gob.NewEncoder(&buf)

		config["host"] = ctypes.ConfigValueStr{Value: os.Getenv("SNAP_OPENTSDB_HOST")}
		config["port"] = ctypes.ConfigValueInt{Value: 4242}

		op := NewOpentsdbPublisher()
		cp, _ := op.GetConfigPolicy()
		cfg, _ := cp.Get([]string{""}).Process(config)
		tags := map[string]string{}
		tags[core.STD_TAG_PLUGIN_RUNNING_ON] = "mac1"

		Convey("Publish float metrics to OpenTSDB", func() {
			metrics := []plugin.MetricType{
				*plugin.NewMetricType(core.NewNamespace("/psutil/load/load15"), time.Now(), tags, "float64", 23.1),
				*plugin.NewMetricType(core.NewNamespace("/psutil/vm/available"), time.Now().Add(2*time.Second), tags, "float64", 23.2),
				*plugin.NewMetricType(core.NewNamespace("/psutil/load/load1"), time.Now().Add(3*time.Second), tags, "float64", 23.3),
			}
			enc.Encode(metrics)

			err := op.Publish(plugin.SnapGOBContentType, buf.Bytes(), *cfg)
			So(err, ShouldBeNil)
		})

		Convey("Publish int metrics to OpenTSDB", func() {
			metrics := []plugin.MetricType{
				*plugin.NewMetricType(core.NewNamespace("/psutil/vm/free"), time.Now().Add(5*time.Second), tags, "int", 23),
			}
			enc.Encode(metrics)

			err := op.Publish(plugin.SnapGOBContentType, buf.Bytes(), *cfg)
			So(err, ShouldBeNil)
		})
	})
}
