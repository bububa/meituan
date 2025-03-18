package model

import (
	"encoding/json"
	"io"
)

// QueryCouponRequest 商品查询接口 API Request
type QueryCouponRequest struct {
	// Platform 商品所属业务一级分类类型：1 到家及其他业务类型，2 到店业务类型（包含到店美食、休闲生活、酒店、门票）；不填则默认1
	Platform int `json:"platform,omitempty"`
	// BizLine 商品所属业务二级分类类型；当字段platform为1，选择到家及其他业务类型时：5 医药 ，不填则默认为null，返回外卖商品券；当字段platform为2，选择到店业务类型时：1 到餐，2 到综 3：酒店 4：门票度假 不填则默认1
	BizLine int `json:"bizLine,omitempty"`
	// Longitude 定位经纬度的经度，请传递经度*100万倍的整形数字，如经度116.404*100万倍为116404000； 针对到店、到家医药商品业务类型，若未输入经纬度，则默认北京；针对到家外卖商品券业务类型，若未输入经纬度，则默认全国
	Longitude int64 `json:"longitude,omitempty"`
	// Latitude 定位经纬度的纬度，请传递纬度*100万倍的整形数字，如纬度39.928*100万倍为39928000; 针对到店、到家医药商品业务类型，若未输入经纬度，则默认北京；针对到家外卖商品券业务类型，若未输入经纬度，则默认全国
	Latitude int64 `json:"latitude,omitempty"`
	// PriceCap 筛选商品售卖价格上限【单位元】
	PriceCap int64 `json:"priceCap,omitempty"`
	// PriceFloor 筛选商品价格下限【单位元】
	PriceFloor int64 `json:"priceFloor,omitempty"`
	// CommissionCap 筛选商品佣金值上限【单位元】，若商品按照佣金值进行范围筛选，则排序只能按照佣金降序，本字段只支持到店业务类型、到家医药业务类型
	CommissionCap int64 `json:"commissionCap,omitempty"`
	// CommissionFloor 筛选商品佣金值下限【单位元】，若商品按照佣金值进行范围筛选，则排序只能按照佣金降序，本字段只支持到店业务类型、到家医药业务类型
	CommissionFloor int64 `json:"commissionFloor,omitempty"`
	// VpSkuViewIDs 商品ID集合，非必填，若填写该字段则不支持其他筛选条件，集合里ID用英文“,”隔开。一次最多支持查询20个售卖券ID
	VpSkuViewIDs string `json:"vpSkuViewIds,omitempty"`
	// ListTopiID 选品池榜单主题ID，到家及其他业务类型支持查询：1 精选，2 今日必推，3 同城热销，4 跟推爆款的商品售卖券（其中到家医药业务类型，本项为必填，且只支持传枚举3）；到店业务类型支持查询：2 今日必推，3 同城热销（全部商品）5 实时热销（其中到店酒店、门票业务类型，本项为必填，且只支持传枚举3）
	ListTopiID int `json:"listTopiId,omitempty"`
	// SearchText 搜索关键字,限制1-100个字符，不支持入参指定Platform、bizLine搜索，搜索范围为全品类。如需使用该字段查询商品信息，则vpSkuViewIds、listTopiId字段必须为空！！！如不为空，则按下述字段优先级执行查询：vpSkuViewIds>listTopiId>searchText。
	SearchText string `json:"searchText,omitempty"`
	// SearchID 仅搜索场景分页使用，首次调用不用填。查询相同搜索关键词、相同排序规则的下一页数据，需携带填写上次查询时出参中的'searchId'。如变更搜索关键字或排序规则，则也无需填写。
	SearchID string `json:"searchId,omitempty"`
	// PageSize 分页大小，不填返回默认分页20
	PageSize int `json:"pageSize,omitempty"`
	// PageNo 页数，不填返回默认页数1
	PageNo int `json:"pageNo,omitempty"`
	// SortField 1）未入参榜单listTopiId时：支持1 售价排序、2 销量排序；2）入参榜单listTopiId时：当platform为1，选择到家业务类型：外卖商品券类型，支持1 售价排序、 2 销量降序、 3佣金降序，不填则默认为1；到家医药类型，支持2 销量降序、 3 佣金降序，不填则默认为2； 当platform为2，选择到店业务类型：支持2 销量降序、 3佣金降序，不填则默认为2。其中listTopiId为5时，仅支持默认排序，sortField不生效；3)通过搜索searchText召回时：支持1综合排序、2价格升序，不填默认为1
	SortField int `json:"sortField,omitempty"`
	// AscDescOrder 仅对到家业务类型生效，未入参榜单listTopiId时：1 升序，2 降序； 入参榜单listTopiId时：1 升序，2 降序，并且仅对sortField为1售价排序的时候生效，其他筛选值不生效； 其他说明：不填则默认为1升序
	AscDescOrder int `json:"ascDescOrder,omitempty"`
}

