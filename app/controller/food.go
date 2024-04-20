package controller

import (
	"guruchan-back/app/domain/service"
)

type FoodController struct {
	FoodService *service.FoodService
}

// PagingQuerry to select all foodlist (max 30 in one page)
func (FoodController *FoodController) PagingQuerryAllFood() {
}
