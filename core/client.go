package core

import (
	"context"
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/bububa/meituan/v2/core/internal/debug"
	"github.com/bububa/meituan/v2/model"
	"github.com/bububa/meituan/v2/util"
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
	client *http.Client
	tracer *Otel
	appKey string
	secret string
	debug  bool
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
func (c *SDKClient) POST(ctx context.Context, gw string, req model.PostRequest, resp model.Response) error {
	return c.post(ctx, BASE_URL, gw, req, resp)
}

func (c *SDKClient) post(ctx context.Context, base string, gw string, req model.PostRequest, resp model.Response) error {
	reqPath := util.StringsJoin(BASE_PATH, gw)
	reqUrl := util.StringsJoin(base, reqPath)
	buffer := util.GetBufferPool()
	defer util.PutBufferPool(buffer)
	req.Encode(buffer)
	ts := time.Now().UnixMilli()
	tsStr := strconv.FormatInt(ts, 10)
	md5Hash := md5.Sum(buffer.Bytes())
	contentMD5 := base64.StdEncoding.EncodeToString(md5Hash[:])
	stringToSign := util.StringsJoin("POST\n", contentMD5, "\nS-Ca-App:", c.appKey, "\nS-Ca-Timestamp:", tsStr, "\n", reqPath)

	signH := hmac.New(sha256.New, []byte(c.secret))
	signH.Write([]byte(stringToSign))
	signature := base64.StdEncoding.EncodeToString(signH.Sum(nil))

	debug.PrintPostJSONRequest(reqUrl, buffer.Bytes(), c.debug)
	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, reqUrl, buffer)
	if err != nil {
		return err
	}
	httpReq.Header.Set("Content-Type", "application/json")
	httpReq.Header.Set("S-Ca-App", c.appKey)
	httpReq.Header.Set("S-Ca-Signature", c.appKey)
	httpReq.Header.Set("S-Ca-Timestamp", tsStr)
	httpReq.Header.Set("Content-MD5", contentMD5)
	httpReq.Header.Set("S-Ca-Signature", signature)
	httpReq.Header.Set("S-Ca-Signature-Headers", "S-Ca-Timestamp,S-Ca-App")
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
			Code:    httpResp.StatusCode,
			Message: string(body),
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
