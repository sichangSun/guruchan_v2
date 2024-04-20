package mysql

import (
	"context"
	"database/sql"
	"errors"
	"guruchan-back/app/domain/repository"
	"guruchan-back/app/domain/service"
	mysql "guruchan-back/app/infra/sqlboiler"

	"github.com/jmoiron/sqlx"
	. "github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type FoodRepository struct {
	db *sqlx.DB
}

// query food(whit collection list or all food list ) typecode=1 is collection list ,
func (f *FoodRepository) QuerryFoodList(ctx context.Context, userID string, typeCode int8, pageInt int) (*mysql.FoodSlice, error) {
	output, err := mysql.Foods(Where("isDel = ?", 0), Limit(30), Offset(pageInt)).All(ctx, f.db)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, repository.RowsNotFoundErr
		}
		return nil, err
	}
	return &output, nil
}

// fuzzy Query food
func (FoodRepository *FoodRepository) FuzzyQuery(ctx context.Context, userID string, foodOrRestaurantName string, tagIds []int8, pageInt int) (*mysql.FoodSlice, error) {
	return nil, nil
}

// update food info
func (FoodRepository *FoodRepository) UpdateFoodById(ctx context.Context, userID string, food *service.FoodInput) error {
	return nil
}

// add a new food
func (FoodRepository *FoodRepository) AddNewFood(ctx context.Context, userID string, food *service.FoodInput) error {
	return nil
}

// softDelete food by ID
func (FoodRepository *FoodRepository) SoftDeleteFood(ctx context.Context, userID string, foodId string) error {
	return nil
}
