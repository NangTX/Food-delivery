package restaurantlikemodel

type Filter struct {
	RestaurantId int `json:"-" form:"restaurant_id"`
	UserId       int `json:"user_id" form:"user_id"`
}
