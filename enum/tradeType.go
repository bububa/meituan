package enum

// TradeType 订单的奖励类型
// 需要等待订单在完成或者核销状态，该字段可查询
// 查询时外卖，闪购，美团电商，优选存在该字段，酒店业务为空，字段书写方式[X,X]，例如[2,3]，含义为点击归因，且为首单，特殊的外卖CPA类订单，T+8可查询。[2,8]表示外卖社群新客关注奖励订单
type TradeType int

const (
	// TradeType_COUPON 券归因
	TradeType_COUPON TradeType = 1
	// TradeType_CLICK 点击归因
	TradeType_CLICK TradeType = 2
	// TradeType_FIRST_PURCHASE 3 首购奖励
	TradeType_FIRST_PURCHASE TradeType = 3
	// TradeType_REPURCHASE 复购
	TradeType_REPURCHASE TradeType = 4
	// TradeType_RETAIN 5 留存奖励
	TradeType_RETAIN TradeType = 5
	// TradeType_SECOND_PURCHASE 二单奖励
	TradeType_SECOND_PURCHASE TradeType = 6
	// TradeType_RECALL 唤起
	TradeType_RECALL TradeType = 7
	// TradeType_FOLLOW 关注
	TradeType_FOLLOW TradeType = 8
)
