/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package config

import (
	"context"
	"fmt"
	"io"
	"path/filepath"
)

type (
	// ProtocolType protocol type enum
	ProtocolType int32

	// PathKey config path key type
	PathKey string
)

var (
	_rootPathTemp                       = "/%s/arana-db"
	DefaultRootPath                     PathKey
	DefaultConfigMetadataPath           PathKey
	DefaultTenantsPath                  PathKey
	DefaultTenantBaseConfigPath         PathKey
	DefaultConfigDataFiltersPath        PathKey
	DefaultConfigDataNodesPath          PathKey
	DefaultConfigDataUsersPath          PathKey
	DefaultConfigDataSourceClustersPath PathKey
	DefaultConfigDataShardingRulePath   PathKey
	DefaultConfigDataShadowRulePath     PathKey
)

func initPath(root string) {
	DefaultRootPath = PathKey(fmt.Sprintf(_rootPathTemp, root))
	DefaultConfigMetadataPath = PathKey(filepath.Join(string(DefaultRootPath), "metadata"))
	DefaultTenantsPath = PathKey(filepath.Join(string(DefaultRootPath), "tenants"))
	DefaultTenantBaseConfigPath = PathKey(filepath.Join(string(DefaultRootPath), "tenants/%s"))
	DefaultConfigDataFiltersPath = PathKey(filepath.Join(string(DefaultTenantBaseConfigPath), "filters"))
	DefaultConfigDataNodesPath = PathKey(filepath.Join(string(DefaultTenantBaseConfigPath), "nodes"))
	DefaultConfigDataUsersPath = PathKey(filepath.Join(string(DefaultTenantBaseConfigPath), "users"))
	DefaultConfigDataSourceClustersPath = PathKey(filepath.Join(string(DefaultTenantBaseConfigPath), "dataSourceClusters"))
	DefaultConfigDataShardingRulePath = PathKey(filepath.Join(string(DefaultConfigDataSourceClustersPath), "shardingRule"))
	DefaultConfigDataShadowRulePath = PathKey(filepath.Join(string(DefaultConfigDataSourceClustersPath), "shadowRule"))
}

const (
	Http ProtocolType = iota
	MySQL
)

const (
	_            DataSourceType = ""
	DBMySQL      DataSourceType = "mysql"
	DBPostgreSQL DataSourceType = "postgresql"
)

var (
	slots        = make(map[string]StoreOperate)
	storeOperate StoreOperate
)

// Register register store plugin
func Register(s StoreOperate) {
	if _, ok := slots[s.Name()]; ok {
		panic(fmt.Errorf("StoreOperate=[%s] already exist", s.Name()))
	}

	slots[s.Name()] = s
}

type (
	//EventSubscriber
	EventSubscriber interface {
		//OnEvent
		OnEvent(event Event)
		//Type
		Type() EventType
	}

	Options struct {
		StoreName string                 `yaml:"name"`
		RootPath  string                 `yaml:"root_path"`
		Options   map[string]interface{} `yaml:"options"`
	}

	//TenantOperate 专门针对租户空间的相关操作
	TenantOperate interface {
		Tenants() ([]string, error)

		CreateTenant(string) error

		RemoveTenant(string) error
	}

	// Center
	Center interface {
		io.Closer

		//Load 每次都是拉取全量的配置数据
		Load(ctx context.Context) (*Tenant, error)

		//Import 导入配置，仅针对某个命名空间下
		Import(ctx context.Context, cfg *Tenant) error

		Subscribe(ctx context.Context, s ...EventSubscriber)

		UnSubscribe(ctx context.Context, s ...EventSubscriber)

		Tenant() string
	}

	// StoreOperate config storage related plugins
	StoreOperate interface {
		io.Closer

		// Init plugin initialization
		Init(options map[string]interface{}) error

		// Save save a configuration data
		Save(key PathKey, val []byte) error

		// Get get a configuration
		Get(key PathKey) ([]byte, error)

		// Watch Monitor changes of the key
		Watch(key PathKey) (<-chan []byte, error)

		// Name plugin name
		Name() string
	}
)
