package mtunion

import (
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/bububa/meituan/enum"
	"github.com/bububa/meituan/model"
)

// SkuRequest 券包商品查询接口 API Request
type SkuRequest struct {
	// Uid 渠道UID
	Uid string `json:"uid,omitempty"`
	// Ts 时间戳，秒时间戳，有效期加减5分钟
	Ts int64 `json:"ts,omitempty"`
	// ActID 活动ID，商品券活动传XX
	ActID uint64 `json:"actId,omitempty"`
	// BusinessLine 业务类型，2-外卖
	BusinessLine enum.BusinessLine `json:"businessLine,omitempty"`
	// Longtitude 定位经度，需要获取商品券对应的【可配送】商家信息时，传入该值，不支持按定位经纬度筛选过滤商品
	Longtitude float64 `json:"longtitude,omitempty"`
	// Latitude 定位纬度，需要获取商品券对应的【可配送】商家信息时，传入该值，不支持按定位经纬度筛选过滤商品
	Latitude float64 `json:"latitude,omitempty"`
	// SkuIDList 商品券包ID列表，长度上限20个。传skuid查询性能较列表查询高。
	// 注：当传入skuIdList时，会忽略经纬度，筛选条件，排序等信息，进行精确查询，直接返回对应商品券包信息。
	SkuIDList []uint64 `json:"skuIdList,omitempty"`
	// PriceMaximum 价格上限，单位分
	PriceMaximum int64 `json:"priceMaximum,omitempty"`
	// PriceMinimum 价格下限，单位分
	PriceMinimum int64 `json:"priceMinimum,omitempty"`
	// SortField 排序字段
	// sales_volumn 销量，selling_price 售价
	SortField string `json:"sortField,omitempty"`
	// SortOrder 排序顺序
	// asc-升序，desc-降序
	SortOrder enum.SortOrder `json:"sortOrder,omitempty"`
	// PageNo 页码
	PageNo int `json:"pageNo,omitempty"`
	// PageSize 每页数量
	PageSize int `json:"pageSize,omitempty"`
}

func (r SkuRequest) Values(values url.Values) {
	values.Set("uid", r.Uid)
	if r.Ts == 0 {
		r.Ts = time.Now().Unix()
	}
	values.Set("ts", strconv.FormatInt(r.Ts, 10))
	values.Set("actId", strconv.FormatUint(r.ActID, 10))
	values.Set("businessLine", strconv.Itoa(int(r.BusinessLine)))
	if r.Longtitude > 1e-15 || r.Longtitude < 1e-15 {
		values.Set("longtitude", strconv.FormatFloat(r.Longtitude, 'f', -1, 64))
	}
	if r.Latitude > 1e-15 || r.Latitude < 1e-15 {
		values.Set("latitude", strconv.FormatFloat(r.Latitude, 'f', -1, 64))
	}
	if len(r.SkuIDList) > 0 {
		ids := make([]string, 0, len(r.SkuIDList))
		for _, v := range r.SkuIDList {
			ids = append(ids, strconv.FormatUint(v, 10))
		}
		values.Set("skuIdList", strings.Join(ids, ","))
	}
	if r.PriceMaximum > 0 {
		values.Set("priceMaximum", strconv.FormatInt(r.PriceMaximum, 10))
	}
	if r.PriceMinimum > 0 {
		values.Set("priceMinimum", strconv.FormatInt(r.PriceMinimum, 10))
	}
	if r.SortField != "" {
		values.Set("sortField", r.SortField)
	}
	if r.SortOrder != "" {
		values.Set("sortOrder", string(r.SortOrder))
	}
	if r.PageNo <= 0 {
		r.PageNo = 1
	}
	values.Set("pageNo", strconv.Itoa(r.PageNo))
	if r.PageSize <= 0 {
		r.PageSize = 20
	}
	values.Set("pageSize", strconv.Itoa(r.PageSize))
}

// SkuResponse 券包商品查询接口 API Response
type SkuResponse struct {
	// Data 商品券包信息
	Data *SkuResult `json:"data,omitempty"`
	model.BaseResponse
}

