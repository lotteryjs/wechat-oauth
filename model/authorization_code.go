package model

// AuthorizeCode Model
//
// 生成网页授权地址的参数模型
//
// swagger:model AuthorizeCode
type AuthorizeCode struct {
	// 公众号的唯一标识
	//
	// required: true
	// example: wx23dfasdfw2
	Appid string `url:"appid"`
	// 公众号的appsecret
	//
	// required: true
	// example: afsasdfsacsdfas2sasdd222
	Secret string `url:"secret"`
	// 填写第一步获取的code参数
	//
	// required: true
	// example: code
	Code string `url:"code"`
	// 填写为authorization_code
	//
	// required: true
	// example: authorization_code
	GrantType string `url:"grant_type"`
}
