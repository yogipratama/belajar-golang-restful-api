package repository

import (
	"context"
	"database/sql"
	"errors"
	"yogipratama/belajar-go-restful-api/helper"
	"yogipratama/belajar-go-restful-api/model/entity"
)

type CategoryRepositoryImpl struct {
}

func NewCategoryRepository() CategoryRepository {
	return &CategoryRepositoryImpl{}
}

func (repository *CategoryRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, category entity.Category) entity.Category {
	querySql := "INSERT INTO categories(name) values (?)"
	result, err := tx.ExecContext(ctx, querySql, category.Name)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	category.Id = int(id)
	return category
}

func (repository *CategoryRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, category entity.Category) entity.Category {
	querySql := "UPDATE categories set name = ? WHERE id = ?"
	_, err := tx.ExecContext(ctx, querySql, category.Name, category.Id)
	helper.PanicIfError(err)

	return category
}

func (repository *CategoryRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, category entity.Category) {
	querySql := "DELETE FROM categories WHERE id = ?"
	_, err := tx.ExecContext(ctx, querySql, category.Id)
	helper.PanicIfError(err)
}

func (repository *CategoryRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, categoryId int) (entity.Category, error) {
	querySql := "SELECT id, name FROM categories WHERE id = ?"
	category := entity.Category{}
	err := tx.QueryRowContext(ctx, querySql, categoryId).Scan(&category.Id, &category.Name)
	if err != nil {
		return category, errors.New("category is not found")
	}

	return category, nil
}

func (repository *CategoryRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []entity.Category {
	querySql := "SELECT id, name FROM categories"
	rows, err := tx.QueryContext(ctx, querySql)
	helper.PanicIfError(err)
	defer rows.Close()

	var categories []entity.Category
	for rows.Next() {
		category := entity.Category{}
		err := rows.Scan(&category.Id, &category.Name)
		helper.PanicIfError(err)

		categories = append(categories, category)
	}
	return categories
}
