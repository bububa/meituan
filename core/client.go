package core

import (
	"context"
	"errors"
	"net/http"
	"sync"
	"time"

	"github.com/bububa/meituan/core/internal/debug"
	"github.com/bububa/meituan/model"
	"github.com/bububa/meituan/util"
)

var (
	onceInit   sync.Once
	httpClient *http.Client
)

func defaultHttpClient() *http.Client {
	onceInit.Do(func() {
		transport := http.DefaultTransport.(*http.Transport).Clone()
		transport.MaxIdleConns = 100
		transport.MaxConnsPerHost = 100
		transport.MaxIdleConnsPerHost = 100
		httpClient = &http.Client{
			Transport: transport,
			Timeout:   time.Second * 60,
		}
	})
	return httpClient
}

// SDKClient sdk client
type SDKClient struct {
	client  *http.Client
	tracer  *Otel
	limiter RateLimiter
	appKey  string
	secret  string
	debug   bool
}

// NewSDKClient 创建SDKClient
func NewSDKClient(appKey string, secret string) *SDKClient {
	return &SDKClient{
		appKey: appKey,
		secret: secret,
		client: defaultHttpClient(),
	}
}

// SetDebug 设置debug模式
func (c *SDKClient) SetDebug(debug bool) {
	c.debug = debug
}

// SetHttpClient 设置http.Client
func (c *SDKClient) SetHttpClient(client *http.Client) {
	c.client = client
}

// SetRateLimiter 设置限流
func (c *SDKClient) SetRateLimiter(limiter RateLimiter) {
	c.limiter = limiter
}

func (c *SDKClient) WithTracer(namespace string) {
	c.tracer = NewOtel(namespace, c.appKey)
}

// Copy 复制SDKClient
func (c *SDKClient) Copy() *SDKClient {
	return &SDKClient{
		appKey: c.appKey,
		secret: c.secret,
		debug:  c.debug,
		client: c.client,
		tracer: c.tracer,
	}
}

// Get get api
func (c *SDKClient) Get(ctx context.Context, gw string, req model.GetRequest, resp model.Response) error {
	return c.get(ctx, BASE_URL, gw, req, resp)
}

func (c *SDKClient) get(ctx context.Context, base string, gw string, req model.GetRequest, resp model.Response) error {
	reqUrl := util.StringsJoin(base, gw)
	if req != nil {
		values := util.GetUrlValues()
		values.Set("appkey", c.appKey)
		req.Values(values)
		c.sign(values)
		reqUrl = util.StringsJoin(reqUrl, "?", values.Encode())
		util.PutUrlValues(values)
	}
	debug.PrintGetRequest(reqUrl, c.debug)
	httpReq, err := http.NewRequestWithContext(ctx, http.MethodGet, reqUrl, nil)
	if err != nil {
		return err
	}
	if c.limiter != nil {
		c.limiter.Take()
	}
	return c.WithSpan(ctx, httpReq, resp, nil, c.fetch)
}

// fetch execute http request
func (c *SDKClient) fetch(httpReq *http.Request, resp model.Response) (*http.Response, error) {
	httpResp, err := c.client.Do(httpReq)
	if err != nil {
		return httpResp, err
	}
	defer httpResp.Body.Close()
	if resp == nil {
		resp = &model.BaseResponse{}
	}
	body, err := debug.DecodeJSONHttpResponse(httpResp.Body, resp, c.debug)
	if httpResp.ContentLength <= 0 {
		httpResp.ContentLength = int64(len(body))
	}
	if err != nil {
		debug.PrintError(err, c.debug)
		return httpResp, errors.Join(err, model.BaseResponse{
			Status: httpResp.StatusCode,
			Des:    string(body),
		})
	}
	if resp.IsError() {
		return httpResp, resp
	}
	return httpResp, nil
}

func (c *SDKClient) WithSpan(ctx context.Context, req *http.Request, resp model.Response, payload []byte, fn func(*http.Request, model.Response) (*http.Response, error)) error {
	if c.tracer == nil {
		_, err := fn(req, resp)
		return err
	}
	return c.tracer.WithSpan(ctx, req, resp, payload, fn)
}
