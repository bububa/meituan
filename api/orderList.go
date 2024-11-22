package api

import (
	"context"

	"github.com/bububa/meituan/core"
	"github.com/bububa/meituan/model"
)

// OrderList 订单列表查询接口（新版）
func OrderList(ctx context.Context, clt *core.SDKClient, req *model.OrderListRequest) (*model.OrderListResponse, error) {
	var resp model.OrderListResponse
	if err := clt.Get(ctx, "orderList", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
