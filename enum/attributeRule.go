package enum

// AttributeRule 归因规则（标品还是单店）
type AttributeRule int

const (
	// AttributeRule_SPU 1.cid-医药-spu-支付
	AttributeRule_SPU AttributeRule = 1
	// AttributeRule_SHOP 2.cid-医药-单店-支付
	AttributeRule_SHOP AttributeRule = 2
)
