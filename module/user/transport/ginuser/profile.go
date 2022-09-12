package ginuser

import (
	"food-delivery/common"
	"food-delivery/component/appctx"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Profile(appCtx appctx.AppContext) func(*gin.Context) {
	return func(c *gin.Context) {
		u := c.MustGet(common.CurrentUser).(common.Requester)

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(u))
	}
}
