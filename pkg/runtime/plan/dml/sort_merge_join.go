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
