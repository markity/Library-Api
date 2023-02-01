package book

import (
	"library-api/service"
	serbook "library-api/service/book"
	"library-api/util/retry"

	"github.com/gin-gonic/gin"
)

func BookSearchByName(ctx *gin.Context) {
	json_, _ := ctx.Get("json")
	json := json_.(map[string]interface{})

	authed_, _ := ctx.Get("authed")
	authed := authed_.(bool)
	var userID int64 = -1
	if authed {
		userID = ctx.GetInt64("user_id")
	}

	bookName_, ok1 := json["book_name"]
	bookName, ok2 := bookName_.(string)
	if !ok1 || !ok2 {
		service.RespInvalidParaError(ctx)
		return
	}

	var book *serbook.BookWithPariseAndFocus

	ok := retry.RetryFrame(func() error {
		b, err := serbook.TryGetBookByName(bookName, userID)
		if err != nil {
			return err
		}

		book = b
		return nil
	}, 3, "api/book/name_search.go BookSearchByName.TryGetBookByName")

	if !ok {
		service.RespServiceNotAvailabelError(ctx)
		return
	}

	if book == nil {
		serbook.RespBookSearchByNameNotFound(ctx)
		return
	}

	serbook.RespBookSearchByNameOK(ctx, book)
	// return
}
