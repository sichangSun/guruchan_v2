package service

type FoodService struct {
}
type FoodOutput struct {
}
type FoodInput struct {
}

// QuerryFoodList
func (service *FoodService) QuerryFoodList(userID string, typeCode string, pageInt int) (*FoodOutput, error) {
	//Todo format argument (default 0,query all food)
	//call interfece repository,create FoodOutput and return
	return nil, nil
}
