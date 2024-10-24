package service

import (
	"gin-template/conf"
	"gin-template/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func CreateToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		type UP struct {
			UserId string `json:"userId" requird:"true"`
			PassWd string `json:"passWd" requird:"true"`
		}
		var Up = UP{}
		c.ShouldBindJSON(&Up)
		token, err := utils.CreateToken(Up.UserId, time.Second*time.Duration(conf.GetAuthConf().TokenExpireSeconds), conf.GetAuthConf().TokenSecretKey)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"code": http.StatusInternalServerError, "msg": err.Error()})
		}
		c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "msg": "success", "data": gin.H{"token": token}})
	}
}

func RefreshToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		type Tk struct {
			Token string `json:"token" requird:"true"`
		}
		var tk = Tk{}
		c.ShouldBindJSON(&tk)

		tS, err := utils.ParseToken(tk.Token, conf.GetAuthConf().TokenSecretKey)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"code": http.StatusInternalServerError, "msg": err.Error()})
		}
		if !tS.Expired() {
			if tS.ExpiresAt-time.Now().Unix() < int64(conf.GetAuthConf().TokenRefreshSeconds) {
				// 生成新token
				newTk, err := utils.CreateTokenFromTokenStruct(tS, time.Second*time.Duration(conf.GetAuthConf().TokenExpireSeconds), conf.GetAuthConf().TokenSecretKey)
				if err != nil {
					c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"code": http.StatusInternalServerError, "msg": err.Error()})
				}
				c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "msg": "success", "data": gin.H{"token": newTk}})
			} else {
				// 返回原token
				c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "msg": "success", "data": gin.H{"token": tk.Token}})
			}
		} else {
			// token 失效，不允许重新生成token
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "msg": "token is expired"})
		}
	}
}
