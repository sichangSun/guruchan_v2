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

// query all food or querry collection list (favorite)
func (f *FoodRepository) QuerryFoodList(ctx context.Context, userID int, typeCode int8, pageInt int) (*mysql.FoodSlice, error) {
	var output mysql.FoodSlice
	err := error(nil)
	if typeCode == 0 {
		output, err = mysql.Foods(Where("userId = ?", userID), Where("isDel = ?", 0), Limit(30), Offset(pageInt)).All(ctx, f.db)
	} else {
		output, err = mysql.Foods(Where("userId = ?", userID), Where("isLikeFlag = ?", 1), Where("isDel = ?", 0), Limit(30), Offset(pageInt)).All(ctx, f.db)
	}
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, repository.RowsNotFoundErr
		}
		return nil, err
	}
	return &output, nil
}

// // querry collection list (favorite)
// func (f *FoodRepository) QuerryCollectionFoodList(ctx context.Context, userID int, typeCode int8, pageInt int) (*mysql.FoodSlice, error) {
// 	output, err := mysql.Foods(Where("userId = ?", userID), Where("isLikeFlag = ?", 1), Where("isDel = ?", 0), Limit(30), Offset(pageInt)).All(ctx, f.db)
// 	if err != nil {
// 		if errors.Is(err, sql.ErrNoRows) {
// 			return nil, repository.RowsNotFoundErr
// 		}
// 		return nil, err
// 	}
// 	return &output, nil
// }

// querry one food by foodID
func (f *FoodRepository) QuerryOneFood(ctx context.Context, userID int) (*mysql.FoodSlice, error) {
	return nil, nil
}

// fuzzy Query food
func (f *FoodRepository) FuzzyQuery(ctx context.Context, userID string, foodOrRestaurantName string, tagId []int8) (*mysql.FoodSlice, error) {
	// name only
	if foodOrRestaurantName != "" {
		if tagId != nil && len(tagId) > 0 {
			output, err := mysql.Foods(Where("userId = ?", userID),
				Where("restaurantName like ? or foodName like ? or address like ?", `%foodOrRestaurantName%`, `%foodOrRestaurantName%`, `%foodOrRestaurantName%`),
				Where("isDel = ?", 0),
				LeftOuterJoin("Tag on Tag."),
				//Todo DB syusei
			).All(ctx, f.db)
			if err != nil {
				if errors.Is(err, sql.ErrNoRows) {
					return nil, repository.RowsNotFoundErr
				}
				return nil, err
			}
			return &output, nil
		}

	}
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
