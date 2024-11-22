package enum

import "strconv"

// OrderStatus 订单状态，外卖、话费、闪购、优选、酒店订单会返回该字段
type OrderStatus int

const (
	// OrderStatus_PAID 1 已付款
	OrderStatus_PAID OrderStatus = 1
	// OrderStatus_COMPLETED 8 已完成
	OrderStatus_COMPLETED OrderStatus = 8
	// OrderStatus_REFUNDED 9 已退款或风控
	OrderStatus_REFUNDED OrderStatus = 9
)

// UnmarshalJSON implement json Unmarshal interface
func (os *OrderStatus) UnmarshalJSON(b []byte) (err error) {
	if b[0] == '"' && b[len(b)-1] == '"' {
		b = b[1 : len(b)-1]
	}
	i, _ := strconv.Atoi(string(b))
	*os = OrderStatus(i)
	return
}

func (os OrderStatus) String() string {
	return strconv.Itoa(int(os))
}
