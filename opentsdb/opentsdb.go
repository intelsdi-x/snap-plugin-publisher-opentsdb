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
	"encoding/json"
	"fmt"
	"net/url"
	"strings"
	"time"

	log "github.com/Sirupsen/logrus"

	"github.com/intelsdi-x/pulse/control/plugin"
	"github.com/intelsdi-x/pulse/control/plugin/cpolicy"
	"github.com/intelsdi-x/pulse/core/ctypes"
)

const (
	name       = "opentsdb"
	version    = 2
	pluginType = plugin.PublisherPluginType
	timeout    = 5
	host       = "host"
	source     = "source"
)

// Meta returns a plugin meta data
func Meta() *plugin.PluginMeta {
	return plugin.NewPluginMeta(name, version, pluginType, []string{plugin.PulseGOBContentType}, []string{plugin.PulseGOBContentType})
}

//NewOpentsdbPublisher returns an instance of the OpenTSDB publisher
func NewOpentsdbPublisher() *opentsdbPublisher {
	return &opentsdbPublisher{}
}

type opentsdbPublisher struct {
}

func (p *opentsdbPublisher) GetConfigPolicy() (*cpolicy.ConfigPolicy, error) {
	cp := cpolicy.New()
	config := cpolicy.NewPolicyNode()

	r1, err := cpolicy.NewStringRule("host", true)
	handleErr(err)
	r1.Description = "Opentsdb host"
	config.Add(r1)

	r2, err := cpolicy.NewIntegerRule("port", true)
	handleErr(err)
	r2.Description = "Opentsdb port"
	config.Add(r2)

	cp.Add([]string{""}, config)
	return cp, nil
}

// Publish publishes metric data to opentsdb.
func (p *opentsdbPublisher) Publish(contentType string, content []byte, config map[string]ctypes.ConfigValue) error {
	logger := log.New()
	var metrics []plugin.PluginMetricType

	switch contentType {
	case plugin.PulseGOBContentType:
		dec := gob.NewDecoder(bytes.NewBuffer(content))
		if err := dec.Decode(&metrics); err != nil {
			logger.Printf("Error decoding GOB: error=%v content=%v", err, content)
			return err
		}
	case plugin.PulseJSONContentType:
		err := json.Unmarshal(content, &metrics)
		if err != nil {
			logger.Printf("Error decoding JSON: error=%v content=%v", err, content)
			return err
		}
	default:
		logger.Printf("Error unknown content type '%v'", contentType)
		return fmt.Errorf("Unknown content type '%s'", contentType)
	}

	u, err := url.Parse(fmt.Sprintf("%s:%d", config["host"].(ctypes.ConfigValueStr).Value, config["port"].(ctypes.ConfigValueInt).Value))
	if err != nil {
		handleErr(err)
	}

	pts := make([]DataPoint, len(metrics))
	var temp DataPoint
	var i = 0
	for _, m := range metrics {
		temp = DataPoint{
			Metric:    StringValue(strings.Join(m.Namespace(), ".")),
			Timestamp: m.Timestamp().Unix(),
			Value:     m.Data(),
			Tags: map[string]StringValue{
				host: StringValue(m.Source()),
			},
		}

		// Omits invalid data points
		if temp.Valid() {
			pts[i] = temp
			i++
		}
	}

	if len(pts) == 0 {
		logger.Printf("Info: '%s' posting metrics: %+v", "no valid data", metrics)
		return nil
	}

	td := time.Duration(timeout * time.Second)
	con := NewClient(u.String(), td)
	err = con.Post(pts)
	if err != nil {
		logger.Printf("Error: '%s' posting metrics: %+v", err.Error(), metrics)
		return err
	}

	return nil
}

func handleErr(e error) {
	if e != nil {
		panic(e)
	}
}
