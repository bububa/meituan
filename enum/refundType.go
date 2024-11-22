package enum

import "strconv"

// RefundType 退款类型
type RefundType int

const (
	// ALL_REFUND(1, "全部退"),
	ALL_REFUND RefundType = 1
	// PART_REFUND(2, "部分退")
	PART_REFUND RefundType = 2
	// RISK_REFUND(3, "风控")
	RISK_REFUND RefundType = 3
)

// UnmarshalJSON implement json Unmarshal interface
func (rt *RefundType) UnmarshalJSON(b []byte) (err error) {
	if b[0] == '"' && b[len(b)-1] == '"' {
		b = b[1 : len(b)-1]
	}
	i, _ := strconv.Atoi(string(b))
	*rt = RefundType(i)
	return
}

func (rt RefundType) String() string {
	return strconv.Itoa(int(rt))
}
