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
// typeCode 0 all, 1 just collection
func (f *FoodRepository) QuerryFoodList(ctx context.Context, userID int, typeCode int8) (*mysql.FoodSlice, error) {
	var output mysql.FoodSlice
	err := error(nil)
	if typeCode == 0 {
		//when typeCode is 0 ,find all food
		output, err = mysql.Foods(
			Select("f.restaurantName", "f.res_address", "f.isLikeflag", "f.typecode", "f.fullAddress", "f.f_addTime",
				"f.f_updataTime", "f.testedFlag", "f.foodImg", "f.foodId", "t.tagTittle", "t.tagId"), From("Food as f"),
			InnerJoin("FoodTags ft on ft.foodId = f.foodId"),
			InnerJoin("Tag t on t.tagId = ft.tagId"),
			Where("userId = ?", userID),
			Where("isDel = ?", 0)).All(ctx, f.db)
	} else {
		//when typeCode is not 0 ,find collection list (favorite)
		output, err = mysql.Foods(
			Select("f.restaurantName", "f.res_address", "f.isLikeflag", "f.typecode", "f.fullAddress", "f.f_addTime",
				"f.f_updataTime", "f.testedFlag", "f.foodImg", "f.foodId", "t.tagTittle", "t.tagId"), From("Food as f"),
			InnerJoin("FoodTags ft on ft.foodId = f.foodId"),
			InnerJoin("Tag t on t.tagId = ft.tagId"),
			Where("userId = ?", userID),
			Where("isDel = ?", 0), Where("f.isLikeflag = ?", 1)).All(ctx, f.db)
	}
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, repository.RowsNotFoundErr
		}
		return nil, err
	}
	return &output, nil
}

// querry one food by foodID
func (f *FoodRepository) QuerryOneFood(ctx context.Context, userID int, foodID int) (*mysql.FoodSlice, error) {
	output, err := mysql.Foods(
		Select("f.restaurantName", "f.res_address", "f.isLikeflag", "f.typecode", "f.fullAddress", "f.f_addTime",
			"f.f_updataTime", "f.testedFlag", "f.foodImg", "f.foodId", "t.tagTittle", "t.tagId"), From("Food as f"),
		InnerJoin("FoodTags ft on ft.foodId = f.foodId"),
		InnerJoin("Tag t on t.tagId = ft.tagId"),
		Where("userId = ?", userID),
		Where("isDel = ?", 0), Where("f.foodID = ?", foodID)).All(ctx, f.db)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, repository.RowsNotFoundErr
		}
		return nil, err
	}
	return &output, nil
}

// fuzzy Query food
// TODO foodOrRestaurantName tobe plus % in service
func (f *FoodRepository) FuzzyQuery(ctx context.Context, userID string, foodOrRestaurantName string, tagId []int) (*mysql.FoodSlice, error) {
	if foodOrRestaurantName != "" {
		//have both foodOrRestaurantName and tagId
		if len(tagId) > 0 {
			output, err := mysql.Foods(
				Select("f.restaurantName", "f.res_address", "f.isLikeflag", "f.typecode", "f.fullAddress", "f.f_addTime",
					"f.f_updataTime", "f.testedFlag", "f.foodImg", "f.foodId", "t.tagTittle", "t.tagId"), From("Food as f"),
				InnerJoin("FoodTags ft on ft.foodId = f.foodId"),
				InnerJoin("Tag t on t.tagId = ft.tagId"),
				Where("userId = ?", userID), Where("isDel = ?", 0),
				Where("AND (f.restaurantName LIKE ? OR f.foodName LIKE ? OR f.address LIKE ? ", foodOrRestaurantName, foodOrRestaurantName, foodOrRestaurantName),
				OrIn("id IN ?", tagId)).All(ctx, f.db)
			if err != nil {
				if errors.Is(err, sql.ErrNoRows) {
					return nil, repository.RowsNotFoundErr
				}
				return nil, err
			}
			return &output, nil
		} else {
			//don't have tag, just foodOrRestaurantName only
			output, err := mysql.Foods(
				Select("f.restaurantName", "f.res_address", "f.isLikeflag", "f.typecode", "f.fullAddress", "f.f_addTime",
					"f.f_updataTime", "f.testedFlag", "f.foodImg", "f.foodId", "t.tagTittle", "t.tagId"), From("Food as f"),
				InnerJoin("FoodTags ft on ft.foodId = f.foodId"),
				InnerJoin("Tag t on t.tagId = ft.tagId"),
				Where("userId = ?", userID), Where("isDel = ?", 0),
				Where("AND (f.restaurantName LIKE ? OR f.foodName LIKE ? OR f.address LIKE ? ",
					foodOrRestaurantName, foodOrRestaurantName, foodOrRestaurantName)).All(ctx, f.db)
			if err != nil {
				if errors.Is(err, sql.ErrNoRows) {
					return nil, repository.RowsNotFoundErr
				}
				return nil, err
			}
			return &output, nil
		}
	} else {
		//just querry by tagId
		if len(tagId) > 0 {
			output, err := mysql.Foods(
				Select("f.restaurantName", "f.res_address", "f.isLikeflag", "f.typecode", "f.fullAddress", "f.f_addTime",
					"f.f_updataTime", "f.testedFlag", "f.foodImg", "f.foodId", "t.tagTittle", "t.tagId"), From("Food as f"),
				InnerJoin("FoodTags ft on ft.foodId = f.foodId"),
				InnerJoin("Tag t on t.tagId = ft.tagId"),
				Where("userId = ?", userID), Where("isDel = ?", 0),
				WhereIn("id IN ?", tagId)).All(ctx, f.db)
			if err != nil {
				if errors.Is(err, sql.ErrNoRows) {
					return nil, repository.RowsNotFoundErr
				}
				return nil, err
			}
			return &output, nil
		} else {
			// TODO to change error description
			return nil, errors.New("error")
		}
	}
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
