// Copyright © 2018 MTRNord <info@nordgedanken.de>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package ui

import (
	"encoding/json"
	"net/http"
	"sort"
	"time"
)

type helloMatrixResp []struct {
	Hostname             string `json:"hostname"`
	Description          string `json:"description"`
	URL                  string `json:"url"`
	Category             string `json:"category"`
	Location             string `json:"location"`
	OnlineSince          int64  `json:"online_since"`
	LastResponse         int64  `json:"last_response"`
	LastResponseTime     int64  `json:"last_response_time"`
	StatusSince          string `json:"status_since"`
	LastVersions         string `json:"last_versions"`
	Measurements         int64  `json:"measurements"`
	Successful           int64  `json:"successful"`
	SumResponseTime      int64  `json:"sum_response_time"`
	MeasurementsShort    int64  `json:"measurements_short"`
	SuccessfulShort      int64  `json:"successful_short"`
	SumResponseTimeShort int64  `json:"sum_response_time_short"`
	UsersActive          int64  `json:"users_active,omitempty"`
	ServerName           string `json:"server_name"`
	ServerVersion        string `json:"server_version"`
	Grade                string `json:"grade"`
	GradeTrustIgnored    string `json:"gradeTrustIgnored"`
	HasWarnings          int64  `json:"hasWarnings"`
	PublicRoomCount      int64  `json:"public_room_count"`
}

func getHelloMatrixList() (resp helloMatrixResp, err error) {
	var httpClient = &http.Client{Timeout: 10 * time.Second}

	url := "https://www.hello-matrix.net/public_servers.php?format=json&only_public=true&client=Morpheusv2"

	r, RespErr := httpClient.Get(url)
	if RespErr != nil {
		err = RespErr
		return
	}
	defer r.Body.Close()

	decodeErr := json.NewDecoder(r.Body).Decode(&resp)
	if decodeErr != nil {
		err = decodeErr
		return
	}

	return
}

func convertHelloMatrixRespToNameSlice(resp helloMatrixResp) (hostnames []string) {
	hostnames = append(hostnames, selectMessage)

	sort.Slice(resp, func(i, j int) bool {
		return resp[i].LastResponseTime < resp[i].LastResponseTime
	})
	for _, v := range resp {
		hostnames = append(hostnames, v.Hostname)
	}

	return
}
