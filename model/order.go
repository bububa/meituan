package model

import (
	"net/url"
	"strconv"

	"github.com/bububa/meituan/enum"
)

// OrderRequest 单订单查询接口（新版） API Request
type OrderRequest struct {
	// ActID 活动id，可以在联盟活动列表中查看获取
	ActID uint64 `json:"actId,omitempty"`
	// BusinessLine 业务线
	// 与actId二者至少择其一
	BusinessLine enum.BusinessLine `json:"businessLine,omitempty"`
	// Full 是否返回完整订单信息(即是否包含返佣、退款信息)
	// 枚举值：
	// 0-非全量查询
	// 1-全量查询
	// 非全量查询：不包含佣金数据和退款数据的订单信息
	// 全量查询：包含佣金数据和退款数据的订单信息
	Full int `json:"full,omitempty"`
	// OrderID 订单ID
	OrderID uint64 `json:"orderId,omitempty"`
	// ProductID 商品ID
	// 使用说明：
	// 美团电商（团好货）业务在单订单查询时，需要增加 productID 参数，为空则返回提示“团好货单订单查询需要productID参数”
	// 到餐业务在单订单查询时，需要增加productID参数，不传则返回为空值
	// 优选业务在单订单查询时，需要增加productID参数，不传则返回为空值
	// CID必须设置商品ID信息
	ProductID uint64 `json:"productId,omitempty"`
}

func (r OrderRequest) Values(values url.Values) {
	values.Set("actId", strconv.FormatUint(r.ActID, 10))
	if r.BusinessLine > 0 {
		values.Set("businessLine", strconv.Itoa(int(r.BusinessLine)))
	}
	if r.Full > 0 {
		values.Set("full", strconv.Itoa(r.Full))
	}
	values.Set("orderId", strconv.FormatUint(r.OrderID, 10))
	values.Set("productId", strconv.FormatUint(r.ProductID, 10))
}

// OrderResponse 单订单查询接口（新版） API Response
type OrderResponse struct {
	BaseResponse
	Data *Order `json:"data,omitempty"`
}

// Order 美团订单
type Order struct {
	// ActID 活动id，可以在联盟活动列表中查看获取
	ActID uint64 `json:"actId,omitempty"`
	// BusinessLine 业务线
	BusinessLine enum.BusinessLine `json:"businessLine,omitempty"`
	// SubBusinessLine 子业务线
	SubBusinessLine enum.SubBusinessLine `json:"subBusinessLine,omitempty"`
	// Quantity 商品数量
	Quantity int `json:"quantity,omitempty"`
	// ModTime 订单信息修改时间，10位时间戳
	ModTime Int64 `json:"modTime,omitempty"`
	// OrderID  订单id
	OrderID Uint64 `json:"orderid,omitempty"`
	// 订单支付时间，10位时间戳
	PayTime Int64 `json:"paytime,omitempty"`
	// PayPrice 订单用户实际支付金额
	PayPrice Float64 `json:"payprice,omitempty"`
	// Profit 订单预估返佣金额
	Profit Float64 `json:"profit,omitempty"`
	// CpaProfit 订单预估cpa总收益（优选、话费券）
	CpaProfit Float64 `json:"cpaProfit,omitempty"`
	// Sid  订单对应的推广位sid
	Sid string `json:"sid,omitempty"`
	// 订单对应的appkey，外卖、话费、闪购、优选订单会返回该字段
	AppKey string `json:"appkey,omitempty"`
	// SmsTitle 订单标题
	SmsTitle string `json:"smstitle,omitempty"`
	// ItemID 店铺ID
	ItemID Uint64 `json:"itemId,omitempty"`
	// ProductID 商品ID
	ProductID Uint64 `json:"productId,omitempty"`
	// ProductName 商品名称
	ProductName string `json:"productName,omitempty"`
	// RefundPrice 订单实际退款金额，外卖、话费、闪购、优选、酒店订单若发生退款会返回该字段
	RefundPrice Float64 `json:"refundprice,omitempty"`
	// RefundTime 订单退款时间，10位时间戳，外卖、话费、闪购、优选、酒店订单若发生退款会返回该字段(退款时间为最近的一次退款)
	RefundTime Int64 `json:"refundtime,omitempty"`
	// RefundProfit 订单需要扣除的返佣金额，外卖、话费、闪购、优选、酒店订单若发生退款会返回该字段
	RefundProfit Float64 `json:"refundprofit,omitempty"`
	// CpaRefundProfit 订单需要扣除的cpa返佣金额（优选、话费券）
	CpaRefundProfit Float64 `json:"cpaRefundProfit,omitempty"`
	// Status 订单状态，外卖、话费、闪购、优选、酒店订单会返回该字段 1 已付款 8 已完成 9 已退款或风控
	Status enum.OrderStatus `json:"status,omitempty"`
	// TradeTypeList 订单的奖励类型 3 首购奖励 5 留存奖励 2 cps 3 首购奖励
	TradeTypeList []enum.TradeType `json:"tradeTypeList,omitempty"`
	// RiskOrder 0表示非风控订单，1表示风控订单
	RiskOrder int `json:"riskOrder,omitempty"`
	// CouponCode 到综业务券码信息标识当前订单的券码编号，用于区分同一笔订单中的不同券码
	// 外卖券包分销，券实体订单使用的券展示IDe
	CouponCode string `json:"couponCode,omitempty"`
	// RefundInfoList 退款列表
	RefundInfoList []RefundInfo `json:"refundInfoList,omitempty"`
	// RefundProfitList 退款佣金明细
	RefundProfitList []RefundProfit `json:"refundProfitList,omitempty"`
	// ConsumeProfitList 核销佣金明细结构
	ConsumeProfitList []ConsumeProfit `json:"consumeProfitList,omitempty"`
	// CidInfo cid扩展信息
}

