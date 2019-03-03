package oauth

import (
	"github.com/google/go-querystring/query"
	"github.com/lotteryjs/wechat-oauth/model"
	"strings"
)

// Auth is the provider
type Auth struct {
	// 公众号的唯一标识
	appid string
	// 公众号的appsecret
	appsecret string
}

// New 构造 OAuth 方法
func New(appid string, appsecret string) *Auth {
	return &Auth{appid, appsecret}
}

// GetAuthorizeURL 获取授权页面的URL地址
func (a *Auth) GetAuthorizeURL(redirect string, state string, scope string) string {
	if strings.TrimSpace(scope) == "" {
		scope = "snsapi_base"
	}
	model := &model.Authorize{
		Appid:       a.appid,
		RedirectURI: redirect,
		Scope:       scope,
		State:       state}

	v, _ := query.Values(model)
	return "https://open.weixin.qq.com/connect/oauth2/authorize?" + v.Encode() + "&response_type=code#wechat_redirect"
}
