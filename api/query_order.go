package api

import (
	"context"

	"github.com/bububa/meituan/v2/core"
	"github.com/bububa/meituan/v2/model"
)

// QueryOrder 查询订单接口
// 查询推广的订单明细及佣金信息，包括到店、到家、买菜等业务类型的订单。支持按付款时间或更新时间查询，查询近3个月的订单明细。支持POST方法查询接口。只接受JSON格式。
func QueryOrder(ctx context.Context, clt *core.SDKClient, req *model.QueryOrderRequest, resp *model.QueryOrderResponse) error {
	return clt.POST(ctx, "query_order", req, resp)
}
