package operate

import (
	"library-api/model"
	errorcodes "library-api/util/error_codes"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ListFocusOKResp struct {
	errorcodes.BasicErrorResp
	Data []*model.FocusListItem `json:"data"`
}

func RespListFocusOK(ctx *gin.Context, data []*model.FocusListItem) {
	resp := ListFocusOKResp{}
	resp.Status = errorcodes.ErrorOKCode
	resp.Info = errorcodes.ErrorOKMsg
	resp.Data = data

	ctx.JSON(http.StatusOK, resp)
}
