/*
Copyright The Kubernetes Authors.

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

// Code generated by client-gen. DO NOT EDIT.

package v1beta1

import (
	rest "k8s.io/client-go/rest"

	"github.com/erda-project/erda/pkg/clientgo/apis/flinkoperator/v1beta1"
	"github.com/erda-project/erda/pkg/clientgo/restclient"
)

type FlinkoperatorV1beta1Interface interface {
	RESTClient() rest.Interface
	FlinkClustersGetter
}

// FlinkoperatorV1beta1Client is used to interact with features provided by the flinkoperator group.
type FlinkoperatorV1beta1Client struct {
	restClient rest.Interface
}

func (c *FlinkoperatorV1beta1Client) FlinkClusters(namespace string) FlinkClusterInterface {
	return newFlinkClusters(c, namespace)
}

// New creates a new FlinkoperatorV1beta1Client for the given RESTClient.
func New(c rest.Interface) *FlinkoperatorV1beta1Client {
	return &FlinkoperatorV1beta1Client{c}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FlinkoperatorV1beta1Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}

// NewFlinkOpeartorClient creates a new OperatorV1beta1 for the given addr.
func NewFlinkOpeartorClient(addr string) (*FlinkoperatorV1beta1Client, error) {
	config := restclient.GetDefaultConfig("/apis")
	config.GroupVersion = &v1beta1.SchemeGroupVersion
	var client rest.Interface
	var err error
	if addr != "" {
		client, err = restclient.NewInetRESTClient(addr, config)
	} else {
		client, err = rest.RESTClientFor(config)
	}
	if err != nil {
		return nil, err
	}
	return New(client), nil
}
