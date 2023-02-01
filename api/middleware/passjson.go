package middleware

import (
	"encoding/json"
	"io"
	"library-api/service"

	"github.com/gin-gonic/gin"
)

func MiddleWarePassJSON(ctx *gin.Context) {
	b, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		service.RespServiceNotAvailabelError(ctx)
		ctx.Abort()
		return
	}

	m := make(map[string]interface{})
	if json.Unmarshal(b, &m) != nil {
		service.RespInvalidParaError(ctx)
		ctx.Abort()
		return
	}

	ctx.Set("json", m)
	ctx.Next()
}
