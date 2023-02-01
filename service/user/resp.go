package user

import (
	errorcodes "library-api/util/error_codes"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func RespOccupiedUsername(ctx *gin.Context) {
	resp := errorcodes.BasicErrorResp{}
	resp.Status = errorcodes.ErrorUsernameOccupiedCode
	resp.Info = errorcodes.ErrorUsernameOccupiedMsg

	ctx.JSON(http.StatusOK, &resp)
}

func RespRegisterOK(ctx *gin.Context) {
	resp := errorcodes.BasicErrorResp{}
	resp.Status = errorcodes.ErrorOKCode
	resp.Info = errorcodes.ErrorOKMsg

	ctx.JSON(http.StatusOK, &resp)
}

func RespUserLoginInfoWrong(ctx *gin.Context) {
	resp := errorcodes.BasicErrorResp{}
	resp.Info = errorcodes.ErrorUserInfoWrongMsg
	resp.Status = errorcodes.ErrorUserInfoWrongCode

	ctx.JSON(http.StatusOK, &resp)
}

type loginOKRespData struct {
	Token        string        `json:"token"`
	Duration     time.Duration `json:"duration"`
	RefreshToken string        `json:"refreshToken"`
}

type loginOKResp struct {
	errorcodes.BasicErrorResp
	Data loginOKRespData `json:"data"`
}

func RespLoginOK(ctx *gin.Context, token string, refreshtoken string, t time.Duration) {
	resp := loginOKResp{}
	resp.Status = errorcodes.ErrorOKCode
	resp.Info = errorcodes.ErrorOKMsg
	resp.Data.Token = token
	resp.Data.RefreshToken = refreshtoken
	resp.Data.Duration = t

	ctx.JSON(http.StatusOK, &resp)
}

type refreshOKRespData struct {
	Token        string        `json:"token"`
	Duration     time.Duration `json:"duration"`
	RefreshToken string        `json:"refresh_token"`
}

type refreshOKResp struct {
	errorcodes.BasicErrorResp
	Data refreshOKRespData `json:"data"`
}

func RespRefreshOK(ctx *gin.Context, token string, refreshtoken string, t time.Duration) {
	resp := refreshOKResp{}
	resp.Status = errorcodes.ErrorOKCode
	resp.Info = errorcodes.ErrorOKMsg
	resp.Data.Token = token
	resp.Data.RefreshToken = refreshtoken
	resp.Data.Duration = t

	ctx.JSON(http.StatusOK, &resp)
}

func RespChangePasswordOK(ctx *gin.Context) {
	resp := errorcodes.BasicErrorResp{}
	resp.Info = errorcodes.ErrorOKMsg
	resp.Status = errorcodes.ErrorOKCode

	ctx.JSON(http.StatusOK, &resp)
}

func RespWrongOldPasswordOK(ctx *gin.Context) {
	resp := errorcodes.BasicErrorResp{}
	resp.Info = errorcodes.ErrorWrongOldPasswordMsg
	resp.Status = errorcodes.ErrorWrongOldPasswordCode

	ctx.JSON(http.StatusOK, &resp)
}

func RespEditUserInfoOK(ctx *gin.Context) {
	resp := errorcodes.BasicErrorResp{}
	resp.Info = errorcodes.ErrorOKMsg
	resp.Status = errorcodes.ErrorOKCode

	ctx.JSON(http.StatusOK, &resp)
}
