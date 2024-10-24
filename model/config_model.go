package model

type ConfStruct struct {
	Auth *AuthStruct `mapstructure:"auth"`
}

type AuthStruct struct {
	TokenSecretKey      string       `mapstructure:"secret_key"`      // token加密密钥
	TokenExpireSeconds  uint         `mapstructure:"expire_seconds"`  // token失效时间
	TokenRefreshSeconds uint         `mapstructure:"refresh_seconds"` // token二次刷新时间
	WhiteList           []ReqRexInfo `mapstructure:"white_list"`      // 白名单
}

type ReqRexInfo struct {
	UrlRex string `json:"url_rex" mapstructure:"url_rex"` // URL 正则
	Method string `json:"method" mapstructure:"method"`   // http请求method
}
