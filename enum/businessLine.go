package enum

import "strconv"

// BusinessLine 业务线
type BusinessLine int

const (
	// BusinessLine_PLATFORM 平台
	// 聚合活动等无法直接归属到某个具体的业务线的订单
	BusinessLine_PLATFORM BusinessLine = 1
	// BusinessLine_ECOM 到家_外卖&闪购&美团电商
	// 当业务线类型=2时，存在如下子业务分类：subBusinessLine=1-外卖订单；subBusinessLine=2-闪购订单；subBusinessLine=3-美团电商订单（团好货）
	// 外卖/闪购业务订单合并为类型2，如需区分订单类型可以通过活动ID区分或使用子业务线类型进行区分
	BusinessLine_ECOM BusinessLine = 2
	// BusinessLine_HOTEL 酒店/酒店券业务订单合并为类型3
	BusinessLine_HOTEL BusinessLine = 3
	// BusinessLine_SELECTED 优选业务订单
	BusinessLine_SELECTED BusinessLine = 4
	// BusinessLine_LOCAL 地推类业务订单
	BusinessLine_LOCAL BusinessLine = 5
	// BusinessLine_RESTAURANT 到餐类业务订单
	BusinessLine_RESTAURANT BusinessLine = 6
	// BusinessLine_OTHERS 到综类业务订单
	BusinessLine_OTHERS BusinessLine = 7
)

// SubBusinessLine 子业务线
type SubBusinessLine int

const (
	// SubBusinessLine_TAKEOUT 1-外卖订单
	SubBusinessLine_TAKEOUT SubBusinessLine = 1
	// SubBusinessLine_FLASH_SALE 2-闪购订单
	SubBusinessLine_FLASH_SALE SubBusinessLine = 2
	// SubBusinessLine_ECOM 3-美团电商订单（团好货）
	SubBusinessLine_ECOM SubBusinessLine = 3
	// SubBusinessLine_CHEAP_TAKEOUT 4-拼好饭
	SubBusinessLine_CHEAP_TAKEOUT SubBusinessLine = 4
	// SubBusinessLine_COUPON 5-虚拟券包
	SubBusinessLine_COUPON SubBusinessLine = 5
)

// UnmarshalJSON implement json Unmarshal interface
func (bl *BusinessLine) UnmarshalJSON(b []byte) (err error) {
	if b[0] == '"' && b[len(b)-1] == '"' {
		b = b[1 : len(b)-1]
	}
	i, _ := strconv.Atoi(string(b))
	*bl = BusinessLine(i)
	return
}

func (bl BusinessLine) String() string {
	return strconv.Itoa(int(bl))
}

// UnmarshalJSON implement json Unmarshal interface
func (bl *SubBusinessLine) UnmarshalJSON(b []byte) (err error) {
	if b[0] == '"' && b[len(b)-1] == '"' {
		b = b[1 : len(b)-1]
	}
	i, _ := strconv.Atoi(string(b))
	*bl = SubBusinessLine(i)
	return
}

func (bl SubBusinessLine) String() string {
	return strconv.Itoa(int(bl))
}
