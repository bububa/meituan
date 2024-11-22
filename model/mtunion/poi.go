package mtunion

import (
	"net/url"
	"strconv"
	"time"

	"github.com/bububa/meituan/enum"
	"github.com/bububa/meituan/model"
)

// PoiRequest 门店POI查询（新版） API Request
type PoiRequest struct {
	// Sid 推广位sid，支持通过接口自定义创建，不受平台200个上限限制，长度不能超过64个字符，支持小写字母和数字，历史已创建的推广位不受这个约束
	Sid string `json:"sid,omitempty"`
	// BusinessLine 固定2=外卖
	BusinessLine enum.BusinessLine `json:"business_line,omitempty"`
	// Longtitude 火星坐标系，经度，
	// 举例：120212010，
	// 其中120为整数部分，212010为小数部分，详见文档末尾数据示例
	Longtitude int64 `json:"longtitude,omitempty"`
	// Latitude 火星坐标系，纬度，
	// 举例：30208400，
	// 其中30为整数部分，208400为小数部分，详见文档末尾数据示例
	Latitude int64 `json:"latitude,omitempty"`
	// CateID 一级品类ID，参考文档链接：单店API外卖品类和筛选项说明
	CateID uint64 `json:"cateId,omitempty"`
	// SecondCateID 二级品类ID，参考文档链接：单店API外卖品类和筛选项说明
	SecondCateID uint64 `json:"secondCateId,omitempty"`
	// FilterConditionCode 筛选项列表，不筛选条件通过逗号分隔，筛选项目不超过6类，超过后将返回异常。
	// 当选中相似的筛选项时，如：1km内和2km内，返回的结果同时包含1km内和2km内的店铺。
	// 参考文档链接：单店API外卖品类和筛选项说明
	// 举例：71,81
	FilterConditionCode string `json:"filterConditionCode,omitempty"`
	// PageNo 页码
	PageNo int `json:"pageNo,omitempty"`
	// PageSize 每页数量
	PageSize int `json:"page_size,omitempty"`
	// PageTraceID 首次查询不传，后面根据第一次接口返回的值传，否则可能导致翻页失败。
	PageTraceID string `json:"pageTraceId,omitempty"`
	// Ts 秒时间戳
	Ts int64 `json:"ts,omitempty"`
}

func (r PoiRequest) Values(values url.Values) {
	values.Set("sid", r.Sid)
	values.Set("businessLine", strconv.Itoa(int(r.BusinessLine)))
	values.Set("longtitude", strconv.FormatInt(r.Longtitude, 10))
	values.Set("latitude", strconv.FormatInt(r.Latitude, 10))
	if r.CateID > 0 {
		values.Set("cateId", strconv.FormatUint(r.CateID, 10))
	}
	if r.SecondCateID > 0 {
		values.Set("secondCateId", strconv.FormatUint(r.SecondCateID, 10))
	}
	if r.FilterConditionCode != "" {
		values.Set("filterConditionCode", r.FilterConditionCode)
	}
	if r.PageNo == 0 {
		r.PageNo = 1
	}
	values.Set("pageNo", strconv.Itoa(r.PageNo))
	if r.PageSize == 0 {
		r.PageSize = 200
	}
	values.Set("pageSize", strconv.Itoa(r.PageSize))
	if r.Ts == 0 {
		r.Ts = time.Now().Unix()
	}
	values.Set("ts", strconv.FormatInt(r.Ts, 10))
}

// PoiResponse 门店POI查询（新版） API Response
type PoiResponse struct {
	Data *PoiResult `json:"data,omitempty"`
	model.BaseResponse
}

type PoiResult struct {
	// PageTraceID // 分页查询参数，第二次查询传回
	PageTraceID string `json:"pageTraceId,omitempty"`
	// DataList
	DataList []PoiInfo `json:"dataList,omitempty"`
}

// PoiInfo Poi门店信息
type PoiInfo struct {
	// PoiViewID POI门店ID
	PoiViewID model.Uint64 `json:"poiViewId,omitempty"`
	// PoiName POI名称
	PoiName string `json:"poiName,omitempty"`
	// PoiPicURL 店铺图URL
	PoiPicURL string `json:"poiPicUrl,omitempty"`
	// PoiScore 店铺评分，满分5分
	PoiScore model.Float64 `json:"poiScore,omitempty"`
	// MonthSale 月售量
	MonthSale model.Int64 `json:"monthSale,omitempty"`
	// ShippingFee 配送费金额，单位元
	ShippingFee model.Float64 `json:"shippingFee,omitempty"`
	// MinPrice 起送金额，单位元
	MinPrice model.Float64 `json:"minPrice,omitempty"`
	// Distance 门店距离，单位米
	Distance model.Float64 `json:"distance,omitempty"`
	// AvgDeliveryTime 配送时长，单位分钟
	AvgDeliveryTime model.Int64 `json:"avgDeliveryTime,omitempty"`
	// ReduceShippingFee 满减配送费
	ReduceShippingFee model.Float64 `json:"reduceShippingFee"`
	// PoiMarkTagUrl 角标信息
	PoiMarkTagUrl string `json:"poiMarkTagUrl,omitempty"`
	// MerchantFullSale 店铺满减,举例：38减25
	MerchantFullSale string `json:"merchantFullSale,omitempty"`
	// MerchantDiscount 店铺折扣，举例：3.4折起
	MerchantDiscount string `json:"merchantDiscount,omitempty"`
	// NewCustomerDiscount 新客立减，举例：新客减1
	NewCustomerDiscount string `json:"newCustomerDiscount,omitempty"`
	// RebateCoupon 返券，举例：返3元券
	RebateCoupon string `json:"rebateCoupon,omitempty"`
	// MerchantCoupon 商家券，举例：领3元券
	MerchantCoupon string `json:"merchantCoupon,omitempty"`
	// FullComplimentary 满赠，举例：满68元得赠品
	FullComplimentary string `json:"fullComplimentary,omitempty"`
}
