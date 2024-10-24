package service

import (
	"gin-template/conf"
	"gin-template/model"
	"gin-template/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func CreateToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		type UP struct {
			UserId string `json:"userId" binding:"required"`
			PassWd string `json:"passWd" binding:"required"`
		}
		var Up = UP{}
		if err := c.ShouldBind(&Up); err != nil {
			c.JSON(http.StatusBadRequest, model.BuildBadReq(err.Error()))
			return
		}
		authConf := conf.GetAuthConf()
		token, err := utils.CreateToken(Up.UserId, time.Second*time.Duration(authConf.TokenExpireSeconds), authConf.TokenSecretKey)
		if err != nil {
			c.JSON(http.StatusInternalServerError, model.BuildInternalErr(err.Error()))
			return
		}
		c.JSON(http.StatusOK, model.BuildSuccess(gin.H{"token": token}))
		return
	}
}

func RefreshToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		type Tk struct {
			Token string `json:"token" binding:"required"`
		}
		var tk = Tk{}
		if err := c.ShouldBind(&tk); err != nil {
			c.JSON(http.StatusBadRequest, model.BuildBadReq(err.Error()))
			return
		}
		tS, err := utils.ParseToken(tk.Token, conf.GetAuthConf().TokenSecretKey)
		if err != nil {
			c.JSON(http.StatusBadRequest, model.BuildBadReq(err.Error()))
			return
		}
		if !tS.Expired() {
			if tS.ExpiresAt-time.Now().Unix() < int64(conf.GetAuthConf().TokenRefreshSeconds) {
				// 生成新token
				newTk, err := utils.CreateTokenFromTokenStruct(tS, time.Second*time.Duration(conf.GetAuthConf().TokenExpireSeconds), conf.GetAuthConf().TokenSecretKey)
				if err != nil {
					c.JSON(http.StatusInternalServerError, model.BuildInternalErr(err.Error()))
					return
				}
				c.JSON(http.StatusOK, model.BuildSuccess(gin.H{"token": newTk}))
				return
			} else {
				// 返回原token
				c.JSON(http.StatusOK, model.BuildSuccess(gin.H{"token": tk.Token}))
				return
			}
		} else {
			// token 失效，不允许重新生成token
			c.JSON(http.StatusBadRequest, model.BuildBadReq("token is expired"))
			return
		}
	}
}
