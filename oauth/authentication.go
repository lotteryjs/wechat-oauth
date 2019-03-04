package oauth

import (
	"github.com/google/go-querystring/query"
	"github.com/lotteryjs/wechat-oauth/model"
	"net/http"
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

// GetAccessToken 根据授权获取到的code，换取access token和openid
func (a *Auth) GetAccessToken(code string) {
	model := &model.AuthorizeCode{
		Appid:     a.appid,
		Secret:    a.appsecret,
		Code:      code,
		GrantType: "authorization_code",
	}

	v, _ := query.Values(model)
	url := "https://api.weixin.qq.com/sns/oauth2/access_token" + v.Encode()

	resp, err := http.DefaultClient.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()
}
