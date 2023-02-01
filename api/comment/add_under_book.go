package comment

import (
	"library-api/service"
	sercomment "library-api/service/comment"
	fieldcheck "library-api/util/field_check"
	"library-api/util/retry"
	"strconv"

	"github.com/gin-gonic/gin"
)

func AddBookComment(ctx *gin.Context) {
	json_, _ := ctx.Get("json")
	json := json_.(map[string]interface{})

	content_, ok1 := json["content"]
	content, ok2 := content_.(string)
	if !ok1 || !ok2 || !fieldcheck.CheckCommentVaile(content) {
		service.RespInvalidParaError(ctx)
		return
	}

	annoymous_, ok3 := json["annoymous"]
	annoymous := false
	var ok4 bool
	if ok3 {
		annoymous, ok4 = annoymous_.(bool)
		if !ok4 {
			service.RespInvalidParaError(ctx)
			return
		}
	}

	userID := ctx.GetInt64("user_id")

	bookIDStr := ctx.Param("book_id")
	bookID, err := strconv.ParseInt(bookIDStr, 10, 64)
	if err != nil {
		service.RespInvalidParaError(ctx)
		return
	}

	var lastinserted int64

	ok := retry.RetryFrame(func() error {
		inserted, err := sercomment.TryAddBookComment(bookID, userID, content, annoymous)
		if err != nil {
			return err
		}

		lastinserted = inserted

		return nil
	}, 3, "api/comment/add_under_book.go AddBookComment.TryAddBookComment")

	if !ok {
		service.RespServiceNotAvailabelError(ctx)
		return
	}

	if lastinserted == -1 {
		sercomment.RespBookNotFound(ctx)
		return
	}

	sercomment.RespBookCommentOK(ctx, lastinserted)
	// return
}
