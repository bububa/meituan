package model

import (
	"encoding/json"
	"io"
)

// GetReferralLinkRequest 获取推广链接接口 API Request
type GetReferralLinkRequest struct {
	// Platform 商品所属业务一级分类类型；请求的商品推广链接所属的业务类型信息，即只有输入skuViewId时才需要传本字段：1 到家及其他业务类型，2 到店业务类型；不填则默认1
	Platform int `json:"platform,omitempty"`
	// BizLine 商品所属业务二级分类类型；请求的商品推广链接所属的业务类型信息，即只有输入skuViewId时才需要传本字段；当字段platform为1，选择到家及其他业务类型时：5 医药，不填则默认null，表示外卖商品券；当字段platform为2，选择到店业务类型时：1 到餐，2 到综 3：酒店 4：门票 不填则默认1
	BizLine int `json:"bizLine,omitempty"`
	// ActID 活动物料ID，我要推广-活动推广中第一列的id信息（和商品id、活动链接三选一填写，不能全填）
	ActID string `json:"actId,omitempty"`
	// SkuViewID 商品id，对商品查询接口返回的skuViewid（和活动物料ID、活动链接三选一，不能全填）
	SkuViewID string `json:"skuViewId,omitempty"`
	// SID 二级媒体身份标识，用于渠道效果追踪，限制64个字符，仅支持英文、数字和下划线
	SID string `json:"sid,omitempty"`
	// LinkType 链接类型，枚举值：1 H5长链接；2 H5短链接；3 deeplink(唤起)链接；4 微信小程序唤起路径；5 团口令
	LinkType int `json:"linkType,omitempty"`
	// Text 只支持到家外卖商品券、买菜业务类型链接和活动物料链接。活动链接，即想要推广的目标链接，出参会返回成自己可推的链接，限定为当前可推广的活动链接或者商品券链接，请求内容尽量保持在200字以内，文本中仅存在一个http协议头的链接
	Text string `json:"text,omitempty"`
}

func (r GetReferralLinkRequest) Encode(w io.Writer) {
	json.NewEncoder(w).Encode(r)
}

// GetReferralLinkResponse 获取推广链接接口 API Response
type GetReferralLinkResponse struct {
	BaseResponse
	// SkuViewID 若用text进行入参取链，且返回的推广链接为商品券链接，则返回对应商品的展示ID，可以根据该ID查商品券接口获取对应的展示信息和佣金信息
	SkuViewID string `json:"skuViewId,omitempty"`
	// Data 返回对应的推广链接，这里的链接才能实现跟单计佣
	Data string `json:"data,omitempty"`
}
