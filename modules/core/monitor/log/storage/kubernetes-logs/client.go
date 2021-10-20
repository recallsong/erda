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

package kuberneteslogs

import (
	"fmt"

	"github.com/bluele/gcache"
	"github.com/erda-project/erda/pkg/k8sclient"
	"k8s.io/client-go/kubernetes"
)

// ClientManager .
type ClientManager interface {
	GetClient(clusterName string) (*kubernetes.Clientset, error)
}

type clientManager struct {
	cache gcache.Cache
}

func newClientManager(c *config) ClientManager {
	cache := gcache.New(c.ClientCacheSize).LRU().Expiration(c.ClientCacheExpiration).LoaderFunc(func(key interface{}) (interface{}, error) {
		client, err := k8sclient.New(key.(string))
		if err != nil {
			return nil, err
		}
		return client.ClientSet, nil
	}).Build()
	return &clientManager{
		cache: cache,
	}
}

func (cm *clientManager) GetClient(clusterName string) (*kubernetes.Clientset, error) {
	val, err := cm.cache.Get(clusterName)
	if err != nil {
		return nil, err
	}
	client, _ := val.(*kubernetes.Clientset)
	if client == nil {
		return nil, fmt.Errorf("not found clientset")
	}
	return client, nil
}