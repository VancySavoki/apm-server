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

package metric

import (
	"github.com/elastic/apm-server/utility"
	"github.com/elastic/beats/libbeat/common"
)

type gauge struct {
	name  string
	value float64
	unit  *string
}

func (md *metricDecoder) decodeGauge(name string, raw map[string]interface{}) *gauge {
	return &gauge{
		name:  name,
		value: md.Float64(raw, "value"),
		unit:  md.StringPtr(raw, "unit"),
	}
}

func (g *gauge) mapstr() common.MapStr {
	v := common.MapStr{
		"type":  "gauge",
		"value": g.value,
	}
	utility.Add(v, "unit", g.unit)
	return v
}

func (g *gauge) transform(m common.MapStr) error {
	m[g.name] = g.mapstr()
	return nil
}
