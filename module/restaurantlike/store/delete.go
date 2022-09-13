package restaurantlikestorage

import (
	"context"
	"food-delivery/common"
	restaurantlikemodel "food-delivery/module/restaurantlike/model"
)

func (s *sqlStore) Delete(context context.Context, userId int, restaurantId int) error {

	db := s.db

	if err := db.Table(restaurantlikemodel.Like{}.TableName()).
		Where("user_id = ? and restaurant_id = ?", userId, restaurantId).
		Delete(map[string]interface{}{"status": 0}).
		Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
