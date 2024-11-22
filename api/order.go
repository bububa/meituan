package api

import (
	"context"

	"github.com/bububa/meituan/core"
	"github.com/bububa/meituan/model"
)

// Order 单订单查询接口（新版）
func Order(ctx context.Context, clt *core.SDKClient, req *model.OrderRequest) (*model.Order, error) {
	var resp model.OrderResponse
	if err := clt.Get(ctx, "order", req, &resp); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
