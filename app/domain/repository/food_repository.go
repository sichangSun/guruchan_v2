package repository

import (
	"context"
	"errors"
	"guruchan-back/app/domain/service"
	mysql "guruchan-back/app/infra/sqlboiler"
)

var (
	RowsNotFoundErr = errors.New("")
)

type FoodRepository interface {
	//query all food list
	QuerryFoodList(ctx context.Context, userID int, typeCode int8) (*mysql.FoodSlice, error)
	//querry one food by foodID
	QuerryOneFood(ctx context.Context, userID int, foodID int) (*mysql.FoodSlice, error)
	//fuzzy Query byTag
	FuzzyQuery(ctx context.Context, userID int, foodOrRestaurantName string, tagId []int) (*mysql.FoodSlice, error)
	//update food
	UpdateFoodById(ctx context.Context, userID int, food *service.FoodInput) (int64, error)
	//add a new food
	AddNewFood(ctx context.Context, userID int, food *service.FoodInput) error
	//softDelete food by ID
	SoftDeleteFood(ctx context.Context, userID int, foodId string) error
}
