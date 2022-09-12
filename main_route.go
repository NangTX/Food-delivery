package main

import (
	"food-delivery/common"
	"food-delivery/component/appctx"
	"food-delivery/middleware"
	"food-delivery/module/restaurant/transport/ginrestaurant"
	"food-delivery/module/upload/transport/ginupload"
	"food-delivery/module/user/transport/ginuser"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func setupRoute(appContext appctx.AppContext, v1 *gin.RouterGroup) {
	v1.POST("/upload", ginupload.UploadImage(appContext))

	v1.POST("/register", ginuser.Register(appContext))
	v1.POST("/authenticate", ginuser.Login(appContext))
	v1.GET("/profile", middleware.RequiredAuth(appContext), ginuser.Profile(appContext))

	restaurants := v1.Group("/restaurants", middleware.RequiredAuth(appContext))

	restaurants.POST("/", ginrestaurant.CreateRestaurant(appContext))

	restaurants.GET("/:id", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}

		var data Restaurant

		appContext.GetMainDBConnection().Where("id = ?", id).First(&data)

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	})

	restaurants.GET("/", ginrestaurant.ListRestaurant(appContext))

	restaurants.PATCH("/:id", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}

		var data RestaurantUpdate

		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		appContext.GetMainDBConnection().Where("id = ?", id).Updates(&data)

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	})

	restaurants.DELETE("/:id", ginrestaurant.DeleteRestaurant(appContext))

}
