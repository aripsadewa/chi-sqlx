package repository

import (
	"context"
	"fmt"
	"math"
	"rest_api/model/domain"

	"github.com/jmoiron/sqlx"
)

type CategoryRepositoryImpl struct {
	DB *sqlx.DB
}

func NewCategoryRepository(DB *sqlx.DB) CategoryRepository {
	return &CategoryRepositoryImpl{
		DB: DB,
	}
}

func (r *CategoryRepositoryImpl) Save(ctx context.Context, category domain.Category) (*domain.Category, error) {
	query := "INSERT INTO category (name) VALUES (:name)"

	rs, err := r.DB.NamedExec(query, category)

	if err != nil {
		return nil, err
	}

	insertId, err := rs.LastInsertId()
	if err != nil {
		return nil, err
	}

	category.ID = int(insertId)

	return &category, nil
}

func (r *CategoryRepositoryImpl) Update(ctx context.Context, category domain.Category) (*domain.Category, error) {
	query := "UPDATE category SET name=:name WHERE id=:id"

	rs, err := r.DB.NamedExec(query, category)

	if err != nil {
		return nil, err
	}

	insertId, err := rs.RowsAffected()
	if err != nil {
		return nil, err
	}

	category.ID = int(insertId)

	return &category, nil
}

func (r *CategoryRepositoryImpl) FindById(ctx context.Context, categoryId int) (*domain.Category, error) {
	query := "SELECT * FROM category WHERE id=?"
	rs := domain.Category{}
	err := r.DB.Get(&rs, query, categoryId)
	if err != nil {
		return nil, err
	}

	return &rs, nil
}

func (r *CategoryRepositoryImpl) Delete(ctx context.Context, categoryId int) (int, error) {

	query := "DELETE FROM category WHERE id=?"
	rs, err := r.DB.Exec(query, categoryId)
	affectedId, _ := rs.RowsAffected()
	if err != nil || int(affectedId) == 0 {
		return 0, err
	}
	return categoryId, nil
}

func (r *CategoryRepositoryImpl) FindAll(ctx context.Context, parPage int) ([]*domain.Category, *domain.CategoryMeta, error) {

	if parPage > 1 {
		limit := 5
		q := "SELECT * FROM category"
		query := fmt.Sprintf("%s LIMIT %d OFFSET %d", q, limit, (parPage-1)*limit)
		categories := []*domain.Category{}
		err := r.DB.Select(&categories, query)
		if err != nil {
			return nil, nil, err
		}
		meta := domain.CategoryMeta{}
		var page float64 = meta.Total / float64(limit)
		total := math.Ceil(page)
		err = r.DB.Get(&meta.Total, "SELECT COUNT(id) AS total FROM category")
		if err != nil {
			return nil, nil, err
		}

		meta = domain.CategoryMeta{
			Total:     meta.Total,
			Page:      parPage,
			TotalPage: total,
		}
		return categories, &meta, nil
	} else {
		limit := 5
		q := "SELECT * FROM category"
		query := fmt.Sprintf("%s LIMIT %d OFFSET %d", q, limit, 0)
		categories := []*domain.Category{}
		err := r.DB.Select(&categories, query)
		if err != nil {
			return nil, nil, err
		}
		meta := domain.CategoryMeta{}
		err = r.DB.Get(&meta.Total, "SELECT COUNT(id) AS total FROM category")
		if err != nil {
			return nil, nil, err
		}

		var page float64 = meta.Total / float64(limit)
		total := math.Ceil(page)

		meta = domain.CategoryMeta{
			Total:     meta.Total,
			Page:      1,
			TotalPage: total,
		}

		return categories, &meta, nil
	}

}
