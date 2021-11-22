package req

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func RequestParams() gin.HandlerFunc {
	return func(context *gin.Context) {
		if context.Request.Method == "POST" {
			params := make(map[string]interface{})
			err := context.ShouldBindJSON(&params)
			if err != nil {
				return
			}
			zap.S().Info(params)
		}
		context.Next()
	}
}
