package helper

import (
	"yogipratama/belajar-go-restful-api/model/entity"
	"yogipratama/belajar-go-restful-api/model/web"
)

func ToCategoryResponse(category entity.Category) web.CategoryResponse {
	return web.CategoryResponse{
		Id:   category.Id,
		Name: category.Name,
	}
}

func ToCategoryResponses(categories []entity.Category) []web.CategoryResponse {
	var categoryResponses []web.CategoryResponse
	for _, category := range categories {
		categoryResponses = append(categoryResponses, ToCategoryResponse(category))
	}

	return categoryResponses
}
