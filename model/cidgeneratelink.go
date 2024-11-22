package model

import (
	"net/url"
	"strconv"

	"github.com/bububa/meituan/enum"
)

// CidGenerateLinkRequest CID取链 API Request
type CidGenerateLinkRequest struct {
	// PoiID
	PoiID string `json:"poiId,omitempty"`
	// SkuID
	SkuID string `json:"skuId,omitempty"`
	// AttributeType 1：美团平台归因
	AttributeType enum.AttributeType `json:"attributeType,omitempty"`
	// Media 抖音：2
	Media enum.Media `json:"media,omitempty"`
	// AttributeRule 1.cid-医药-spu-支付
	// 2.cid-医药-单店-支付
	// 归因规则（标品还是单店）
	AttributeRule enum.AttributeRule `json:"attributeRule,omitempty"`
	// ActID 活动id,本次默认用366号活动
	ActID uint64 `json:"actId,omitempty"`
}

func (r CidGenerateLinkRequest) Values(values url.Values) {
	values.Set("poiId", r.PoiID)
	values.Set("skuId", r.SkuID)
	values.Set("attributeType", strconv.Itoa(int(r.AttributeType)))
	values.Set("media", strconv.Itoa(int(r.Media)))
	values.Set("attributeRule", strconv.Itoa(int(r.AttributeRule)))
	values.Set("actId", strconv.FormatUint(r.ActID, 10))
}

// CidGenerateLinkResponse CID取链 API Response
type CidGenerateLinkResponse struct {
	Data *CidLinkInfo `json:"data,omitempty"`
	BaseResponse
}

type CidLinkInfo struct {
	// DeepLinkURL 直达链接（老版）
	DeepLinkURL string `json:"deepLinkUrl,omitempty"`
	// NewDeepLinkURL 新版直达链接（推荐使用）
	NewDeepLinkURL string `json:"newDeepLinkUrl,omitempty"`
	// WechatLink 小程序链接
	WechatLink string `json:"wechatLink,omitempty"`
	// MonitorLink 点击监测链接
	// 如果使用取链接口获取直达/小程序链接，监测链接也需要使用通过取链接口获取的。人工离线提供的链接和取链接口获取的链接不能混用，否则事件回传不准确。
	MonitorLink string `json:"monitorLink,omitempty"`
	// H5Link H5 活动页链接（可作为投放兜底链接）
	H5Link string `json:"h5Link,omitempty"`
}
