package restaurantstorage

import (
	"context"
	"food-delivery/common"
	restaurantmodel "food-delivery/module/restaurant/model"
)

func (s *sqlStore) ListRestaurant(ctx context.Context,
	filter *restaurantmodel.Filter,
	paging *common.Paging, moreKeys ...string) ([]restaurantmodel.Restaurant, error) {
	offset := (paging.Page - 1) * paging.Limit
	var result []restaurantmodel.Restaurant

	db := s.db
	if v := filter.OwnerId; v > 0 {
		db = db.Where("owner_id = >", v)
	}

	if err := db.Table(restaurantmodel.Restaurant{}.TableName()).Count(&paging.Total).Offset(offset).Limit(paging.Limit).Order("id desc").Find(&result).Error; err != nil {
		return nil, err
	}

	return result, nil
}
