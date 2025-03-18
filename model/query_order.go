package model

import (
	"encoding/json"
	"io"
)

// QueryOrderRequest 查询订单接口 API Request
type QueryOrderRequest struct {
	// Platform 商品所属业务一级分类类型；请求的商品推广链接所属的业务类型信息，即只有输入skuViewId时才需要传本字段：1 到家及其他业务类型，2 到店业务类型；不填则默认1
	Platform int `json:"platform,omitempty"`
	// BusinessLine 业务线标识；1）当platform为1，选择到家及其他业务类型时，业务线枚举为1：外卖订单 WAIMAI 2：闪购红包 3：酒旅 4：美团电商订单（团好货） 5：医药 6：拼好饭 7：商品超值券包 COUPON 8：买菜 MAICAI 9：美团私域 11：闪购商品 12：省钱包；不传则默认传空表示非售卖券包、省钱包订单类型的全部查询。若输入参数含7 商品超值券包，则只返回商品超值券包订单，查询省钱包仅传省钱包参数值；2）当platform为2，选择到店业务类型 时，业务线枚举1：到餐 2：到综 3：酒店 4：门票 5：民宿 6：度假 15：到店红包，不填则默认传1, 999：新客CPA(到家和到店都包含)
	BusinessLine []int `json:"businessLine,omitempty"`
	// CategoryIDs 订单品类；1）当platform为1，当businessLine为11时，枚举值支持：大型连锁商超便利店(12)，小型商超便利店(14)，线上便利店(21)，日百服饰(128)，数码家电(106)，美妆日化(107)，母婴玩具(108)，宠物(110)，生鲜食材(24)，鲜花(16)，水果(15)，酒饮(26)，休闲食品(25)，旗舰店(137)，其他(-2)；2）当platform为1，当businessLine为9时，枚举值支持：进群(1)，下单(2)，首关注(3)；3）当platform为2，当businessLine为3时，枚举值支持：酒店(209)，非标住宿(2327)；4）当platform为2，当businessLine为2时，枚举值支持：休闲娱乐(3)，结婚(338)，教育培训(289)，养车/用车(390)，运动健身(206)，家居(600)，购物(379)，亲子(389)，医疗健康(450)，生活服务(4)，K歌(1853)，宠物(1861)，其他(-1) 4）当businessLine为999时，枚举值支持 到餐-纯新用户(1001),到餐-召回用户(1002),到综-纯新用户(1003),到综-召回用户(1004),闪购-纯新用户(1005),闪购-召回用户(1006),外卖-纯新用户(1007),外卖-沉默用户(1008),外卖-预警用户(1009),外卖-流失用户(1010)
	CategoryIDs []uint64 `json:"categoryIds,omitempty"`
	// ActID 活动物料id，我要推广-活动推广中第一列的id信息，不传则返回所有actId的数据，省钱包订单不传
	ActID string `json:"actId,omitempty"`
	// SID 二级媒体身份标识，用于渠道效果追踪，限制64个字符，仅支持英文、数字和下划线
	SID string `json:"sid,omitempty"`
	// OrderID 订单id，入参后可与业务线标识businessLine配合使用，输入的orderId需要与businessLine能对应上。举例：如查询商品超值券包订单时orderId传券包订单号，businessLine传7；除此以外其他查询筛选条件不生效，不传业务线标识businessLine则默认仅查非券包订单
	OrderID string `json:"orderId,omitempty"`
	// StartTime 查询时间类型对应的查询开始时间，10位时间戳表示，单位秒
	StartTime int64 `json:"startTime,omitempty"`
	// EndTime 查询时间类型对应的查询结束时间，10位时间戳表示，单位秒
	EndTime int64 `json:"endTime,omitempty"`
	// Page 页码，默认1，从1开始,若searchType选择2，本字段必须传1，若不传参数默认1
	Page int `json:"page,omitempty"`
	// Limit 每页限制条数，默认100，最大支持100
	Limit int `json:"limit,omitempty"`
	// QueryTimeType 查询时间类型，枚举值， 1 按订单支付时间查询， 2 按照更新时间查询， 默认为1
	QueryTimeType int `json:"query_time_type,omitempty"`
	// TradeType 交易类型，1表示CPS，2表示CPA
	TradeType int `json:"trade_type,omitempty"`
	// ScrollID 分页id，当searchType选择2逐页查询时，本字段为必填。若不填写，默认查询首页。取值为上一页查询时出参的scrollId字段
	ScrollID string `json:"scrollId,omitempty"`
	// SearchType 订单分页查询方案选择，不填则默认为1。1 分页查询（最多能查询到1万条订单），当选择本查询方案，page参数不能为空。此查询方式后续不再维护，建议使用2逐页查询。2 逐页查询（不限制查询订单数，只能逐页查询，不能指定页数），当选择本查询方案，需配合scrollId参数使用，省钱包查询仅支持2
	SearchType int `json:"searchType,omitempty"`
	// CityNames 可输入城市名称圈定特定城市的订单，单次最多查询10个城市（英文逗号分隔）。不传则默认全部城市订单。 注：如需确认城市具体名称，可参考后台订单明细页的城市筛选项，或参考具体活动的城市命名。目前支持到家业务类型-商品超值券包业务线；到店业务类型-到餐、到综、酒店、门票、民宿、度假业务线。
	CityNames []string `json:"cityNames,omitempty"`
}

