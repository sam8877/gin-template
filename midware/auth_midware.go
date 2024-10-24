package midware

import (
	"fmt"
	"gin-template/conf"
	"gin-template/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"regexp"
)

type reqRexInfo struct {
	UrlRex string `json:"url_rex"`
	Method string `json:"method"`
}

func AuthMidware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		url := ctx.Request.URL.Path
		method := ctx.Request.Method
		if inWhiteList(url, method) {
			// 白名单直接放过
			ctx.Next()
		} else {
			// 认证
			tokenStr := ctx.GetHeader("token")
			if tokenStr == "" {
				panic("token is empty")
			}
			token, err := utils.ParseToken(tokenStr, conf.GetAuthConf().TokenSecretKey)
			if err != nil {
				panic("invalid token")
			}
			if token.Expired() {
				panic("token is expired")
			}

			// TODO 授权 暂不实现

			// 都成功，下一个中间件
			ctx.Next()
		}
	}
}

func inWhiteList(url string, method string) bool {
	whiteList := getWhiteList()
	for _, info := range whiteList {
		if info.Method == method && matchPath(info.UrlRex, url) {
			return true
		}
	}
	return false
}

func matchPath(pattern, path string) bool {
	r, err := regexp.Compile(pattern)
	if err != nil {
		fmt.Println("Error compiling regex:", err)
		return false
	}
	return r.MatchString(path)
}

func getWhiteList() []reqRexInfo {
	// TODO 读取白名单
	return []reqRexInfo{{UrlRex: "/create_token", Method: http.MethodPost}, {UrlRex: "/refresh_token", Method: http.MethodPost}}
}
