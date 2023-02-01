package comment

import (
	"fmt"
	"library-api/service"
	sercomment "library-api/service/comment"
	"library-api/util/retry"
	"strconv"

	"github.com/gin-gonic/gin"
)

// /comment/{comment_id}
func DeleteComment(ctx *gin.Context) {
	userID := ctx.GetInt64("user_id")

	commentIDStr := ctx.Param("comment_id")
	commentID, err := strconv.ParseInt(commentIDStr, 10, 64)
	if err != nil {
		fmt.Println(err)
		service.RespInvalidParaError(ctx)
		return
	}
	println("fuck")

	var exists, permissionDeny bool

	ok := retry.RetryFrame(func() error {
		exists_, permissionDeny_, err := sercomment.TryDeleteComment(commentID, userID)
		if err != nil {
			return err
		}

		exists = exists_
		permissionDeny = permissionDeny_
		return nil
	}, 3, "api/comment/delete_comment.go")

	if !ok {
		service.RespServiceNotAvailabelError(ctx)
		return
	}

	if !exists {
		sercomment.RespNoSuchCommentToDelete(ctx)
		return
	}

	if permissionDeny {
		sercomment.RespNoPermissionToDeleteComment(ctx)
		return
	}

	sercomment.RespDeleteCommentOK(ctx)
	// return
}
