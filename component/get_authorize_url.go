package component

import (
	"github.com/google/go-querystring/query"
	"github.com/lotteryjs/wechat-oauth/model"
)

// GetAuthorizeURL 获取授权页面的URL地址
func GetAuthorizeURL(authorize *model.Authorize) string {
	v, _ := query.Values(authorize)
	return "https://open.weixin.qq.com/connect/oauth2/authorize?" + v.Encode() + "&response_type=code#wechat_redirect"
}
