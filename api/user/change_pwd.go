package user

import (
	"library-api/service"
	seruser "library-api/service/user"
	fieldcheck "library-api/util/field_check"
	"library-api/util/md5"
	"library-api/util/retry"

	"github.com/gin-gonic/gin"
)

func ChangePassword(ctx *gin.Context) {
	json_, _ := ctx.Get("json")
	json := json_.(map[string]interface{})
	userID_, _ := ctx.Get("user_id")
	userID := userID_.(int64)
	oldPassword_, ok1 := json["old_password"]
	oldPassword, ok2 := oldPassword_.(string)
	newPassword_, ok3 := json["new_password"]
	newPassword, ok4 := newPassword_.(string)
	if !ok1 || !ok2 || !ok3 || !ok4 || !fieldcheck.CheckPasswordValid(oldPassword) ||
		!fieldcheck.CheckPasswordValid(newPassword) || oldPassword == newPassword {
		service.RespInvalidParaError(ctx)
		return
	}

	changeOK := false

	ok := retry.RetryFrame(func() error {
		ok, err := seruser.TryChangePassword(userID, md5.ToMD5(oldPassword), md5.ToMD5(newPassword))
		if err != nil {
			return err
		}

		changeOK = ok

		return nil
	}, 3, "api/user/change_pad.go ChangePassword.TryChangePassword")

	if !ok {
		service.RespServiceNotAvailabelError(ctx)
		return
	}

	if !changeOK {
		seruser.RespWrongOldPasswordOK(ctx)
		return
	}

	seruser.RespChangePasswordOK(ctx)
	// return
}
