package model

import (
	"net/url"
	"strconv"

	"github.com/bububa/meituan/enum"
)

// MiniCodeRequest 小程序生成二维码 API Request
type MiniCodeRequest struct {
	// Sid 推广位sid，支持通过接口自定义创建，不受平台200个上限限制，长度不能超过64个字符，支持小写字母和数字，历史已创建的推广位不受这个约束
	Sid string `json:"sid,omitempty"`
	// ActID 活动id
	ActID uint64 `json:"actId,omitempty"`
	// LinkType WECHAT(4, "微信小程序-美团小程序"),
	// YOUXUAN_WXAPP(8, "微信小程序-优选小程序");
	LinkType enum.LinkType `json:"linkType,omitempty"`
	// CityID 城市ID
	CityID uint64 `json:"cityId,omitempty"`
	// SkuID 商品ID
	SkuID uint64 `json:"skuId,omitempty"`
	// CategoryID 商品类目ID
	CategoryID uint64 `json:"categoryId,omitempty"`
}

func (r MiniCodeRequest) Values(values url.Values) {
	values.Set("sid", r.Sid)
	values.Set("actId", strconv.FormatUint(r.ActID, 10))
	values.Set("linkType", strconv.Itoa(int(r.LinkType)))
	if r.CityID > 0 {
		values.Set("cityId", strconv.FormatUint(r.CityID, 10))
	}
	if r.SkuID > 0 {
		values.Set("skuId", strconv.FormatUint(r.SkuID, 10))
	}
	if r.CategoryID > 0 {
		values.Set("categoryId", strconv.FormatUint(r.CategoryID, 10))
	}
}

// MiniCodeResponse 小程序生成二维码 API Response
type MiniCodeResponse struct {
	Data string `json:"data,omitempty"`
	BaseResponse
}
