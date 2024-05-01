package mysql

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"guruchan-back/app/domain/repository"
	"guruchan-back/app/domain/service"
	mysql "guruchan-back/app/infra/sqlboiler"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	. "github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type FoodRepository struct {
	db *sqlx.DB
}

// query all food
func (f *FoodRepository) QuerryFoodList(ctx context.Context, userID int) (*mysql.FoodSlice, error) {
	output, err := mysql.Foods(
		Select("f.restaurantName", "f.res_address", "f.isLikeflag", "f.typecode", "f.fullAddress", "f.f_addTime",
			"f.f_updataTime", "f.testedFlag", "f.foodImg", "f.foodId", "t.tagTittle", "t.tagId"), From("Food as f"),
		InnerJoin("FoodTags ft on ft.foodId = f.foodId"),
		InnerJoin("Tag t on t.tagId = ft.tagId"),
		Where("userId = ?", userID),
		Where("isDel = ?", 0)).All(ctx, f.db)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, repository.RowsNotFoundErr
		}
		return nil, err
	}
	return &output, nil
}

// querry collection list (favorite)
func (f *FoodRepository) QuerryCollectionFoodList(ctx context.Context, userID int, typeCode int8) (*mysql.FoodSlice, error) {
	//when typeCode is 1 ,find collection list (favorite)
	output, err := mysql.Foods(
		Select("f.restaurantName", "f.res_address", "f.isLikeflag", "f.typecode", "f.fullAddress", "f.f_addTime",
			"f.f_updataTime", "f.testedFlag", "f.foodImg", "f.foodId", "t.tagTittle", "t.tagId"), From("Food as f"),
		InnerJoin("FoodTags ft on ft.foodId = f.foodId"),
		InnerJoin("Tag t on t.tagId = ft.tagId"),
		Where("userId = ?", userID),
		Where("isDel = ?", 0), Where("f.isLikeflag = ?", typeCode)).All(ctx, f.db)
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
// TODO foodOrRestaurantName tobe "% %" in foodservice
func (f *FoodRepository) FuzzyQueryByName(ctx context.Context, userID int, foodOrRestaurantName string) (*mysql.FoodSlice, error) {
	if foodOrRestaurantName == "%%" {
		return nil, errors.New("mysql: no foodOrRestaurantName provided for insertion")
	} else {
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
}

// fuzzy Query food by tagIds
func (f *FoodRepository) FuzzyQueryByTagIDs(ctx context.Context, userID int, tagId []int) (*mysql.FoodSlice, error) {
	// just querry by tagId
	if len(tagId) == 0 {
		return nil, errors.New("mysql: no tagsIDs provided for insertion")
	} else {
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
	}
}

// fuzzy Query
func (f *FoodRepository) FuzzyQuery(ctx context.Context, userID int, foodOrRestaurantName string, tagId []int) (*mysql.FoodSlice, error) {
	//have both foodOrRestaurantName and tagId
	if len(tagId) == 0 || foodOrRestaurantName == "%%" {
		return nil, errors.New("mysql: no foodOrRestaurantName provided for insertion")
	} else {
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
	}
}

// update food info
// return 1 susscful  0 failed
func (f *FoodRepository) UpdateFoodById(ctx context.Context, userID int, food *service.FoodInput) (int8, error) {
	if food == nil {
		return 0, errors.New("mysql: no Food provided for insertion")
	} else {
		output, err := mysql.Foods(Where("foodID = ?", food.FoodId), Where("userId = ?", userID), Where("isDel = ?", 0)).One(ctx, f.db)
		if err != nil {
			//querry error
			return 0, err
		} else if output == nil {
			// can not find output
			return 0, errors.New("mysql: No matching rows in the database")
		} else {
			//update
			output.RestaurantName = food.RestaurantName
			output.FoodName = null.String(food.FoodName)
			output.ResAddress = null.String(food.ResAddress)
			output.FullAddress = null.String(food.FullAddress)
			output.IsLikeflag = food.IsLikeflag
			output.Typecode = food.Typecode
			output.FUpdataTime = null.TimeFrom(time.Now())
			output.TestedFlag = food.TestedFlag
			output.FoodImg = null.String(food.FoodImg)
			output.FoodDetail = null.String(food.FoodDetail)

			_, err := output.Update(ctx, f.db, boil.Infer())
			if err != nil {
				return 1, err
			} else {
				// if food.NewAddTags != nil && len(food.NewAddTags) > 0 {
				// 	err = updateFoodTags(ctx, f.db, food.FoodId, food.UserId, food.NewAddTags, true)
				// 	if err != nil {
				// 		return resultInt, fmt.Errorf("mysql: error adding tag - %w", err)
				// 	}
				// 	return resultInt + 2, nil
				// }
				// if food.DelTags != nil && len(food.DelTags) > 0 {
				// 	err = updateFoodTags(ctx, f.db, food.FoodId, food.UserId, food.DelTags, false)
				// 	if err != nil {
				// 		return resultInt, fmt.Errorf("mysql: error deleting tag - %w", err)
				// 	}
				// 	return resultInt + 2, nil
			}
			return 1, nil
		}
	}

}

// add a new food
func (f *FoodRepository) AddNewFood(ctx context.Context, userID int, food *service.FoodInput) (int8, error) {
	if food == nil {
		return 0, errors.New("mysql: no Food provided for insertion")
	} else {
		newFood := mysql.Food{
			UserId:         userID,
			RestaurantName: food.RestaurantName,
			FoodName:       food.FoodName,
			ResAddress:     food.ResAddress,
			FullAddress:    food.FullAddress,
			IsLikeflag:     food.IsLikeflag,
			Typecode:       food.Typecode,
			FAddTime:       null.TimeFrom(time.Now()),
			FirstTime:      null.TimeFrom(food.FirstTime.Time),
			TestedFlag:     food.TestedFlag,
			FoodImg:        food.FoodImg,
			IsDel:          0,
			FoodDetail:     food.FoodDetail,
		}
		err := newFood.Insert(ctx, f.db, boil.Infer())
		if err != nil {
			return 0, fmt.Errorf("mysql: failed to insert new food - %w", err)
		}
	}
	return 1, nil
}

// softDelete food by ID
func (f *FoodRepository) SoftDeleteFood(ctx context.Context, userID int, foodId string) (int8, error) {
	output, err := mysql.Foods(Where("foodID = ?", foodId), Where("userId = ?", userID), Where("isDel = ?", 0)).One(ctx, f.db)
	if err != nil {
		//querry err
		return 0, err
	} else if output != nil {
		//del
		output.IsDel = 1
		_, err := output.Update(ctx, f.db, boil.Infer())
		if err != nil {
			return 0, err
		} else {
			return 1, nil
		}
	} else {
		// can not find output
		return 0, errors.New("mysql: No matching rows in the database")
	}

}

// updateFoodTags add or delet
func (f *FoodRepository) UpdateFoodTags(ctx context.Context, foodID int, userID int, tags []int, isAdd bool) (int8, error) {
	if tags != nil && len(tags) > 0 {
		for _, tagID := range tags {
			foodTag := mysql.FoodTag{
				FoodId: foodID,
				TagId:  tagID,
				UserId: userID,
			}
			hasTag, err := querryFoodTags(ctx, f.db, foodID, userID, foodTag.TagId)
			if err != nil {
				return 0, err
			}
			if isAdd {
				if !hasTag {
					err = foodTag.Insert(ctx, f.db, boil.Infer())
					if err != nil {
						return 0, err
					}
				}
			} else {
				if hasTag {
					_, err = foodTag.Delete(ctx, f.db)
					if err != nil {
						return 0, err
					}
				}
			}
		}
	}
	return 1, nil
}

// querry foodTag
func querryFoodTags(ctx context.Context, db *sqlx.DB, foodID int, userID int, tagID int) (bool, error) {
	_, err := mysql.FoodTags(Where("foodId = ?", foodID), Where("userId = ?", userID), Where("tagId = ?", tagID)).One(ctx, db)
	if err != nil {
		return false, err
	}
	return true, nil
}
