package mtunion

import (
	"context"

	"github.com/bububa/meituan/core"
	"github.com/bububa/meituan/model/mtunion"
)

// Sku 券包商品查询接口
func Sku(ctx context.Context, clt *core.SDKClient, req *mtunion.SkuRequest) (*mtunion.SkuResult, error) {
	var resp mtunion.SkuResponse
	if err := clt.Get(ctx, "mtunion/sku", req, &resp); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
