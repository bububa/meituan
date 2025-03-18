package api

import (
	"context"

	"github.com/bububa/meituan/v2/core"
	"github.com/bububa/meituan/v2/model"
)

// GetReferralLink 获取推广链接接口
// 支持获取活动物料、到店/到家/买菜业务类型的推广链接；支持按活动物料ID、商品券展示ID、目标链接的形式获取对应的推广链接；支持appkey-sid两级渠道追踪推广效果。需要用POST方法调用接口。
func GetReferralLink(ctx context.Context, clt *core.SDKClient, req *model.GetReferralLinkRequest, resp *model.GetReferralLinkResponse) error {
	return clt.POST(ctx, "get_referral_link", req, resp)
}
