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
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

const (
	putEndPoint     = "/api/put"
	contentTypeJson = "application/json"
	userAgent       = "snap-publisher"
)

type HttpClient struct {
	url        string
	httpClient *http.Client
	userAgent  string
}

type Client interface {
	NewClient(url string, timeout time.Duration) *HttpClient
}

//NewClient creates an instance of HttpClient which times out at
//the givin duration.
func NewClient(url string, timeout time.Duration) *HttpClient {
	return &HttpClient{
		url: url,
		httpClient: &http.Client{
			Timeout: timeout,
		},
		userAgent: userAgent,
	}
}

func (hc *HttpClient) getUrl() string {
	u := url.URL{
		Scheme: "http",
		Host:   hc.url,
		Path:   putEndPoint,
	}
	return u.String()
}

// Post stores slides of Datapoint to OpenTSDB
func (hc *HttpClient) Post(dps []DataPoint) error {
	url := hc.getUrl()

	buf, err := json.Marshal(dps)
	if err != nil {
		return err
	}

	resp, err := hc.httpClient.Post(url, contentTypeJson, bytes.NewReader(buf))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNoContent {
		return nil
	}

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var result map[string]int
	if err := json.Unmarshal(content, &result); err != nil {
		return err
	}
	return fmt.Errorf("failed to post %d data to OpenTSDB, %d succeeded", result["failed"], result["success"])
}
