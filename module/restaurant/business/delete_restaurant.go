package restaurantbiz

import (
	"context"
	"food-delivery/common"
	restaurantmodel "food-delivery/module/restaurant/model"
)

type DeleteRestaurantStore interface {
	Delete(context context.Context, id int) error
	FindDataWithCondition(
		context context.Context,
		condition map[string]interface{},
		moreKeys ...string,
	) (*restaurantmodel.Restaurant, error)
}

type deleteRestaurantBiz struct {
	store     DeleteRestaurantStore
	requester common.Requester
}

func NewDeleteRestaurantBiz(store DeleteRestaurantStore, requester common.Requester) *deleteRestaurantBiz {
	return &deleteRestaurantBiz{store: store, requester: requester}
}

func (biz *deleteRestaurantBiz) DeleteRestaurant(context context.Context, id int) error {

	oldData, err := biz.store.FindDataWithCondition(context, map[string]interface{}{"id": id})

	if err != nil {
		return common.ErrEntityNotFound(restaurantmodel.EntityName, err)
	}

	if oldData.Status == 0 {
		return common.ErrEntityDeleted(restaurantmodel.EntityName, nil)
	}

	if oldData.UserId != biz.requester.GetUserId() {
		return common.ErrNoPermission(nil)
	}

	if err := biz.store.Delete(context, id); err != nil {
		return common.ErrCannotDeleteEntity(restaurantmodel.EntityName, nil)
	}

	return nil
}
