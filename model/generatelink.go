package model

import (
	"net/url"
	"strconv"

	"github.com/bububa/meituan/enum"
)

// GenerateLinkRequest 自助取链接口（新版） API Request
type GenerateLinkRequest struct {
	// Sid 推广位sid，支持通过接口自定义创建，不受平台200个上限限制，长度不能超过64个字符，支持小写字母和数字，历史已创建的推广位不受这个约束
	// 备注：
	// 1、酒店业务不能包含大写字母及“:”
	// 2、优选业务不能包含“:”
	Sid string `json:"sid,omitempty"`
	// SkuViewID 商品券展示ID（加密过的）
	SkuViewID string `json:"skuViewId,omitempty"`
	// ActID 活动id
	ActID uint64 `json:"actId,omitempty"`
	// LinkType 链接类型，枚举值：
	// 1 h5链接
	// 2 deeplink(唤起)链接
	// 3 中间页唤起链接
	// 4 微信小程序唤起路径
	// 5 团口令
	LinkType enum.LinkType `json:"linkType,omitempty"`
	// ShortLink 获取长链还是短链
	// 0表示获取长链
	// 1表示获取短链
	ShortLink int `json:"shortLink,omitempty"`
}

func (r GenerateLinkRequest) Values(values url.Values) {
	if r.ActID > 0 {
		values.Set("actId", strconv.FormatUint(r.ActID, 10))
	}
	if r.Sid != "" {
		values.Set("sid", r.Sid)
	}
	if r.LinkType > 0 {
		values.Set("linkType", strconv.Itoa(int(r.LinkType)))
	}
	if r.ShortLink > 0 {
		values.Set("shortLink", strconv.Itoa(r.ShortLink))
	}
	if r.SkuViewID != "" {
		values.Set("skuViewId", r.SkuViewID)
	}
}

// GenerateLinkResponse 自助取链接口（新版） API Response
type GenerateLinkResponse struct {
	// Data 最终的推广链接
	// 默认生成的都是长链
	Data string `json:"data,omitempty"`
	BaseResponse
}
