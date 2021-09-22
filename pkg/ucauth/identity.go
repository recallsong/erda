// Copyright (c) 2021 Terminus, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package ucauth

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/erda-project/erda/pkg/http/httpclient"
)

func getIdentity(kratosPrivateAddr string, userID string) (*OryKratosIdentity, error) {
	var body bytes.Buffer
	r, err := httpclient.New(httpclient.WithCompleteRedirect()).
		Get(kratosPrivateAddr).
		Path("/identities/" + userID).
		Do().Body(&body)
	if err != nil {
		return nil, err
	}
	if !r.IsOK() {
		return nil, fmt.Errorf("get identity: statuscode: %d, body: %v", r.StatusCode(), body.String())
	}
	var i OryKratosIdentity
	if err := json.Unmarshal(body.Bytes(), &i); err != nil {
		return nil, err
	}
	return &i, nil
}
