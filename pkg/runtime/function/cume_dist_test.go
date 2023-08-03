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

package function

import (
	"context"
	"fmt"
	"testing"
)

import (
	"github.com/stretchr/testify/assert"
)

import (
	"github.com/arana-db/arana/pkg/proto"
)

func TestFuncCumeDist(t *testing.T) {
	fn := proto.MustGetFunc(FuncCumeDist)
	assert.Equal(t, 0, fn.NumInput())

	type tt struct {
		inputs []proto.Value
		want   string
	}
	for _, v := range []tt{
		{
			[]proto.Value{
				proto.NewValueFloat64(1),
				proto.NewValueFloat64(1),
				proto.NewValueFloat64(1),
				proto.NewValueFloat64(2),
				proto.NewValueFloat64(3),
				proto.NewValueFloat64(3),
				proto.NewValueFloat64(3),
				proto.NewValueFloat64(4),
				proto.NewValueFloat64(4),
				proto.NewValueFloat64(5),
			},
			"0.2222222222222222",
		},
		{
			[]proto.Value{
				proto.NewValueFloat64(2),
				proto.NewValueFloat64(1),
				proto.NewValueFloat64(1),
				proto.NewValueFloat64(2),
				proto.NewValueFloat64(3),
				proto.NewValueFloat64(3),
				proto.NewValueFloat64(3),
				proto.NewValueFloat64(4),
				proto.NewValueFloat64(4),
				proto.NewValueFloat64(5),
			},
			"0.3333333333333333",
		},
		{
			[]proto.Value{
				proto.NewValueFloat64(3),
				proto.NewValueFloat64(1),
				proto.NewValueFloat64(1),
				proto.NewValueFloat64(2),
				proto.NewValueFloat64(3),
				proto.NewValueFloat64(3),
				proto.NewValueFloat64(3),
				proto.NewValueFloat64(4),
				proto.NewValueFloat64(4),
				proto.NewValueFloat64(5),
			},
			"0.6666666666666666",
		},
		{
			[]proto.Value{
				proto.NewValueFloat64(4),
				proto.NewValueFloat64(1),
				proto.NewValueFloat64(1),
				proto.NewValueFloat64(2),
				proto.NewValueFloat64(3),
				proto.NewValueFloat64(3),
				proto.NewValueFloat64(3),
				proto.NewValueFloat64(4),
				proto.NewValueFloat64(4),
				proto.NewValueFloat64(5),
			},
			"0.8888888888888888",
		},
		{
			[]proto.Value{
				proto.NewValueFloat64(5),
				proto.NewValueFloat64(1),
				proto.NewValueFloat64(1),
				proto.NewValueFloat64(2),
				proto.NewValueFloat64(3),
				proto.NewValueFloat64(3),
				proto.NewValueFloat64(3),
				proto.NewValueFloat64(4),
				proto.NewValueFloat64(4),
				proto.NewValueFloat64(5),
			},
			"1",
		},
	} {
		t.Run(v.want, func(t *testing.T) {
			var inputs []proto.Valuer
			for i := range v.inputs {
				inputs = append(inputs, proto.ToValuer(v.inputs[i]))
			}
			out, err := fn.Apply(context.Background(), inputs...)
			assert.NoError(t, err)
			assert.Equal(t, v.want, fmt.Sprint(out))
		})
	}
}

func TestFuncCumeDist_Error(t *testing.T) {
	fn := proto.MustGetFunc(FuncCumeDist)
	assert.Equal(t, 0, fn.NumInput())

	type tt struct {
		name   string
		inputs []proto.Value
	}
	for _, v := range []tt{
		{
			"Test_Nil",
			[]proto.Value{
				nil,
			},
		},
		{
			"Test_Nil_1",
			[]proto.Value{
				proto.NewValueFloat64(5),
				nil,
			},
		},
	} {
		t.Run(v.name, func(t *testing.T) {
			var inputs []proto.Valuer
			for i := range v.inputs {
				inputs = append(inputs, proto.ToValuer(v.inputs[i]))
			}
			out, err := fn.Apply(context.Background(), inputs...)
			assert.Nil(t, err)
			assert.Nil(t, out)
		})
	}
}
