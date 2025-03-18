package api

import (
	"context"

	"github.com/bububa/meituan/v2/core"
	"github.com/bububa/meituan/v2/model"
)

// QueryCoupon 商品查询接口
// 查询售卖商品接口，支持全量查询、精确查询、榜单主题查询。需用POST方式调用。只接受JSON格式。
func QueryCoupon(ctx context.Context, clt *core.SDKClient, req *model.QueryCouponRequest, resp *model.QueryCouponResponse) error {
	return clt.POST(ctx, "query_coupon", req, resp)
}
