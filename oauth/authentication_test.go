package oauth

import "testing"

func TestAuth_GetAuthorizeURL(t *testing.T) {
	type fields struct {
		appid     string
		appsecret string
	}
	type args struct {
		redirect string
		state    string
		scope    string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			`GetAuthorizeURL("http://lotteryjs.com/","","")`,
			fields{
				"appid",
				"appsecret",
			},
			args{
				"http://lotteryjs.com/",
				"",
				"",
			},
			"https://open.weixin.qq.com/connect/oauth2/authorize?appid=appid&redirect_uri=http%3A%2F%2Flotteryjs.com%2F&scope=snsapi_base&state=&response_type=code#wechat_redirect",
		},
		{
			`GetAuthorizeURL("http://lotteryjs.com/","hehe","")`,
			fields{
				"appid",
				"appsecret",
			},
			args{
				"http://lotteryjs.com/",
				"hehe",
				"",
			},
			"https://open.weixin.qq.com/connect/oauth2/authorize?appid=appid&redirect_uri=http%3A%2F%2Flotteryjs.com%2F&scope=snsapi_base&state=hehe&response_type=code#wechat_redirect",
		},
		{
			`GetAuthorizeURL("http://lotteryjs.com/","hehe","snsapi_userinfo")`,
			fields{
				"appid",
				"appsecret",
			},
			args{
				"http://lotteryjs.com/",
				"hehe",
				"snsapi_userinfo",
			},
			"https://open.weixin.qq.com/connect/oauth2/authorize?appid=appid&redirect_uri=http%3A%2F%2Flotteryjs.com%2F&scope=snsapi_userinfo&state=hehe&response_type=code#wechat_redirect",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Auth{
				appid:     tt.fields.appid,
				appsecret: tt.fields.appsecret,
			}
			if got := a.GetAuthorizeURL(tt.args.redirect, tt.args.state, tt.args.scope); got != tt.want {
				t.Errorf("Auth.GetAuthorizeURL() = %v, want %v", got, tt.want)
			}
		})
	}
}
