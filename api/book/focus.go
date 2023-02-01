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

	isAlreadyFocused := false

	ok := retry.RetryFrame(func() error {
		inserted, err := serbook.TryFocus(userID, bookID)
		if err != nil {
			return err
		}

		isAlreadyFocused = !inserted

		return nil
	}, 3, "api/book/focus.go Focus.TryFocus")

	println(isAlreadyFocused)

	if !ok {
		service.RespServiceNotAvailabelError(ctx)
		return
	}

	if isAlreadyFocused {
		serbook.RespAlreadyFoucused(ctx)
		return
	}

	serbook.RespFocusOK(ctx)
	// return
}
