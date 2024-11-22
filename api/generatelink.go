package api

import (
	"context"

	"github.com/bububa/meituan/core"
	"github.com/bububa/meituan/model"
)

// GenerateLink  自助取链接口（新版）
func GenerateLink(ctx context.Context, clt *core.SDKClient, req *model.GenerateLinkRequest) (string, error) {
	var resp model.GenerateLinkResponse
	if err := clt.Get(ctx, "generatelink", req, &resp); err != nil {
		return "", err
	}
	return resp.Data, nil
}
