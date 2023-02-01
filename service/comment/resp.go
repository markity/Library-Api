package comment

import (
	"library-api/model"
	errorcodes "library-api/util/error_codes"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RespBookNotFound(ctx *gin.Context) {
	resp := errorcodes.BasicErrorResp{}
	resp.Info = errorcodes.ErrorCommentBookNotFoundMsg
	resp.Status = errorcodes.ErrorCommentBookNotFoundCode

	ctx.JSON(http.StatusOK, resp)
}

type bookCommentOKResp struct {
	errorcodes.BasicErrorResp
	Data int64 `json:"data"`
}

func RespBookCommentOK(ctx *gin.Context, lastinserted int64) {
	resp := bookCommentOKResp{}
	resp.Info = errorcodes.ErrorOKMsg
	resp.Status = errorcodes.ErrorOKCode
	resp.Data = lastinserted

	ctx.JSON(http.StatusOK, resp)
}

type commentCommentOKResp struct {
	errorcodes.BasicErrorResp
	Data int64 `json:"data"`
}

func RespCommentCommentOK(ctx *gin.Context, lastinserted int64) {
	resp := commentCommentOKResp{}
	resp.Info = errorcodes.ErrorOKMsg
	resp.Status = errorcodes.ErrorOKCode
	resp.Data = lastinserted

	ctx.JSON(http.StatusOK, resp)
}

func RespCommentCommentNotFound(ctx *gin.Context) {
	resp := errorcodes.BasicErrorResp{}
	resp.Info = errorcodes.ErrorCommentBookNotFoundMsg
	resp.Status = errorcodes.ErrorCommentCommentNotFoundCode

	ctx.JSON(http.StatusOK, resp)
}

func RespNoCommentToUpdate(ctx *gin.Context) {
	resp := errorcodes.BasicErrorResp{}
	resp.Info = errorcodes.ErrorNoSuchCommentToUpdateMsg
	resp.Status = errorcodes.ErrorNoSuchCommentToUpdateCode

	ctx.JSON(http.StatusOK, resp)
}

func RespNoPermissionToUpdateComment(ctx *gin.Context) {
	resp := errorcodes.BasicErrorResp{}
	resp.Info = errorcodes.ErrorNoPermissionToUpdateCommentMsg
	resp.Status = errorcodes.ErrorNoPermissionToUpdateCommentCode

	ctx.JSON(http.StatusOK, resp)
}

func RespUpdateCommentOK(ctx *gin.Context) {
	resp := errorcodes.BasicErrorResp{}
	resp.Info = errorcodes.ErrorOKMsg
	resp.Status = errorcodes.ErrorOKCode

	ctx.JSON(http.StatusOK, resp)
}

func RespNoSuchCommentToDelete(ctx *gin.Context) {
	resp := errorcodes.BasicErrorResp{}
	resp.Info = errorcodes.ErrorNoSuchCommentToDeleteMsg
	resp.Status = errorcodes.ErrorNoSuchCommentToDeleteCode

	ctx.JSON(http.StatusOK, resp)
}

func RespNoPermissionToDeleteComment(ctx *gin.Context) {
	resp := errorcodes.BasicErrorResp{}
	resp.Info = errorcodes.ErrorNoPermissionToDeleteCommentMsg
	resp.Status = errorcodes.ErrorNoPermissionToDeleteCommentCode

	ctx.JSON(http.StatusOK, resp)
}

func RespDeleteCommentOK(ctx *gin.Context) {
	resp := errorcodes.BasicErrorResp{}
	resp.Info = errorcodes.ErrorOKMsg
	resp.Status = errorcodes.ErrorOKCode

	ctx.JSON(http.StatusOK, resp)
}

func RespNoSuchBookToListComments(ctx *gin.Context) {
	resp := errorcodes.BasicErrorResp{}
	resp.Info = errorcodes.ErrorNoSuchBookToListCommentsMsg
	resp.Status = errorcodes.ErrorNoSuchBookToListCommentsCode

	ctx.JSON(http.StatusOK, resp)
}

type listBookCommetsOKResp struct {
	errorcodes.BasicErrorResp
	Data []*model.CommentWithParise `json:"data"`
}

func RespListBookCommentsOK(ctx *gin.Context, data []*model.CommentWithParise) {
	resp := listBookCommetsOKResp{}
	resp.Status = errorcodes.ErrorOKCode
	resp.Info = errorcodes.ErrorOKMsg
	resp.Data = data
	ctx.JSON(http.StatusOK, resp)
}

func RespAlreadyPraised(ctx *gin.Context) {
	resp := errorcodes.BasicErrorResp{}
	resp.Status = errorcodes.ErrorAlreadyPraisedCommentCode
	resp.Info = errorcodes.ErrorAlreadyPraisedCommentMsg

	ctx.JSON(http.StatusOK, resp)
}

func RespNoSuchCommentToPraise(ctx *gin.Context) {
	resp := errorcodes.BasicErrorResp{}
	resp.Status = errorcodes.ErrorNoSuchCommentToPraiseCode
	resp.Info = errorcodes.ErrorNoSuchCommentToPraiseMsg

	ctx.JSON(http.StatusOK, resp)
}

func RespPraiseOK(ctx *gin.Context) {
	resp := errorcodes.BasicErrorResp{}
	resp.Info = errorcodes.ErrorOKMsg
	resp.Status = errorcodes.ErrorOKCode

	ctx.JSON(http.StatusOK, resp)
}
