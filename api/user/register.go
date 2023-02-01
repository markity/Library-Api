package user

import (
	"library-api/service"
	seruser "library-api/service/user"
	fieldcheck "library-api/util/field_check"
	"library-api/util/md5"
	"library-api/util/retry"
	"time"

	"github.com/gin-gonic/gin"
)

func Register(ctx *gin.Context) {
	json_, _ := ctx.Get("json")
	json := json_.(map[string]interface{})
	username_, ok1 := json["username"]
	username, ok2 := username_.(string)
	password_, ok3 := json["password"]
	password, ok4 := password_.(string)

	// 检验参数合法性
	if !ok1 || !ok2 || !ok3 || !ok4 || !fieldcheck.CheckUsernameValid(username) ||
		!fieldcheck.CheckPasswordValid(password) {
		service.RespInvalidParaError(ctx)
		return
	}

	isAlreadyRegistered := false
	ok := retry.RetryFrame(func() error {
		err, insertOK := seruser.TryCreateUser(username, md5.ToMD5(password), time.Now())
		if err != nil {
			return err
		}
		isAlreadyRegistered = !insertOK
		return nil
	}, 3, "service/user/register.go Register.TryCreateUser")

	if !ok {
		service.RespServiceNotAvailabelError(ctx)
		return
	}

	if isAlreadyRegistered {
		seruser.RespOccupiedUsername(ctx)
		return
	}

	seruser.RespRegisterOK(ctx)
	// return
}
