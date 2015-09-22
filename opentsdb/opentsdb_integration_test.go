//
// +build integration

package opentsdb

import (
	"bytes"
	"encoding/gob"
	"os"
	"testing"
	"time"

	"github.com/intelsdi-x/pulse/control/plugin"
	"github.com/intelsdi-x/pulse/core/ctypes"

	. "github.com/smartystreets/goconvey/convey"
)

func TestOpentsdbPublish(t *testing.T) {
	config := make(map[string]ctypes.ConfigValue)

	Convey("Pulse Plugin integration testing with OpenTSDB", t, func() {
		var buf bytes.Buffer
		buf.Reset()
		enc := gob.NewEncoder(&buf)

		config["host"] = ctypes.ConfigValueStr{Value: os.Getenv("PULSE_OPENTSDB_HOST")}
		config["port"] = ctypes.ConfigValueInt{Value: 4242}

		op := NewOpentsdbPublisher()
		cp := op.GetConfigPolicy()
		cfg, _ := cp.Get([]string{""}).Process(config)

		Convey("Publish float metrics to OpenTSDB", func() {
			metrics := []plugin.PluginMetricType{
				*plugin.NewPluginMetricType([]string{"/psutil/load/load15"}, time.Now(), "mac1", 23.1),
				*plugin.NewPluginMetricType([]string{"/psutil/vm/available"}, time.Now().Add(2*time.Second), "mac2", 23.2),
				*plugin.NewPluginMetricType([]string{"/psutil/load/load1"}, time.Now().Add(3*time.Second), "linux3", 23.3),
			}
			enc.Encode(metrics)

			err := op.Publish(plugin.PulseGOBContentType, buf.Bytes(), *cfg)
			So(err, ShouldBeNil)
		})

		Convey("Publish int metrics to OpenTSDB", func() {
			metrics := []plugin.PluginMetricType{
				*plugin.NewPluginMetricType([]string{"/psutil/vm/free"}, time.Now().Add(5*time.Second), "linux7", 23),
			}
			enc.Encode(metrics)

			err := op.Publish(plugin.PulseGOBContentType, buf.Bytes(), *cfg)
			So(err, ShouldBeNil)
		})
	})
}
