package operate

import (
	"library-api/model"
	"library-api/service"
	"library-api/service/operate"
	"library-api/util/retry"

	"github.com/gin-gonic/gin"
)

func ListFocus(ctx *gin.Context) {
	userID := ctx.GetInt64("user_id")

	var result []*model.FocusListItem

	ok := retry.RetryFrame(func() error {
		res, err := operate.TryListFocus(userID)
		if err != nil {
			return err
		}

		result = res
		return nil
	}, 3, "api/operate/list_focus.go ListFocus.TryListFocus")

	if !ok {
		service.RespServiceNotAvailabelError(ctx)
		return
	}

	operate.RespListFocusOK(ctx, result)

}
