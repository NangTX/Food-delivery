package rstlikebiz

import (
	"context"
	restaurantlikemodel "food-delivery/module/restaurantlike/model"
)

type UserDislikeRestaurantStore interface {
	Delete(ctx context.Context, UserId, restaurantId int) error
}

type userDislikeRestaurantBiz struct {
	store UserDislikeRestaurantStore
}

func NewUserDislikeRestaurantBiz(store UserDislikeRestaurantStore,
) *userDislikeRestaurantBiz {
	return &userDislikeRestaurantBiz{store: store}
}

func (biz *userDislikeRestaurantBiz) DislikeRestaurant(
	ctx context.Context,
	UserId,
	restaurantId int,
) error {
	err := biz.store.Delete(ctx, UserId, restaurantId)

	if err != nil {
		return restaurantlikemodel.ErrCannotDislikeRestaurant(err)
	}

	return nil
}
