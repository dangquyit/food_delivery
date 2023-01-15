package restaurantbusiness

import "context"

type DeleteRestaurant interface {
	Delete(context context.Context, id int) error
}

type deleteRestaurant struct {
	store DeleteRestaurant
}

func NewDeleteRestaurantBusiness(store DeleteRestaurant) *deleteRestaurant {
	return &deleteRestaurant{store: store}
}

func (bsn *deleteRestaurant) DeleteRestaurant(context context.Context, id int) error {
	
	if err := bsn.store.Delete(context, id); err != nil {
		return err
	}
	return nil
}
