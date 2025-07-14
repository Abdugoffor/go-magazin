package category_dto

import (
	"category-crud/helper"
	category_model "category-crud/module/category/model"
)

type Create struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
}

type Update struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
}

type Response struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

type ResponseWithProducts struct {
	ID   uint
	Name string `json:"name"`
	// Description string `json:"description"`
}

func ToResponse(cat category_model.Category) Response {
	return Response{
		ID:          cat.ID,
		Name:        cat.Name,
		Description: cat.Description,
		CreatedAt:   helper.FormatDate(cat.CreatedAt),
		UpdatedAt:   helper.FormatDate(cat.UpdatedAt),
	}
}
