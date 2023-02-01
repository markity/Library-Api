package book

import (
	"library-api/service"
	serbook "library-api/service/book"
	fieldcheck "library-api/util/field_check"
	"library-api/util/retry"

	"github.com/gin-gonic/gin"
)

func SearchBooksByLabel(ctx *gin.Context) {
	json_, _ := ctx.Get("json")
	json := json_.(map[string]interface{})

	authed_, _ := ctx.Get("authed")
	authed := authed_.(bool)
	var userID int64 = -1
	if authed {
		userID = ctx.GetInt64("user_id")
	}

	label_, ok1 := json["label"]
	label, ok2 := label_.(string)
	if !ok1 || !ok2 || !fieldcheck.CheckLabelVaile(label) {
		service.RespInvalidParaError(ctx)
		return
	}

	var books []serbook.BookWithPariseAndFocus

	ok := retry.RetryFrame(func() error {
		b, err := serbook.TrySearchLabelBooks(userID, label)
		if err != nil {
			return err
		}

		books = b
		return nil
	}, 3, "api/book/search_by_label.go SearchBooksByLabel.TrySearchLabelBooks")

	if !ok {
		service.RespServiceNotAvailabelError(ctx)
		return
	}

	serbook.RespLableSearchOK(ctx, books)

}
