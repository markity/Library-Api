package comment

import (
	"library-api/service"
	sercomment "library-api/service/comment"
	fieldcheck "library-api/util/field_check"
	"library-api/util/retry"
	"strconv"

	"github.com/gin-gonic/gin"
)

func AddCommentComment(ctx *gin.Context) {
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

	commentIDStr := ctx.Param("comment_id")
	commentID, err := strconv.ParseInt(commentIDStr, 10, 64)
	if err != nil {
		service.RespInvalidParaError(ctx)
		return
	}

	lastinseted := int64(-1)
	ok := retry.RetryFrame(func() error {
		inserted, err := sercomment.TryAddCommentComment(commentID, userID, content, annoymous)
		if err != nil {
			return err
		}

		lastinseted = inserted
		return nil
	}, 3, "api/comment/add_user_comment.go AddCommentComment.TryAddCommentComment")

	if !ok {
		service.RespServiceNotAvailabelError(ctx)
	}

	if lastinseted == -1 {
		sercomment.RespCommentCommentNotFound(ctx)
		return
	}

	sercomment.RespCommentCommentOK(ctx, lastinseted)
	// return
}
