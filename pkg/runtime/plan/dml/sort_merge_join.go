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

package dml

import (
	"context"

	"github.com/arana-db/arana/pkg/dataset"
	"github.com/arana-db/arana/pkg/proto"
	"github.com/arana-db/arana/pkg/resultx"
	"github.com/arana-db/arana/pkg/runtime/ast"
)

type SortMergeJoin struct {
	Stmt *ast.SelectStatement

	LeftQuery  proto.Plan
	RightQuery proto.Plan

	JoinType ast.JoinType
	LeftKey  string
	RightKey string
}

func (h *SortMergeJoin) Type() proto.PlanType {
	return proto.PlanTypeQuery
}

func (h *SortMergeJoin) ExecIn(ctx context.Context, conn proto.VConn) (proto.Result, error) {
	rightDs := fetchDs(h.RightQuery, ctx, conn)
	leftDs := fetchDs(h.LeftQuery, ctx, conn)

	joinColumn := &dataset.JoinColumn{}
	joinColumn.SetColumn(h.LeftKey)
	ds, err := dataset.NewSortMergeJoin(
		h.JoinType,
		joinColumn,
		rightDs,
		leftDs,
	)
	if err != nil {
		return nil, err
	}

	return resultx.New(resultx.WithDataset(ds)), nil
}

func fetchDs(plan proto.Plan, ctx context.Context, conn proto.VConn) proto.Dataset {
	res, err := plan.ExecIn(ctx, conn)
	if err != nil {
		return nil
	}

	ds, err := res.Dataset()
	if err != nil {
		return nil
	}

	return ds
}
