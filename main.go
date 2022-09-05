package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

type Restaurant struct {
	Id   int    `json:"id" gorm:"column:id;"`
	Name string `json:"name" gorm:"column:name;"`
	Addr string `json:"addr" gorm:"column:addr;"`
}

// trả về tên bảng trong DB
func (Restaurant) TableName() string {
	return "restaurants"
}

type RestaurantUpdate struct {
	Name *string `json:"name" gorm:"column:name;"`
	Addr *string `json:"addr" gorm:"column:addr;"`
}

func (RestaurantUpdate) TableName() string {
	return Restaurant{}.TableName()
}

func main() {
	// Lay bien env tu file .env
	envErr := godotenv.Load(".env")
	if envErr != nil {
		fmt.Println("Khong the mo file .env")
		os.Exit(1)
	}

	dsn := os.Getenv("MYSQL_CONN_STRING")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(db)

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	v1 := r.Group("/v1")
	restaurants := v1.Group("/restaurants")

	restaurants.POST("/", func(c *gin.Context) {
		var data Restaurant

		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		db.Create(&data)

		c.JSON(http.StatusOK, gin.H{
			"message": data,
		})
	})

	restaurants.GET("/:id", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}

		var data Restaurant

		db.Where("id = ?", id).First(&data)

		c.JSON(http.StatusOK, gin.H{
			"message": data,
		})
	})

	restaurants.GET("/", func(c *gin.Context) {

		var data []Restaurant

		type Paging struct {
			Page  int `json:"page" form:"page"`
			Limit int `json:"limit" form:"limit"`
		}

		var pagingData Paging

		if err := c.ShouldBind(&pagingData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		if pagingData.Page <= 0 {
			pagingData.Page = 1
		}

		if pagingData.Limit <= 0 {
			pagingData.Limit = 5
		}
		// có phân trang
		db.Offset((pagingData.Page - 1) * pagingData.Limit).
			Order("id desc").
			Limit(pagingData.Limit).
			Find(&data)

		c.JSON(http.StatusOK, gin.H{
			"message": data,
		})
	})

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

		db.Where("id = ?", id).Updates(&data)

		c.JSON(http.StatusOK, gin.H{
			"message": data,
		})
	})

	restaurants.DELETE("/:id", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}

		var data RestaurantUpdate

		db.Table(Restaurant{}.TableName()).Where("id = ?", id).Delete(nil)

		c.JSON(http.StatusOK, gin.H{
			"message": data,
		})
	})

	r.Run()
	// create database
	// newRestaurant := Restaurant{Name: "Tani", Addr: "281 truong dinh"}

	// if err := db.Create(&newRestaurant).Error; err != nil {
	// 	log.Println(err)
	// }
	//log.Println("New id :", newRestaurant.Id)

	// query database
	// var myRestaurant Restaurant

	// if err := db.Where("id= ?", 3).First(&myRestaurant).Error; err != nil {
	// 	log.Println(err)
	// }

	// log.Println(myRestaurant)

	// update database
	// myRestaurant.Name = "highland"

	// if err := db.Where("id= ?", 3).Updates(&myRestaurant).Error; err != nil {
	// 	log.Println(err)
	// }

	// log.Println(myRestaurant)

	// update database recommend
	// newName := "Highland"
	// updateData := RestaurantUpdate{Name: &newName}
	// if err := db.Where("id= ?", 3).Updates(&updateData).Error; err != nil {
	// 	log.Println(err)
	// }

	// log.Println(myRestaurant)

	//  delete database

	// if err := db.Table(Restaurant{}.TableName()).Where("id= ?", 1).Delete(nil).Error; err != nil {
	// 	log.Println(err)
	// }

}
