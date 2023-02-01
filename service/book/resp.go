package book

import (
	errorcodes "library-api/util/error_codes"
	"net/http"

	"github.com/gin-gonic/gin"
)

type bookListResp struct {
	errorcodes.BasicErrorResp
	Data []BookWithPariseAndFocus `json:"data"`
}

func RespBookListOK(ctx *gin.Context, data []BookWithPariseAndFocus) {
	bookListResp := bookListResp{}
	bookListResp.Status = errorcodes.ErrorOKCode
	bookListResp.Info = errorcodes.ErrorOKMsg
	bookListResp.Data = data

	ctx.JSON(http.StatusOK, bookListResp)
}

type bookSearchByNameResp struct {
	errorcodes.BasicErrorResp
	Data *BookWithPariseAndFocus `json:"data"`
}

func RespBookSearchByNameOK(ctx *gin.Context, data *BookWithPariseAndFocus) {
	resp := bookSearchByNameResp{}
	resp.Status = errorcodes.ErrorOKCode
	resp.Info = errorcodes.ErrorOKMsg
	resp.Data = data

	ctx.JSON(http.StatusOK, resp)
}

func RespBookSearchByNameNotFound(ctx *gin.Context) {
	resp := errorcodes.BasicErrorResp{}
	resp.Status = errorcodes.ErrorBookNameNotFoundCode
	resp.Info = errorcodes.ErrorBookNameNotFoundMsg

	ctx.JSON(http.StatusOK, resp)
}

func RespFocusOK(ctx *gin.Context) {
	resp := errorcodes.BasicErrorResp{}
	resp.Status = errorcodes.ErrorOKCode
	resp.Info = errorcodes.ErrorOKMsg

	ctx.JSON(http.StatusOK, resp)
}

func RespPraiseOK(ctx *gin.Context) {
	resp := errorcodes.BasicErrorResp{}
	resp.Status = errorcodes.ErrorOKCode
	resp.Info = errorcodes.ErrorOKMsg

	ctx.JSON(http.StatusOK, resp)
}

func RespAlreadyFoucused(ctx *gin.Context) {
	resp := errorcodes.BasicErrorResp{}
	resp.Status = errorcodes.ErrorAlreadyFocusedCode
	resp.Info = errorcodes.ErrorAlreadyFocusedMsg

	ctx.JSON(http.StatusOK, resp)
}

func RespAlreadyPraised(ctx *gin.Context) {
	resp := errorcodes.BasicErrorResp{}
	resp.Status = errorcodes.ErrorAlreadyPraisedBookCode
	resp.Info = errorcodes.ErrorAlreadyPraisedBookMsg

	ctx.JSON(http.StatusOK, resp)
}

func RespNoSuchBookToPraise(ctx *gin.Context) {
	resp := errorcodes.BasicErrorResp{}
	resp.Status = errorcodes.ErrorNoSuchBookToPraiseCode
	resp.Info = errorcodes.ErrorNoSuchBookToPraiseMsg

	ctx.JSON(http.StatusOK, resp)
}

type labelSearchOKRespData struct {
	Books []BookWithPariseAndFocus `json:"books"`
}

type labelSearchOKResp struct {
	errorcodes.BasicErrorResp
	Data labelSearchOKRespData `json:"data"`
}

func RespLableSearchOK(ctx *gin.Context, list []BookWithPariseAndFocus) {
	resp := labelSearchOKResp{}
	resp.Status = errorcodes.ErrorOKCode
	resp.Info = errorcodes.ErrorOKMsg
	resp.Data.Books = list

	ctx.JSON(http.StatusOK, resp)
}
