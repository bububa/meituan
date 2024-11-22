# 美团联盟及CID接口 Golang SDK

[![Go Reference](https://pkg.go.dev/badge/github.com/bububa/meituan.svg)](https://pkg.go.dev/github.com/bububa/meituan)
[![Go](https://github.com/bububa/meituan/actions/workflows/go.yml/badge.svg)](https://github.com/bububa/meituan/actions/workflows/go.yml)
[![goreleaser](https://github.com/bububa/meituan/actions/workflows/goreleaser.yml/badge.svg)](https://github.com/bububa/meituan/actions/workflows/goreleaser.yml)
[![GitHub go.mod Go version of a Go module](https://img.shields.io/github/go-mod/go-version/bububa/meituan.svg)](https://github.com/bububa/meituan)
[![GoReportCard](https://goreportcard.com/badge/github.com/bububa/meituan)](https://goreportcard.com/report/github.com/bububa/meituan)
[![GitHub license](https://img.shields.io/github/license/bububa/meituan.svg)](https://github.com/bububa/meituan/blob/master/LICENSE)
[![GitHub release](https://img.shields.io/github/release/bububa/meituan.svg)](https://GitHub.com/bububa/meituan/releases/)

- 自助取链接口 [ GenerateLink(ctx context.Context, clt *core.SDKClient, req *model.GenerateLinkRequest) (string, error) ]
- CID取链 [ CidGenerateLink(ctx context.Context, clt *core.SDKClient, req *model.CidGenerateLinkRequest) (*model.CidLinkInfo, error) ]
- 订单列表查询接口 [ OrderList(ctx context.Context, clt *core.SDKClient, req *model.OrderListRequest) (*model.OrderListResponse, error) ]
- 单订单查询接口 [ Order(ctx context.Context, clt *core.SDKClient, req *model.OrderRequest) (*model.Order, error) ]
- 小程序生成二维码 [ MiniCode(ctx context.Context, clt *core.SDKClient, req *model.MiniCodeRequest) (string, error) ]
- 门店POI查询 [ mtunion.Poi(ctx context.Context, clt *core.SDKClient, req *mtunion.PoiRequest) (*mtunion.PoiResult, error) ]
- 券包商品查询接口 [ mtunion.Sku(ctx context.Context, clt *core.SDKClient, req *mtunion.SkuRequest) (*mtunion.SkuResult, error) ]
- 订单回推接口 [ serve.Handler(ctx context.Context, w http.ResponseWriter, r *http.Request) (*serve.Order, error) ]

