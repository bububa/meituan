package mtunion

import (
	"context"

	"github.com/bububa/meituan/core"
	"github.com/bububa/meituan/model/mtunion"
)

// Poi 门店POI查询（新版）
func Poi(ctx context.Context, clt *core.SDKClient, req *mtunion.PoiRequest) (*mtunion.PoiResult, error) {
	var resp mtunion.PoiResponse
	if err := clt.Get(ctx, "mtunion/poi", req, &resp); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
