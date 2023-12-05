package main

import (
	"fmt"
	"net/http"
	"yogipratama/belajar-go-restful-api/app"
	"yogipratama/belajar-go-restful-api/controller"
	"yogipratama/belajar-go-restful-api/helper"
	"yogipratama/belajar-go-restful-api/middleware"
	"yogipratama/belajar-go-restful-api/repository"
	"yogipratama/belajar-go-restful-api/service"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db := app.NewDB()
	validate := validator.New()
	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)

	router := app.NewRouter(categoryController)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: middleware.NewAuthMiddleware(router),
	}
	fmt.Print("App is running...")

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
