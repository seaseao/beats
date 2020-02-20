// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

// Code generated by beats/dev-tools/cmd/asset/asset.go - DO NOT EDIT.

package greenplum

import (
	"github.com/elastic/beats/libbeat/asset"
)

func init() {
	if err := asset.SetFields("filebeat", "greenplum", asset.ModuleFieldsPri, AssetGreenplum); err != nil {
		panic(err)
	}
}

// AssetGreenplum returns asset data.
// This is the base64 encoded gzipped contents of module/greenplum.
func AssetGreenplum() string {
	return "eJzMl8Fu2zAMhu95CqL3vkAOe4MBewODllhHiCRqIrXVbz+4aB3HsbHmIiuHHBjq1xfxt0i/wpXGMwyZKCZfwglAnXo6w8scezkBWBKTXVLH8Qw/TgBwWwM/2RZPJ4A3R97K+ePnV4gY6F56+uiYPqJc0mdkQ/teainneZhjq6W/PKEQoLXL+CKb3jGk6b8tMj9ji6xHwC2eG5G6QKIYltnPkn2X7WvPIpQrbPd1GFca/3K2GyAWFXsUagImZTYk0jnbBI5eMqFthebCohVBlN51q0Sca1J4jsMGhZCI49iJYhWc2REZo6CZVOra4j/n0IhFDYeA0XaGS2zAJUPqhIZAVVnWfWer8ywMNaZ7kefxvg+4X8Ab0MpJx+DM9vbOUBMPmnWi2fVFyXaHXQL7B+XZoG8RTErfIJWnP+SbIBFFpc6wbWMACySCQ02WnTHDkqKrWaIdjour2jx2KH4XymMjGF1icQ/6h1zKhuOK9DC39mU4HsOULJwbKtCb8zXvkudGr+m7+qzz4Ne5JblYn2bvrULRXGvb+V8AAAD//4wOTP8="
}
