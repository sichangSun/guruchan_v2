package service

import "github.com/volatiletech/null"

type FoodService struct {
}
type FoodOutput struct {
}
type FoodInput struct {
	FoodId         int
	UserId         int
	RestaurantName string
	FoodName       null.String
	ResAddress     null.String
	FullAddress    null.String
	IsLikeflag     int
	Typecode       int
	FAddTime       null.Time
	FUpdataTime    null.Time
	FirstTime      null.Time
	TestedFlag     int
	FoodImg        null.String
	IsDel          int
	FoodDetail     null.String
	OldTags        []int
	NewAddTags     []int
	DelTags        []int
}

// QuerryFoodList
func (service *FoodService) QuerryFoodList(userID string, typeCode string, pageInt int) (*FoodOutput, error) {
	//Todo format argument (default 0,query all food)
	//call interfece repository,create FoodOutput and return
	return nil, nil
}
