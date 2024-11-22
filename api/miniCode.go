package api

import (
	"context"

	"github.com/bububa/meituan/core"
	"github.com/bububa/meituan/model"
)

// MiniCode 小程序生成二维码
func MiniCode(ctx context.Context, clt *core.SDKClient, req *model.MiniCodeRequest) (string, error) {
	var resp model.MiniCodeResponse
	if err := clt.Get(ctx, "miniCode", req, &resp); err != nil {
		return "", err
	}
	return resp.Data, nil
}
