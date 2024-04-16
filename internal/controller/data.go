package controller

import (
	"context"
	"goframe-shop-v2/api/backend"
	"goframe-shop-v2/internal/service"
)

// Data 内容管理
var Data = cData{}

type cData struct{}

func (a *cData) DataHead(ctx context.Context, req *backend.DateHeadReq) (res *backend.DateHeadRes, err error) {
	out, err := service.Data().DataHead(ctx)
	if err != nil {
		return nil, err
	}
	return &backend.DateHeadRes{
		TodayOrderCount: out.TodayOrderCount,
		DAU:             out.DAU,
		ConversionRate:  out.ConversionRate,
	}, err
}
