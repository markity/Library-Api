package user

import (
	"library-api/api/middleware"
	"library-api/service"
	seruser "library-api/service/user"
	"time"

	"github.com/gin-gonic/gin"
)

func RefreshToken(ctx *gin.Context) {
	json_, _ := ctx.Get("json")
	json := json_.(map[string]interface{})
	refresh_, ok1 := json["refresh_token"]
	refresh, ok2 := refresh_.(string)

	if !ok1 || !ok2 {
		service.RespInvalidParaError(ctx)
		return
	}
	ok, p := middleware.JwtSignaturer.CheckAndUnpackPayload(refresh)
	if !ok || p.TokenType != "refresh" {
		service.RespInvalidParaError(ctx)
		return
	}

	jwtToken := middleware.JwtSignaturer.Signature(p.UserID, "token", time.Hour*2)

	seruser.RespRefreshOK(ctx, jwtToken, refresh, time.Hour*2)
}
