package restaurantbusiness

import (
	"context"
	restaurantmodel "food-delivery/module/restaurant/model"
)

type GetRestaurantStore interface {
	GetRestaurant(
		ctx context.Context,
		cond map[string]interface{},
		moreKeys ...string,
	) (*restaurantmodel.Restaurant, error)
}

type getRestaurantByIdBusiness struct {
	store GetRestaurantStore
}

func NewGetRestaurantByIdBusiness(store GetRestaurantStore) *getRestaurantByIdBusiness {
	return &getRestaurantByIdBusiness{store: store}
}

func (biz *getRestaurantByIdBusiness) GetRestaurantById(ctx context.Context, id int) (*restaurantmodel.Restaurant, error) {
	restaurant, err := biz.store.GetRestaurant(ctx, map[string]interface{}{"id": id})
	if err != nil {
		return nil, err
	}
	return restaurant, nil
}
