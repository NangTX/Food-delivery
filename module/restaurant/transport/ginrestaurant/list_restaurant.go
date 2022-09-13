package ginrestaurant

import (
	"food-delivery/common"
	"food-delivery/component/appctx"
	restaurantbiz "food-delivery/module/restaurant/business"
	restaurantmodel "food-delivery/module/restaurant/model"
	restaurantrepo "food-delivery/module/restaurant/repository"
	restaurantstorage "food-delivery/module/restaurant/storage"

	"net/http"

	"github.com/gin-gonic/gin"
)

func ListRestaurant(appCtx appctx.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		db := appCtx.GetMainDBConnection()

		var pagingData common.Paging

		if err := c.ShouldBind(&pagingData); err != nil {

			panic(common.ErrInvalidRequest(err))
		}

		pagingData.Fulfill()

		var filter restaurantmodel.Filter
		if err := c.ShouldBind(&filter); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		if err := c.ShouldBind(&pagingData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		filter.Status = []int{1}

		store := restaurantstorage.NewSQLStore(db)
		// likeStore := restaurantlikestorage.NewSQLStore(db)
		repo := restaurantrepo.NewListRestaurantRepo(store)
		biz := restaurantbiz.NewListRestaurantBiz(repo)

		result, err := biz.ListRestaurant(c.Request.Context(), &filter, &pagingData)
		if err != nil {
			panic(err)
		}

		for i := range result {
			result[i].Mask(false)
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(result, pagingData, filter))
	}

}
