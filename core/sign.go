package core

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"io"
	"net/url"
	"sort"
)

// 签名(sign)生成逻辑（新版）
// https://union.meituan.com/v2/apiDetail?id=27
func (c *SDKClient) sign(params url.Values) string {
	// 参数按照参数名的字典升序排列
	var keys []string
	for k := range params {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	signStr := bytes.NewBufferString(c.secret)
	for _, k := range keys {
		signStr.WriteString(k)
		signStr.WriteString(params.Get(k))
	}
	signStr.WriteString(c.secret)
	// md5加密
	has := md5.New()
	io.Copy(has, signStr)
	ret := hex.EncodeToString(has.Sum(nil))
	params.Set("sign", ret)
	return ret
}