// RefundInfo 退款信息
type RefundInfo struct {
	// ID 退款单id
	ID Uint64 `json:"id,omitempty"`
	// RefundPrice 订单实际退款金额，外卖、话费、闪购、优选、酒店订单若发生退款会返回该字段
	RefundPrice Float64 `json:"refundPrice,omitempty"`
	// RefundTime 订单退款时间，10位时间戳
	RefundTime Int64 `json:"refundTime,omitempty"`
	// RefundType 退款类型
	RefundType enum.RefundType `json:"refundType,omitempty"`
}

// RefundProfit 退款佣金明细
type RefundProfit struct {
	// ID 退款单id
	ID Uint64 `json:"id,omitempty"`
	// RefundProfit 退款佣金
	RefundProfit Float64 `json:"refundProfit,omitempty"`
	// RefundFinishTime 佣金产生时间,10位时间戳
	RefundFinishTime Int64 `json:"refundFinishTime,omitempty"`
	// Type 查询时外卖，闪购，美团电商，优选存在该字段，酒店业务为空，字段书写方式[X,X]，例如[2,3]，含义为点击归因，且为首单
	// CPS订单：
	// 1-券归因
	// 2-点击归因
	// ____________________________
	// CPA订单：
	// 3-首单
	// 4-复购
	// 5-留存
	// 6-二单
	// 7-唤起
	Type enum.TradeType `json:"type,omitempty"`
}

// ConsumeProfit 核销佣金明细结构
type ConsumeProfit struct {
	// ID 核销id
	ID Uint64 `json:"id,omitempty"`
	// ConsumeProfit 核销佣金
	ConsumeProfit Float64 `json:"consumeProfit,omitempty"`
	// ConsumeFinishTime 佣金产生时间,10位时间戳
	ConsumeFinishTime Int64 `json:"consumeFinishTime,omitempty"`
	// Type
	Type enum.TradeType `json:"type,omitempty"`
}

// CidInfo cid扩展信息
type CidInfo struct {
	// CidMediaID 媒体信息：示例jinritoutiao
	CidMediaID string `json:"cidMediaId,omitempty"`
	// CidAccountID 账户id
	CidAccountID Uint64 `json:"cidAccountId,omitempty"`
	// CidAdGroupID 广告组id
	CidAdGroupID Uint64 `json:"cidAdGroupId,omitempty"`
	// CidPlanID 广告计划id
	CidPlanID Uint64 `json:"cidPlanId,omitempty"`
	// CidCreativeID 广告创意id
	CidCreativeID Uint64 `json:"cidCreativeId,omitempty"`
	// CidCallbackParam 回调参数
	CidCallbackParam string `json:"cidCallbackParam,omitempty"`
	// CidCallbackEvent 事件类型，当前可忽略
	CidCallbackEvent string `json:"cidCallbackEvent,omitempty"`
	// CidClickTimeStamp 点击时间戳
	CidClickTimeStamp int64 `json:"CidClickTimeStamp,omitempty"`
	// CidOsType 操作系统
	CidOsType string `json:"cidOsType,omitempty"`
	// CidAppName 下单app名称
	CidAppName string `json:"cidAppName,omitempty"`
	// CidRecipientID 收货城市id
	CidRecipientID Uint64 `json:"cidRecipientId,omitempty"`
	// CidRecipientName 收货城市name
	CidRecipientName string `json:"cidRecipientName,omitempty"`
}
