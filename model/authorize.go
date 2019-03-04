package model

// Authorize Model
//
// 生成网页授权地址的参数模型
//
// swagger:model Authorize
type Authorize struct {
	// 公众号的唯一标识
	//
	// required: true
	// example: wx23dfasdfw2
	Appid string `url:"appid"`
	// 授权后重定向的回调链接地址， 请使用 urlEncode 对链接进行处理
	//
	// required: true
	// example: https://lotteryjs.com/resolve
	RedirectURI string `url:"redirect_uri"`
	// 应用授权作用域
	// snsapi_base （不弹出授权页面，直接跳转，只能获取用户openid）
	// snsapi_userinfo （弹出授权页面，可通过openid拿到昵称、性别、所在地。并且， 即使在未关注的情况下，只要用户授权，也能获取其信息 ）
	//
	// required: true
	// example: snsapi_userinfo
	Scope string `url:"scope"`
	// 公众号的唯一标识
	//
	// required: false
	// example: 1
	State string `url:"state"`
}
