package book

import (
	"library-api/service"
	serbook "library-api/service/book"
	"library-api/util/retry"

	"github.com/gin-gonic/gin"
)

func Praise(ctx *gin.Context, bookID int64) {
	userID := ctx.GetInt64("user_id")

	exists := false
	inserted := false

	ok := retry.RetryFrame(func() error {
		exists_, inserted_, err := serbook.TryPraise(userID, bookID)
		if err != nil {
			return err
		}

		exists = exists_
		inserted = inserted_

		return nil
	}, 3, "api/book/praise.go Praise.TryPraise")

	if !ok {
		service.RespServiceNotAvailabelError(ctx)
		return
	}

	if !exists {
		serbook.RespNoSuchBookToPraise(ctx)
		return
	}

	if !inserted {
		serbook.RespAlreadyPraised(ctx)
		return
	}

	serbook.RespPraiseOK(ctx)
	// return
}