func (r QueryOrderRequest) Encode(w io.Writer) {
	json.NewEncoder(w).Encode(r)
}

// QueryOrderResponse 查询订单接口 API Response
type QueryOrderResponse struct {
	BaseResponse
	// Data 响应结果信息
	Data *QueryOrderResult `json:"data,omitempty"`
}

type QueryOrderResult struct {
	// ActID 活动物料id，我要推广-活动推广中第一列的id信息
	ActID Uint64 `json:"actId,omitempty"`
	// SkuCount 查询返回本页的数量合计（无实际使用场景，若查询订单购买商品数可以看返回的dataList中skuCount）
	SkuCount int64 `json:"skuCount,omitempty"`
	// ScrollID 分页id，当searchType选择2逐页查询时，出参会返回本字段。用于下一页查询的scrollId字段入参使用
	ScrollID string `json:"scrollId,omitempty"`
	// DataList 数据列表
	DataList []Order `json:"dataList,omitempty"`
}

type Order struct {
	// BusinessLine 业务线，同入参枚举说明
	BusinessLine int `json:"businessLine,omitempty"`
	// OrderID 订单ID
	OrderID string `json:"orderId,omitempty"`
	// PayTime 订单支付时间
	PayTime int64 `json:"payTime,omitempty"`
	// PayPrice 订单支付价格。针对到餐、到综、酒店、闪购、医药业务类型，为父订单的支付价格，单位元
	PayPrice Float64 `json:"payPrice,omitempty"`
	// UpdateTime 订单最近一次的更新时间。到家外卖商品券、到家医药、到家闪购商品业务、到店到餐、到综、酒店类型，订单时间为用户买券包的更新时间，非每张券的更新时间。针对以上业务类型，建议查询单张券的更新时间
	UpdateTime int64 `json:"update_time,omitempty"`
	// CommissionRate 订单预估佣金比例，300表示3%
	CommissionRate Int `json:"commissionRate,omitempty"`
	// Profit cps类型的预估佣金收入，单位元，1.60表示1.6元
	Profit Float64 `json:"profit,omitempty"`
	// CpaProfit cpa类型的预估佣金收入，单位元，6.50表示6.5元
	CpaProfit Float64 `json:"cpaProfit,omitempty"`
	// SID 二级媒体身份标识，用于渠道效果追踪
	SID string `json:"sid,omitempty"`
	// ProductID 产品ID，对应商品查询接口的skuViewId，目前只支持到家外卖商品券、到家医药、到家闪购商品业务、到店业务类型
	ProductID string `json:"productId,omitempty"`
	// ProductName 产品名称，外卖订单展示店铺名称，到店取单个商品券的名称、其他展示全部商品名称
	ProductName string `json:"productName,omitempty"`
	// SpecificationName 规格信息，同一个商品名称下可以包括不同的规格，对应不同的价格和佣金
	SpecificationName string `json:"specificationName,omitempty"`
	// RefundPrice 只对非到店到餐、非到综、非酒店业务类型有效。订单维度退款价格，该笔订单用户发生退款行为时的退款计佣金额之和，超值券包订单本期不返回退款数据，单位元
	RefundPrice Float64 `json:"refundPrice,omitempty"`
	// RefundTime 只对非到店到餐、非到综、非酒店业务类型有效。订单维度最新一次发生退款的时间；超值券包订单本期不返回退款数据，单位元
	RefundTime int64 `json:"refundTime,omitempty"`
	// RefundProfit 只对非到店到餐、非到综、非酒店业务类型有效。订单维度退款预估佣金，该笔订单用户发生退款行为时的退款预估佣金金额之和；超值券包订单本期不返回退款数据，单位元
	RefundProfit Float64 `json:"refundProfit,omitempty"`
	// CpaRefundProfit cpa退款预估佣金，单位元
	CpaRefundProfit Float64 `json:"cpaRefundProfit,omitempty"`
	// Status 表示订单维度状态，枚举有 2：付款（如果是CPA订单则表示奖励已创建） 3：完成 4：取消 5：风控 6：结算。 针对到家商品券订单、到家闪购订单、到家医药订单、到店到餐、到综、酒店业务类型订单则为父订单相关状态，枚举有2：付款，父订单仅付款，至少有任意一个子订单未核销； 3：完成，父订单中所有子订单都核销完成； 4：取消，父订单中的子订单全部退款或过期未使用； 5：风控，父订单中的子订单全部变成风控状态； 6：结算，父订单中所有子订单都结算完成。（CPA订单只有到家闪购订单、到家医药订单、到店到餐、到综业务类型有本状态） 说明： 1、若到店到餐、到综、酒店业务类型父订单、到家闪购商品父订单、到家医药父订单，含有多个状态混合的子订单，则随机取子订单状态作为父订单状态，建议以orderDetail中每张券状态为准 2、含多个商品或券包的订单不建议使用该字段，实际计佣状态以orderDetail中每张券的计佣状态为准。
	Status Int `json:"status,omitempty"`
	// TradeType 交易类型，1：cps，2：cpa
	TradeType int `json:"tradeType,omitempty"`
	// ActID 活动物料id，我要推广-活动推广中第一列的id信息
	ActID Uint64 `json:"actId,omitempty"`
	// AppKey 归因到的appKey，对应取链时入参的appkey
	AppKey string `json:"appKey,omitempty"`
	// SkuCount 表示sku数量，团好货和券包类型的CPS订单返回有值，其余类型订单不返回该值
	SkuCount int64 `json:"skuCount,omitempty"`
	// CityName 订单所属的城市，目前支持三级城市粒度。目前支持到家业务类型-商品超值券包业务线；到店业务类型-到餐、到综、酒店、门票、民宿、度假业务线。
	CityName string `json:"cityName,omitempty"`
	// CategoryID 订单品类id。
	CategoryID uint64 `json:"categoryId,omitempty"`
	// CategoryName 订单品类名称。
	CategoryName string `json:"categoryName,omitempty"`
	// OrderDetail 订单详情，只支持到家外卖商品券、到家医药、到家闪购商品业务、到店到餐、到综、酒店类型返回数据
	OrderDetail *OrderDetail `json:"orderDetail,omitempty"`
}

