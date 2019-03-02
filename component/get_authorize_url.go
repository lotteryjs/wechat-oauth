package component

import (
	"github.com/google/go-querystring/query"
	"github.com/lotteryjs/wechat-oauth/model"
)

// GetAuthorizeURL 获取授权页面的URL地址
func GetAuthorizeURL(authorize *model.Authorize) string {
	v, _ := query.Values(authorize)
	return v.Encode()
}
