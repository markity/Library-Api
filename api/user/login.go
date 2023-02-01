package user

import (
	"library-api/api/middleware"
	"library-api/service"
	seruser "library-api/service/user"
	fieldcheck "library-api/util/field_check"
	"library-api/util/md5"
	"library-api/util/retry"
	"time"

	"github.com/gin-gonic/gin"
)

func Login(ctx *gin.Context) {
	json_, _ := ctx.Get("json")
	json := json_.(map[string]interface{})
	username_, ok1 := json["username"]
	username, ok2 := username_.(string)
	password_, ok3 := json["password"]
	password, ok4 := password_.(string)

	if !ok1 || !ok2 || !ok3 || !ok4 || !fieldcheck.CheckUsernameValid(username) ||
		!fieldcheck.CheckPasswordValid(password) {
		service.RespInvalidParaError(ctx)
		return
	}

	var loginOK bool
	var userID int64

	ok := retry.RetryFrame(func() error {
		// exist代表是否查询到 username = xxx, password = xxx的条目
		// 如果查到, 就代表密码正确, 否则就返回用户账户或密码不正确
		ok, id, err := seruser.TryCheckLoginInfo(username, md5.ToMD5(password))
		if err != nil {
			return err
		}

		loginOK = ok
		userID = id

		return nil
	}, 3, "service/user/login.go Login.TryCheckLoginInfo")

	if !ok {
		service.RespServiceNotAvailabelError(ctx)
		return
	}

	if !loginOK {
		seruser.RespUserLoginInfoWrong(ctx)
		return
	}

	// 登录成功, 签发jwt
	jwtToken := middleware.JwtSignaturer.Signature(userID, "token", time.Hour*2)
	refreshToken := middleware.JwtSignaturer.Signature(userID, "refresh", 0)

	seruser.RespLoginOK(ctx, jwtToken, refreshToken, time.Hour*2)
}
