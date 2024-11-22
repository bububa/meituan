package serve

type Response struct {
	Errmsg  string `json:"errmsg"`
	Errcode int    `json:"errcode"`
}

var (
	// SuccessResponse 返回正常
	SuccessResponse = Response{"ok", 0}
	// Error 返回错误
	ErrorResponse = Response{"err", 1}
)
