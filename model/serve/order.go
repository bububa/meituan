package serve

import (
	"github.com/bububa/meituan/enum"
	"github.com/bububa/meituan/model"
)

// Order 订单回推
type Order struct {
	// SmsTitle 订单标题
	SmsTitle string `json:"smstitle,omitempty"`
	// Quantity 订单数量
	Quantity model.Int64 `json:"quantity,omitempty"`
	// OrderID 订单id
	OrderID model.Uint64 `json:"orderid,omitempty"`
	// DealID 店铺id（部分存在）
	DealID model.Uint64 `json:"dealid,omitempty"`
	// PayTime 订单支付时间，10位时间戳
	PayTime model.Int64 `json:"paytime,omitempty"`
	// ActID 活动id，可以在联盟活动列表中查看获取
	ActID model.Uint64 `json:"actId,omitempty"`
	// BusinessLine 详见业务线类型
	BusinessLine enum.BusinessLine `json:"businessLine,omitempty"`
	// SubBusinessLine 子业务线
	SubBusinessLine enum.SubBusinessLine `json:"subBusinessLine,omitempty"`
	// OrderTime 下单时间，10位时间戳
	OrderTime model.Int64 `json:"ordertime,omitempty"`
	// Sid 媒体推广位sid
	Sid string `json:"sid,omitempty"`
	// AppKey 媒体名称，可在推广者备案-媒体管理中查询
	Appkey string `json:"appkey,omitempty"`
	// Uid  渠道id
	Uid string `json:"uid,omitempty"`
	// Status 订单状态，枚举值同订单查询接口返回定义
	Status enum.OrderStatus `json:"status,omitempty"`
	// Total 订单总金额
	Total model.Float64 `json:"total,omitempty"`
	// PayPrice 订单实际支付金额
	PayPrice model.Float64 `json:"payPrice,omitempty"`
	// ModTime 订单修改时间
	ModTime model.Int64 `json:"modTime,omitempty"`
	// ProductID 商品ID
	ProductID model.Uint64 `json:"productId,omitempty"`
	// ProductName 商品名称
	ProductName string `json:"productName,omitempty"`
	// Direct 订单实际支付金额
	Direct model.Float64 `json:"direct,omitempty"`
	// Ratio 订单返佣比例，cps活动的订单会返回该字段
	Ratio model.Float64 `json:"ratio,omitempty"`
	// Sign 订单签名字段，计算方法参见文档中签名(sign)生成逻辑
	Sign string `json:"sign,omitempty"`
	// TradeTypeList 优选订单类型返回该字段
	TradeTypeList []enum.TradeType `json:"tradeTypeList,omitempty"`
	// ConsumeType 核销类型
	// 核销类型
	// 0 未核销
	// 1 已核销
	// 当前对到到餐、到综业务生效，在退款状态下推送，标识订单内在退款下的，不同消费券的核销状态
	ConsumeType int `json:"consumeType,omitempty"`
	// RefundType 退款类型
	// 退款类型
	// ALL_REFUND(1, "全部退"),
	// PART_REFUND(2, "部分退")
	// RISK_REFUND(3, "风控")
	RefundType enum.RefundType `json:"refundType,omitempty"`
	// EncryptionVoucherId 消费券加密券ID
	EncryptionVoucherId string `json:"encryptionVoucherId,omitempty"`
	// CouponCode 到综业务券码信息标识当前订单的券码编号，用于区分同一笔订单中的不同券码
	// 外卖券包分销，券实体订单使用的券展示IDe
	CouponCode string `json:"couponCode,omitempty"`
}
