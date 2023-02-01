package operate

import (
	"library-api/api/book"
	"library-api/api/comment"
	"library-api/service"

	"github.com/gin-gonic/gin"
)

func Praise(ctx *gin.Context) {
	json_, _ := ctx.Get("json")
	json := json_.(map[string]interface{})

	model_, ok1 := json["model"]
	model, ok2 := model_.(float64)
	if !ok1 || !ok2 || (model != 1 && model != 0) {
		service.RespInvalidParaError(ctx)
		return
	}

	targetID_, ok3 := json["target_id"]
	targetID, ok4 := targetID_.(float64)
	if !ok3 || !ok4 {
		service.RespInvalidParaError(ctx)
		return
	}

	if model == 1 {
		book.Praise(ctx, int64(targetID))
		return
	}

	if model == 0 {
		comment.Praise(ctx, int64(targetID))
		return
	}
}
