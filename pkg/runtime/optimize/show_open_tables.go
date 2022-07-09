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

package optimize

import (
	"context"
)

import (
	"github.com/arana-db/arana/pkg/proto"
	"github.com/arana-db/arana/pkg/runtime/ast"
	rcontext "github.com/arana-db/arana/pkg/runtime/context"
	"github.com/arana-db/arana/pkg/runtime/namespace"
	"github.com/arana-db/arana/pkg/runtime/plan"
	"github.com/arana-db/arana/pkg/security"
	"github.com/arana-db/arana/pkg/transformer"
)

func init() {
	registerOptimizeHandler(ast.SQLTypeShowOpenTables, optimizeShowOpenTables)
}

func optimizeShowOpenTables(ctx context.Context, o *optimizer) (proto.Plan, error) {
	var invertedIndex map[string]string
	for logicalTable, v := range o.rule.VTables() {
		t := v.Topology()
		t.Each(func(x, y int) bool {
			if _, phyTable, ok := t.Render(x, y); ok {
				if invertedIndex == nil {
					invertedIndex = make(map[string]string)
				}
				invertedIndex[phyTable] = logicalTable
			}
			return true
		})
	}

	stmt := o.stmt.(*ast.ShowOpenTables)

	clusters := security.DefaultTenantManager().GetClusters(rcontext.Tenant(ctx))
	plans := make([]proto.Plan, 0, len(clusters))
	for _, cluster := range clusters {
		ns := namespace.Load(cluster)
		// 配置里原子库 都需要执行一次
		groups := ns.DBGroups()
		for i := 0; i < len(groups); i++ {
			ret := plan.NewShowOpenTablesPlan(stmt)
			ret.BindArgs(o.args)
			ret.SetInvertedShards(invertedIndex)
			ret.SetDatabase(groups[i])
			plans = append(plans, ret)
		}
	}

	unionPlan := &plan.UnionPlan{
		Plans: plans,
	}

	aggregate := &plan.AggregatePlan{
		Plan:       unionPlan,
		Combiner:   transformer.NewCombinerManager(),
		AggrLoader: transformer.LoadAggrs(nil),
	}

	return aggregate, nil
}
