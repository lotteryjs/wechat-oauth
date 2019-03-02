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
					Appid:       "wx23dfasdfw2",
					RedirectURI: "https://lotteryjs.com/resolve",
					Scope:       "snsapi_userinfo",
					State:       "",
				},
			},
			"appid=wx23dfasdfw2&redirect_uri=https%3A%2F%2Flotteryjs.com%2Fresolve&scope=snsapi_userinfo&state=",
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
