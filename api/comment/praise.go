package comment

import (
	"library-api/service"
	sercomment "library-api/service/comment"
	"library-api/util/retry"

	"github.com/gin-gonic/gin"
)

func Praise(ctx *gin.Context, bookID int64) {
	userID := ctx.GetInt64("user_id")

	exists := false
	inserted := false

	ok := retry.RetryFrame(func() error {
		exists_, inserted_, err := sercomment.TryPraise(userID, bookID)
		if err != nil {
			return err
		}

		exists = exists_
		inserted = inserted_

		return nil
	}, 3, "api/comment/praise.go Praise.TryPraise")

	if !ok {
		service.RespServiceNotAvailabelError(ctx)
		return
	}

	if !exists {
		sercomment.RespNoSuchCommentToPraise(ctx)
		return
	}

	if !inserted {
		sercomment.RespAlreadyPraised(ctx)
		return
	}

	sercomment.RespPraiseOK(ctx)
	// return
}
