package restaurantbusiness

import (
	"context"
	restaurantmodel "food-delivery/module/restaurant/model"
)

type GetRestaurantByIdStore interface {
	GetRestaurantById(ctx context.Context, id int) (*restaurantmodel.Restaurant, error)
}

type getRestaurantByIdBusiness struct {
	store GetRestaurantByIdStore
}

func NewGetRestaurantByIdBusiness(store GetRestaurantByIdStore) *getRestaurantByIdBusiness {
	return &getRestaurantByIdBusiness{store: store}
}

func (biz *getRestaurantByIdBusiness) GetRestaurantById(ctx context.Context, id int) (*restaurantmodel.Restaurant, error) {
	restaurant, err := biz.store.GetRestaurantById(ctx, id)
	if err != nil {
		return nil, err
	}
	return restaurant, nil
}
