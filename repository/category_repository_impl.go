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
	query := "INSERT INTO category (name, created_at) VALUES (:name, now())"

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

func (r *CategoryRepositoryImpl) FindById(ctx context.Context, categoryId int) (*domain.Category, error) {
	query := "SELECT id,name FROM category WHERE id=?"
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
func (r *CategoryRepositoryImpl) Update(ctx context.Context, category domain.Category) (*domain.Category, error) {
	query := "UPDATE category SET name=:name, updated_at=now() WHERE id=:id"

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

func (r *CategoryRepositoryImpl) FindAll(ctx context.Context, request domain.CategoryMeta) ([]*domain.Category, *domain.CategoryMeta, error) {
	sort := request.Sort
	order := ""
	sortName := ""
	sortValue := ""
	limit := request.Limit
	var parPage float64 = float64(request.Page)
	var ofset float64 = 0
	meta := domain.CategoryMeta{}
	err := r.DB.Get(&meta.Total, "SELECT COUNT(id) AS total FROM category")
	if err != nil {
		return nil, nil, err
	}
	var page float64 = meta.Total / float64(limit)
	total := math.Ceil(page)
	if parPage > 1 {
		ofset = (parPage - 1) * limit
	} else {
		parPage = 1
		ofset = 0
	}

	meta = domain.CategoryMeta{
		Limit:     limit,
		Total:     meta.Total,
		Page:      int(parPage),
		TotalPage: total,
	}
	q := "SELECT id,name FROM category"
	if sort != "" {
		order = "ORDER BY"
		sortName = sort
		sortValue = request.SortValue
	}
	query := fmt.Sprintf("%s %s %s %s LIMIT %d OFFSET %d", q, order, sortName, sortValue, int(limit), int(ofset))
	categories := []*domain.Category{}
	err = r.DB.Select(&categories, query)
	if err != nil {
		return nil, nil, err
	}

	return categories, &meta, nil

}
