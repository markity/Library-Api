package comment

import (
	"library-api/service"
	sercomment "library-api/service/comment"
	fieldcheck "library-api/util/field_check"
	"library-api/util/retry"
	"strconv"

	"github.com/gin-gonic/gin"
)

func UpdateComment(ctx *gin.Context) {
	json_, _ := ctx.Get("json")
	json := json_.(map[string]interface{})

	commentIDStr := ctx.Param("comment_id")
	commentID, err := strconv.ParseInt(commentIDStr, 10, 64)
	if err != nil {
		service.RespInvalidParaError(ctx)
		return
	}

	userID := ctx.GetInt64("user_id")

	content_, ok1 := json["content"]
	var content string
	var updateContent bool
	if !ok1 {
		updateContent = false
	} else {
		updateContent = true
		var ok bool
		content, ok = content_.(string)
		if !ok {
			service.RespInvalidParaError(ctx)
			return
		}
		if !fieldcheck.CheckCommentVaile(content) {
			service.RespInvalidParaError(ctx)
			return
		}
	}

	annoymous_, ok2 := json["annoymous"]
	var annoymous bool
	var updateAnnoymous bool
	if !ok2 {
		updateAnnoymous = false
	} else {
		updateAnnoymous = true
		var ok bool
		annoymous, ok = annoymous_.(bool)
		if !ok {
			service.RespInvalidParaError(ctx)
			return
		}
	}

	// 啥都不更新, 视为参数错误
	if !updateAnnoymous && !updateContent {
		service.RespInvalidParaError(ctx)
		return
	}

	var exists, permissionDeny bool

	ok := retry.RetryFrame(func() error {
		exists_, permissionDeny_, err := sercomment.TryUpdateComment(userID,
			commentID, updateAnnoymous, annoymous, updateContent, content)
		if err != nil {
			return err
		}

		exists = exists_
		permissionDeny = permissionDeny_

		return nil

	}, 3, "api/comment/update_comment.go UpdateComment.TryUpdateComment")

	if !ok {
		service.RespServiceNotAvailabelError(ctx)
		return
	}

	if !exists {
		sercomment.RespNoCommentToUpdate(ctx)
		return
	}

	if permissionDeny {
		sercomment.RespNoPermissionToUpdateComment(ctx)
		return
	}

	sercomment.RespUpdateCommentOK(ctx)
	// return
}
