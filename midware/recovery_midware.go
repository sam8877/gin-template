package midware

import (
	"gin-template/model"
	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	"net/http"
)

func Recovery() gin.HandlerFunc {
	return gin.CustomRecovery(func(c *gin.Context, errIn any) {
		b, err := json.Marshal(errIn)
		if err != nil {
			c.JSON(http.StatusInternalServerError, model.BuildInternalErr(err.Error()))
			return
		}
		c.JSON(http.StatusInternalServerError, model.BuildInternalErr(string(b)))
	})
}