// SkuResult 商品券包详细信息
type SkuResult struct {
	// PageInfo 分页信息
	PageInfo *model.PageInfo `json:"pageInfo,omitempty"`
	// SkuInfoList sku信息列表
	SkuInfoList []SkuInfo `json:"skuInfoList,omitempty"`
}

// SkuInfo sku信息
type SkuInfo struct {
	// SkuID  sku编号
	SkuID model.Uint64 `json:"skuId,omitempty"`
	// SkuName sku名称
	SkuName string `json:"skuName,omitempty"`
	// CouponNum 商品券单个券包购买张数
	CouponNum int `json:"couponNum,omitempty"`
	// Pic 商品主图
	Pic string `json:"pic,omitempty"`
	// SalesVolume 当前sku销量
	SalesVolume string `json:"salesVolume,omitempty"`
	// RemainingStock 剩余库存，例：100+，1000+，10000+
	RemainingStock string `json:"remainingStock,omitempty"`
	// SaleStatus 售卖状态，可售为true，不可售为false或者不展示该商品券信息
	SaleStatus bool `json:"saleStatus,omitempty"`
	// OriginalPrice 活动原价，单位分
	OriginalPrice int64 `json:"originalPrice,omitempty"`
	// SellingPrice 活动售价，单位分
	SellingPrice int64 `json:"sellingPrice,omitempty"`
	// BrandName 品牌名称，无品牌则不展示
	BrandName string `json:"brandName,omitempty"`
	// BrandLogoURL 品牌logo，无品牌则不展示
	BrandLogoURL string `json:"brandLogoUrl,omitempty"`
	// AvailablePoiNum 可用门店数，表示该商品券适用于多少个商家门店
	AvailablePoiNum int `json:"availablePoiNum,omitempty"`
	// CoverCitys 覆盖城市集合，表示该商品券可用门店所在城市的集合
	CoverCitys []string `json:"coverCitys,omitempty"`
	// CoverCityNum 覆盖城市数，
	// 表示该商品券可用门店所在城市去重统计数量
	CoverCityNum int `json:"coverCityNum,omitempty"`
	// CommissionPercent 佣金比例，100表示1.00%
	CommissionPercent int `json:"commissionPercent,omitempty"`
	// DailyPurchaseLimit 单日限购数量
	DailyPurchaseLimit int `json:"dailyPurchaseLimit,omitempty"`
	// TotalPurchaseLimit 活动期间限购数量
	TotalPurchaseLimit int `json:"totalPurchaseLimit,omitempty"`
	// ExpirationDate 商品有效期，分为两种描述文案。
	// 购买后x天有效：例如，自购买后90天内有效
	// 时间范围：例如，2023-04-01 00:00:00~2023-05-01 00:00:00
	ExpirationDate string `json:"expirationDate,omitempty"`
	// SkuAvailablePoiInfoList 商品券可配送门店信息，无则不返回
	// 注：当skuIdList为空，入参经纬度可展示附近配送门店，若无入参或无可配送商家则不展示该信息。后续会开放精确查询情况下的传入经纬度获取对应附近可配送门店的逻辑。
	SkuAvailablePoiInfoList []SkuAvailablePoiInfo `json:"skuAvailablePoiInfoList,omitempty"`
}

// SkuAvailablePoiInfo 商品券可配送门店信息
type SkuAvailablePoiInfo struct {
	// PoiName 可配送的商家名称，若无可配送门店则不展示
	PoiName string `json:"poiName,omitempty"`
	// PoiLogoURL 对应商家的门店logo链接，若无可配送门店则不展示
	PoiLogoURL string `json:"poiLogoUrl,omitempty"`
	// DeliveryDistance 对应可配送商家的配送距离，若无可配送门店则不展示，单位km
	DeliveryDistance model.Float64 `json:"deliveryDistance,omitempty"`
	// DeliveryCost 对应可配送商家的预估配送费，单位元
	DeliveryCost model.Float64 `json:"deliveryCost,omitempty"`
	// DeliveryDuration 对应可配送商家的预估配送时长，单位分钟
	DeliveryDuration model.Int64 `json:"deliveryDuration,omitempty"`
	// LastDeliveryFee 对应可配送商家的起送费，单位元
	LastDeliveryFee model.Float64 `json:"lastDeliveryFee,omitempty"`
}