func (r QueryCouponRequest) Encode(w io.Writer) {
	json.NewEncoder(w).Encode(r)
}

// QueryCouponResponse 商品查询接口 API Response
type QueryCouponResponse struct {
	BaseResponse
	// HasNext 分页使用，看是否有下一页
	HasNext bool `json:"hasNext,omitempty"`
	// SearchID 搜索场景出参,用于相同条件下一页请求入参
	SearchID string `json:"searchId,omitempty"`
	// Data 响应结果信息
	Data []QueryCouponResult `json:"data,omitempty"`
}

type QueryCouponResult struct {
	// AvailablePOIInfo 可用门店信息
	AvailablePOIInfo *AvailablePOIInfo `json:"availablePoiInfo,omitempty"`
	// BrandInfo 品牌信息
	BrandInfo *BrandInfo `json:"brandInfo,omitempty"`
	// CommissionInfo 佣金信息
	CommissionInfo *CommissionInfo `json:"commissionInfo,omitempty"`
	// CouponPackDetail 商品详情
	CouponPackDetail *CouponPackDetail `json:"couponPackDetail,omitempty"`
	// DeliverablePOIInfo 只支持到家外卖商品券业务类型，可配送门店信息
	DeliverablePOIInfo *DeliverablePOIInfo `json:"deliverablePoiInfo,omitempty"`
	// PurchaseLimitInfo 购买限制信息
	PurchaseLimitInfo *PurchaseLimitInfo `json:"purchaseLimitInfo,omitempty"`
	// CouponValidTimeInfo 只支持到家外卖商品券业务类型，券包活动有效时间信息
	CouponValidTimeInfo *CouponValidTimeInfo `json:"CouponValidTimeInfo,omitempty"`
}

// AvailablePOIInfo 可用门店信息
type AvailablePOIInfo struct {
	// AvailablePOINum 可用门店数量。针对到店、到家医药业务类型商品，若传入经纬度信息，则为经纬度所在城市可用的门店数。若不传入经纬度信息，则输出北京可用的门店数
	AvailablePOINum int64 `json:"availablePoiNum,omitempty"`
}

// BrandInfo 品牌信息
type BrandInfo struct {
	// BrandName 品牌名称
	BrandName string `json:"brandName,omitempty"`
	// BrandLogoURL 品牌Logo的url
	BrandLogoURL string `json:"brandLogoUrl,omitempty"`
}

// CommissionInfo 佣金信息
type CommissionInfo struct {
	// CommissionPercent 查询当时生效的佣金比例， 商品券拉取、通过商品券ID查询、通过榜单listTopiId查询，返回的数据需要除以100表示对应的佣金比例，如返回400表示佣金比例为4%
	CommissionPercent Int `json:"commissionPercent,omitempty"`
	// Commission 只支持到店、到家医药业务类型。查询当时生效的佣金值。单位元，保留小数点后两位
	Commission Float64 `json:"commission,omitempty"`
}

