package enum

// LinkType 链接类型，枚举值：
type LinkType int

const (
	// LinkType_H5 1 h5链接
	LinkType_H5 LinkType = 1
	// LinkType_DEEPLINK 2 deeplink(唤起)链接
	LinkType_DEEPLINK LinkType = 2
	// LinkType_JUMP_PAGE 3 中间页唤起链接
	LinkType_JUMP_PAGE LinkType = 3
	// LinkType_WECHAT 4 微信小程序唤起路径
	LinkType_WECHAT LinkType = 4
	// LinkType_COUPON_TOKEN 5 团口令
	LinkType_COUPON_TOKEN LinkType = 5
	// LinkType_YOUXUAN_WXAPP (8, "微信小程序-优选小程序")
	LinkType_YOUXUAN_WXAPP LinkType = 8
)
