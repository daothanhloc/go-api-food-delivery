package restaurantstorage

import (
	"context"
	restaurantmodel "food-delivery/module/restaurant/model"
)

func (s *sqlStore) GetRestaurant(ctx context.Context, cond map[string]interface{},
	moreKeys ...string) (*restaurantmodel.Restaurant, error) {
	var restaurant restaurantmodel.Restaurant
	if err := s.db.Table(restaurantmodel.Restaurant{}.TableName()).Where(cond).First(&restaurant).Error; err != nil {
		return nil, err
	}
	return &restaurant, nil
}