// CouponPackDetail 商品详情
type CouponPackDetail struct {
	// Name 商品名称
	Name string `json:"name,omitempty"`
	// SkuViewID 商品skuViewId，传入开放平台取链接口的skuViewId，取得对应推广链接才能正常归因订单
	SkuViewID string `json:"skuViewId,omitempty"`
	// Specification 规格信息，只支持到家医药商品业务类型
	Specification string `json:"specification,omitempty"`
	// CouponNum 只支持到家外卖商品券业务类型，券包中券的数量
	CouponNum int `json:"couponNum,omitempty"`
	// ValidTime 只支持到家外卖商品券业务类型，活动截止有效日期，仅作参考，具体结束时间详见couponValidTimeInfo中的信息a
	ValidTime int64 `json:"validTime,omitempty"`
	// HeadURL 商品头图的url                  a
	HeadURL string `json:"headUrl,omitempty"`
	// SaleVolume 美团累计销量，例：100+，1000+，10000+
	SaleVolume string `json:"saleVolume,omitempty"`
	// StartTime 只支持到家外卖商品券业务类型，活动有效期开始时间
	StartTime int64 `json:"startTime,omitempty"`
	// EndTime 只支持到家外卖商品券业务类型，活动有效期结束时间
	EndTime int64 `json:"endTime,omitempty"`
	// SaleStatus 售卖状态，可售为是，不可售为否。不可售商品不返回商品数据
	SaleStatus bool `json:"saleStatus,omitempty"`
	// OriginalPrice 原始价格，如划线价(元）
	OriginalPrice Float64 `json:"originalPrlice,omitempty"`
	// SellPrice 售卖价格(元）
	SellPrice Float64 `json:"sellPrice,omitempty"`
	// Platform 平台，1-到家、2-到店
	Platform int `json:"platform,omitempty"`
	// DeliverablePOIInfo 只支持到家外卖商品券业务类型，可配送门店信息
	DeliverablePOIInfo *DeliverablePOIInfo `json:"deliverablePoiInfo,omitempty"`
	// PurchaseLimitInfo 购买限制信息
	PurchaseLimitInfo *PurchaseLimitInfo `json:"purchaseLimitInfo,omitempty"`
	// CouponValidTimeInfo 只支持到家外卖商品券业务类型，券包活动有效时间信息
	CouponValidTimeInfo *CouponValidTimeInfo `json:"couponValidTimeInfo,omitempty"`
}

// DeliverablePOIInfo 只支持到家外卖商品券业务类型，可配送门店信息
type DeliverablePOIInfo struct {
	// POIName 门店名称，商品券可配送门店信息，无则不返回 注：入参经纬度可展示附近配送门店名称。按主题榜单查询时不展示该字段
	POIName string `json:"poiName,omitempty"`
	// POILogoURL 门店Logo的url 注：入参经纬度可展示附近配送门店logo。按主题榜单查询时不展示该字段。
	POILogoURL string `json:"poiLogoURL,omitempty"`
	// DeliveryDistance 配送距离 注：入参经纬度可展示附近配送门店的配送距离。按主题榜单查询时不展示该字段。
	DeliveryDistance Float64 `json:"deliveryDistance,omitempty"`
	// DistributionCost 配送费 注：入参经纬度可展示附近配送门店的配送费。按主题榜单查询时不展示该字段。
	DistributionCost Float64 `json:"distributionCost,omitempty"`
	// DeliveryDuration 配送时长 注：入参经纬度可展示附近配送门店的配送时长。按主题榜单查询时不展示该字段。
	DeliveryDuration Int64 `json:"deliveryDuration,omitempty"`
	// LastDeliveryFee 起送额 注：入参经纬度可展示附近配送门店的起送金额。按主题榜单查询时不展示该字段。
	LastDeliveryFee Float64 `json:"lastDeliveryFee,omitempty"`
}

// PurchaseLimitInfo 购买限制信息
type PurchaseLimitInfo struct {
	// SingleDayPurchaseLimit 单日售卖上限
	SingleDayPurchaseLimit int64 `json:"singleDayPurchaseLimit,omitempty"`
}

// CouponValidTimeInfo 只支持到家外卖商品券业务类型，券包活动有效时间信息
type CouponValidTimeInfo struct {
	// CouponValidTimeType 券包活动生效时间类型,1:按生效天数,2:按时间段
	CouponValidTimeType int `json:"couponValidTimeType,omitempty"`
	// CouponValidDay 券生效天数；couponValidTimeType为1有效
	CouponValidDay int `json:"couponValidDay,omitempty"`
	// CouponValidSTime 券开始时间戳，单位秒；couponValidTimeType为2有效
	CouponValidSTime int64 `json:"couponValidSTime,omitempty"`
	// CouponValidETime 券结束时间戳，单位秒；couponValidTimeType为2有效
	CouponValidETime int64 `json:"couponValidETime,omitempty"`
}