// OrderDetail 订单详情，只支持到家外卖商品券、到家医药、到家闪购商品业务、到店到餐、到综、酒店类型返回数据
type OrderDetail struct {
	// CouponStatus 本期只有到到家外卖商品券、到家医药、到家闪购商品业务、到店到餐、到综、酒店业务类型展示订单明细，表示商品券/子订单推广计佣状态，1、付款，2、完成（或券已核销），3、结算，4、失效（含取消或风控的情况）
	CouponStatus Int `json:"couponStatus,omitempty"`
	// ItemOrderID 针对到店到餐、到综、酒店商品券，返回商品券的子订单号。其他业务类型不返回
	ItemOrderID string `json:"itemOrderId,omitempty"`
	// FinishTime 1、针对到家外卖商品券，返回商品券核销完成履约的实物菜品订单号对应的完成时间；2、针对到家医药&闪购商品，返回商品订单完成时间；3、针对到店到餐、到综、酒店子订单，返回子订单对应的券核销时间
	FinishTime Int64 `json:"finishTime,omitempty"`
	// BasicAmount 商品的计佣金额，每个商品对应的支付分摊金额，单位元
	BasicAmount Float64 `json:"basicAmount,omitempty"`
	// CouponFee 商品的佣金，当推广状态为失效、取消、风控时，佣金值为0，单位元
	CouponFee Float64 `json:"couponFee,omitempty"`
	// OrderViewID 只对到家外卖商品券有效。商品券的核销完成履约的实物菜品订单号
	OrderViewID string `json:"orderViewId,omitempty"`
	// RefundAmount 到店到餐、到综、酒店子订单、到家闪购商品、到家医药业务类型的退款金额，到家其他业务类型不返回数据，单位元
	RefundAmount Float64 `json:"refundAmount,omitempty"`
	// RefundFee 到店到餐、到综、酒店子订单、到家闪购商品、到家医药业务类型的退款佣金，到家其他业务类型不返回数据，单位元
	RefundFee Float64 `json:"refundFee,omitempty"`
	// RefundTime 到店到餐、到综、酒店子订单、到家闪购商品、到家医药业务类型的退款时间，到家其他业务类型不返回数据
	RefundTime Int64 `json:"refundTime,omitempty"`
	// SettleTime 到家商品券/到家闪购商品/到店到餐/到综/酒店子订单的结算时间，完成并且进入结算账期时则变为结算状态。若存在多次结算记录则取最新结算时间
	SettleTime Int64 `json:"settleTime,omitempty"`
	// UpdateTime 到家商品券/到家闪购商品/到家医药/到店到餐、到综、酒店子订单的更新时间
	UpdateTime Int64 `json:"updateTime,omitempty"`
}
