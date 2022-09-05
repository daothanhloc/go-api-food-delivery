package restaurantbusiness

import (
	"context"
	"food-delivery/common"
	restaurantmodel "food-delivery/module/restaurant/model"
)

type CreateRestaurantStore interface {
	CreateRestaurant(ctx context.Context, data *restaurantmodel.RestaurantCreate) error
}

func NewCreateRestaurantBusiness(store CreateRestaurantStore) *createRestaurantBusiness {
	return &createRestaurantBusiness{store: store}
}

type createRestaurantBusiness struct {
	store CreateRestaurantStore
}

func (biz *createRestaurantBusiness) CreateRestaurant(ctx context.Context, data *restaurantmodel.RestaurantCreate) error {

	if err := data.Validate(); err != nil {
		return err
	}
	if err := biz.store.CreateRestaurant(ctx, data); err != nil {
		return common.NewErrorResponse(err, "create restaurant failed", "", "DB_ERROR")
	}
	return nil
}
