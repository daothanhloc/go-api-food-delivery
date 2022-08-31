package restaurantbusiness

import "context"

type DeleteRestaurantStore interface {
	DeleteRestaurant(ctx context.Context, id int) error
}

type deleteRestaurantBusiness struct {
	store DeleteRestaurantStore
}

func NewDeleteRestaurantBusiness(store DeleteRestaurantStore) *deleteRestaurantBusiness {
	return &deleteRestaurantBusiness{store: store}
}

func (biz *deleteRestaurantBusiness) DeleteRestaurant(ctx context.Context, id int) error {
	if err := biz.store.DeleteRestaurant(ctx, id); err != nil {
		return err
	}
	return nil
}
