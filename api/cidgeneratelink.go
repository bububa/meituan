package api

import (
	"context"

	"github.com/bububa/meituan/core"
	"github.com/bububa/meituan/model"
)

// CidGenerateLink CID取链
func CidGenerateLink(ctx context.Context, clt *core.SDKClient, req *model.CidGenerateLinkRequest) (*model.CidLinkInfo, error) {
	var resp model.CidGenerateLinkResponse
	if err := clt.Get(ctx, "cidgeneratelink", req, &resp); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
