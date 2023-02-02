package book

import (
	"library-api/service"
	serbook "library-api/service/book"
	"library-api/util/retry"

	"github.com/gin-gonic/gin"
)

func Focus(ctx *gin.Context) {
	json_, _ := ctx.Get("json")
	json := json_.(map[string]interface{})
	bookID_, ok1 := json["book_id"]
	bookID__, ok2 := bookID_.(float64)
	bookID := int64(bookID__)
	userID := ctx.GetInt64("user_id")

	if !ok1 || !ok2 {
		println(ok1, ok2)
		service.RespInvalidParaError(ctx)
		return
	}

	var exists bool
	var inserted bool

	ok := retry.RetryFrame(func() error {
		exists_, inserted_, err := serbook.TryFocus(userID, bookID)
		if err != nil {
			return err
		}

		exists = exists_
		inserted = inserted_

		return nil
	}, 3, "api/book/focus.go Focus.TryFocus")

	if !ok {
		service.RespServiceNotAvailabelError(ctx)
		return
	}

	if !exists {
		serbook.RespNoSuchBookToFocus(ctx)
		return
	}

	if !inserted {
		serbook.RespAlreadyFoucused(ctx)
		return
	}

	serbook.RespFocusOK(ctx)
	// return
}
