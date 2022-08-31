package restaurantstorage

import (
	"context"
	restaurantmodel "food-delivery/module/restaurant/model"
)

func (s *sqlStore) GetRestaurantById(ctx context.Context, id int) (*restaurantmodel.Restaurant, error) {
	var restaurant restaurantmodel.Restaurant
	if err := s.db.Table(restaurantmodel.Restaurant{}.TableName()).Where("id = ?", id).First(&restaurant).Error; err != nil {
		return nil, err
	}
	return &restaurant, nil
}
