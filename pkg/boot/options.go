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

package boot

import (
	"github.com/arana-db/arana/pkg/config"
)

type (
	BootOptions struct {
		Config    *config.Options `yaml:"config"`
		Listeners []*Listener     `validate:"required,dive" yaml:"listeners" json:"listeners"`
	}

	// SocketAddress specify either a logical or physical address and port, which are
	// used to tell server where to bind/listen, connect to upstream and find
	// management servers
	SocketAddress struct {
		Address string `default:"0.0.0.0" yaml:"address" json:"address"`
		Port    int    `default:"13306" yaml:"port" json:"port"`
	}

	Listener struct {
		ProtocolType  string         `yaml:"protocol_type" json:"protocol_type"`
		SocketAddress *SocketAddress `yaml:"socket_address" json:"socket_address"`
		ServerVersion string         `yaml:"server_version" json:"server_version"`
	}
)
