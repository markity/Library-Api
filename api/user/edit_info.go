package user

import (
	"library-api/service"
	seruser "library-api/service/user"
	fieldcheck "library-api/util/field_check"
	"library-api/util/retry"

	"github.com/gin-gonic/gin"
)

func ChangeUserInfo(ctx *gin.Context) {
	json_, _ := ctx.Get("json")
	json := json_.(map[string]interface{})
	userID_, _ := ctx.Get("user_id")
	userID := userID_.(int64)

	avartar_, doAvatar := json["avatar_link"]
	var avartar *string
	if doAvatar {
		if avartar_ == nil {
			avartar = nil
		} else {
			a, ok := avartar_.(string)
			if !ok || !fieldcheck.CheckLinkVaild(a) {
				service.RespInvalidParaError(ctx)
				return
			}
			avartar = &a
		}
	}

	intro_, doIntro := json["introduction"]
	var intro *string
	if doIntro {
		if intro_ == nil {
			intro = nil
		} else {
			a, ok := intro_.(string)
			if !ok || !fieldcheck.CheckIntroductionVaild(a) {
				service.RespInvalidParaError(ctx)
				return
			}
			intro = &a
		}
	}

	phone_, doPhone := json["phone"]
	var phone *string
	if doPhone {
		if phone_ == nil {
			phone = nil
		} else {
			a, ok := phone_.(string)
			if !ok || !fieldcheck.CheckPhoneVaild(a) {
				service.RespInvalidParaError(ctx)
				return
			}
			phone = &a
		}
	}

	gender_, doGender := json["gender"]
	var gender *int
	if doGender {
		if gender_ == nil {
			gender = nil
		} else {
			a, ok := gender_.(int)
			if !ok || !fieldcheck.CheckGenderVaild(a) {
				service.RespInvalidParaError(ctx)
				return
			}

			gender = &a
		}
	}

	email_, doEmail := json["email"]
	var email *string
	if doEmail {
		if email_ == nil {
			email = nil
		} else {
			a, ok := email_.(string)
			if !ok || !fieldcheck.CheckEmailVaild(a) {
				service.RespInvalidParaError(ctx)
				return
			}
			email = &a
		}
	}

	birthday_, doBirthday := json["birthday"]
	var birthday *string
	if doBirthday {
		if birthday_ == nil {
			birthday = nil
		} else {
			a, ok := birthday_.(string)
			if !ok || !fieldcheck.CheckBirthdayVaild(a) {
				service.RespInvalidParaError(ctx)
				return
			}
			birthday = &a
		}
	}

	ok := retry.RetryFrame(func() error {
		err := seruser.TryEditUserInfo(userID, doAvatar, avartar, doBirthday, birthday, doEmail, email,
			doGender, gender, doIntro, intro, doPhone, phone)
		if err != nil {
			return err
		}

		return nil
	}, 3, "api/user/edit_info.go ChangeUserInfo. TryEditUserInfo")

	if !ok {
		service.RespServiceNotAvailabelError(ctx)
		return
	}

	seruser.RespEditUserInfoOK(ctx)
	// return
}
