package component

import (
	"testing"

	"github.com/lotteryjs/wechat-oauth/model"
)

func TestGetAuthorizeURL(t *testing.T) {
	type args struct {
		authorize *model.Authorize
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"生成授权Url",
			args{
				&model.Authorize{
					Appid:       "wx77317c56edf67e5a",
					RedirectURI: "https://lotteryjs.com/home/resolve",
					Scope:       "snsapi_userinfo",
					State:       "",
				},
			},
			"https://open.weixin.qq.com/connect/oauth2/authorize?appid=wx77317c56edf67e5a&redirect_uri=https%3A%2F%2Flotteryjs.com%2Fhome%2Fresolve&scope=snsapi_userinfo&state=&response_type=code#wechat_redirect",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetAuthorizeURL(tt.args.authorize); got != tt.want {
				t.Errorf("GetAuthorizeURL() = %v, want %v", got, tt.want)
			}
		})
	}
}
