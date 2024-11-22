package serve

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/bububa/meituan/model/serve"
)

// Handler 订单回推接口（新版）
// https://union.meituan.com/v2/apiDetail?id=22
func Handler(ctx context.Context, w http.ResponseWriter, r *http.Request) (*serve.Order, error) {
	var o serve.Order
	if err := json.NewDecoder(r.Body).Decode(&o); err != nil {
		json.NewEncoder(w).Encode(serve.ErrorResponse)
		return nil, err
	}
	json.NewEncoder(w).Encode(serve.SuccessResponse)
	return &o, nil
}
