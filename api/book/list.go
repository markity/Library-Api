package book

import (
	"fmt"
	"library-api/service"
	serbook "library-api/service/book"
	"library-api/util/retry"

	"github.com/gin-gonic/gin"
)

func BookList(ctx *gin.Context) {
	authed_, _ := ctx.Get("authed")
	authed := authed_.(bool)
	var userID int64 = -1
	if authed {
		userID = ctx.GetInt64("user_id")
	}

	var bookdata []serbook.BookWithPariseAndFocus
	ok := retry.RetryFrame(func() error {
		books, err := serbook.TryGetBooksWithPraiseAndFocus(userID)
		if err != nil {
			println("fuck")
			return err
		}

		fmt.Println(books)

		bookdata = books
		return nil
	}, 3, "api/book/list.go BookList.TryGetBooksWithPraiseAndFocus")

	if !ok {
		service.RespServiceNotAvailabelError(ctx)
		return
	}

	serbook.RespBookListOK(ctx, bookdata)
	// return
}
