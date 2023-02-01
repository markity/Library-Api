package comment

import (
	"library-api/model"
	"library-api/service"
	"library-api/util/retry"
	"strconv"

	sercomment "library-api/service/comment"

	"github.com/gin-gonic/gin"
)

func ListCommentsUnderBook(ctx *gin.Context) {
	authed_, _ := ctx.Get("authed")
	authed := authed_.(bool)
	var userID int64 = -1
	if authed {
		userID = ctx.GetInt64("user_id")
	}

	bookIDStr := ctx.Param("book_id")
	bookID, err := strconv.ParseInt(bookIDStr, 10, 64)
	if err != nil {
		service.RespInvalidParaError(ctx)
		return
	}

	var comments []*model.CommentWithParise
	var exists bool

	ok := retry.RetryFrame(func() error {
		comments_, exists_, err := sercomment.TryListCommentsUnderBook(userID, bookID)
		if err != nil {
			return err
		}

		comments = comments_
		exists = exists_

		return nil
	}, 3, "api/comment/list_under_book.go ListCommentsUnderBook.TryListCommentsUnderBook")

	if !ok {
		service.RespServiceNotAvailabelError(ctx)
		return
	}

	if !exists {
		sercomment.RespNoSuchBookToListComments(ctx)
		return
	}

	sercomment.RespListBookCommentsOK(ctx, comments)
	//return
}
