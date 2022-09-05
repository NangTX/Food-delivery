package restaurantmodel

type Filter struct {
	OnwerId int   `json:"onwer_id,omitempty" form:"onwer_id"`
	Status  []int `json:"-"`
}
