package middleware

// 一个中间件, 对于需要jwt鉴权的用户

import (
	"library-api/service"
	"library-api/util/jwt"
	"strings"

	"github.com/gin-gonic/gin"
)

// 用于JWT签名以及鉴权
var JwtSignaturer jwt.JWTSignaturer

// 用于鉴权后传递给下一级中间件的用户信息
type UserAuthInfo struct {
	UserID int64
}

// 加载该包的时候生成一个jwt签名器
func init() {
	JwtSignaturer = jwt.NewUserJWTSignaturer(jwt.NewRsaSHA256Cryptor())
}

func MiddleWareJWTVerify(ctx *gin.Context) {
	jwtStr := ctx.GetHeader("Authorization")
	// Bearer[空格]token
	if !strings.HasPrefix(jwtStr, "Bearer ") {
		service.RespJWTError(ctx)
		ctx.Abort()
		return
	}
	jwtStr = jwtStr[7:]

	valid, payload := JwtSignaturer.CheckAndUnpackPayload(jwtStr)
	if !valid || payload.TokenType != "token" {
		// 未鉴权的错误
		service.RespJWTError(ctx)
		ctx.Abort()
		return
	}

	// 鉴权成功, set用户数据, 并移交给下一个中间件
	ctx.Set("user_id", payload.UserID)
	ctx.Next()
}

// 如果用户登录, 获得用户的user_id, 但用户不登录, 也可以访问
func MiddleWareJWTMention(ctx *gin.Context) {
	jwtStr := ctx.GetHeader("Authorization")
	// Bearer[空格]token
	if !strings.HasPrefix(jwtStr, "Bearer ") {
		ctx.Set("authed", false)
		ctx.Next()
		return
	}
	jwtStr = jwtStr[7:]

	valid, payload := JwtSignaturer.CheckAndUnpackPayload(jwtStr)
	if !valid || payload.TokenType != "token" {
		// 未鉴权的错误
		ctx.Set("authed", false)
		ctx.Next()
		return
	}

	// 鉴权成功, set用户数据, 并移交给下一个中间件
	ctx.Set("authed", true)
	ctx.Set("user_id", payload.UserID)
	ctx.Next()
}
