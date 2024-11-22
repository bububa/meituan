package model

import (
	"net/url"
	"strconv"
	"time"

	"github.com/bububa/meituan/enum"
)

// OrderListRequest 订单列表查询接口（新版） API Request
type OrderListRequest struct {
	// ActID 活动id，可以在联盟活动列表中查看获取
	ActID uint64 `json:"actId,omitempty"`
	// Ts 请求时刻10位时间戳(秒级)，有效期60s
	Ts int64 `json:"ts,omitempty"`
	// BusinessLine 业务线
	// 与actId二者至少择其一
	BusinessLine enum.BusinessLine `json:"businessLine,omitempty"`
	// StartTime 查询起始时间10位时间戳，以下单时间为准
	// 特别提示：为了保障查询性能，最长查询时间为1天
	StartTime int64 `json:"startTime,omitempty"`
	// EndTime 查询截止时间10位时间戳，以下单时间为准
	// 特别提示：为了保障查询性能，最长查询时间为1天
	EndTime int64 `json:"endTime,omitempty"`
	// Page 分页参数，起始值从1开始
	Page int `json:"page,omitempty"`
	// Limit 每页显示数据条数，最大值为100
	Limit int `json:"limit,omitempty"`
	// QueryTimeType 查询时间类型，枚举值
	// 1 按订单支付时间查询
	QueryTimeType enum.QueryTimeType `json:"queryTimeType,omitempty"`
}

func (r OrderListRequest) Values(values url.Values) {
	if r.Ts == 0 {
		r.Ts = time.Now().Unix()
	}
	values.Set("ts", strconv.FormatInt(r.Ts, 10))
	if r.ActID > 0 {
		values.Set("actId", strconv.FormatUint(r.ActID, 10))
	}
	if r.BusinessLine > 0 {
		values.Set("businessLine", strconv.Itoa(int(r.BusinessLine)))
	}
	values.Set("startTime", strconv.FormatInt(r.StartTime, 10))
	values.Set("endTime", strconv.FormatInt(r.EndTime, 10))
	if r.Page == 0 {
		r.Page = 1
	}
	values.Set("page", strconv.Itoa(r.Page))
	if r.Limit == 0 {
		r.Limit = 100
	}
	values.Set("limit", strconv.Itoa(r.Limit))
	if r.QueryTimeType > 0 {
		values.Set("queryTimeType", strconv.Itoa(int(r.QueryTimeType)))
	}
}

// OrderListResponse 订单列表查询接口（新版） API Response
type OrderListResponse struct {
	// DataList 订单列表
	DataList []Order `json:"dataList,omitempty"`
	BaseResponse
	// Total 查询条件命中的总数据条数，用于计算分页参数
	Total int `json:"total,omitempty"`
}
