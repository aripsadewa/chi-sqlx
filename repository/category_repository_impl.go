package repository

import (
	"context"
	"fmt"
	"math"
	"os"
	"rest_api/model/domain"
	"rest_api/web"

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

func (r *CategoryRepositoryImpl) FindAll(ctx context.Context, request web.GetParamRequest) ([]*domain.Category, *domain.CategoryMeta, error) {
	var page float64 = 1
	if request.Page.Valid {
		page = request.Page.Float64
	}
	query, count, totalPage := r.cekParam(request)
	meta := domain.CategoryMeta{}
	meta = domain.CategoryMeta{
		Limit:     request.Limit.Float64,
		Total:     count,
		Page:      page,
		TotalPage: totalPage,
	}
	// fmt.Println("query ", query)
	categories := []*domain.Category{}
	err := r.DB.Select(&categories, query)
	if err != nil {
		return nil, nil, err
	}

	return categories, &meta, nil
}

func (r *CategoryRepositoryImpl) cekParam(request web.GetParamRequest) (string, int, int) {
	count, _ := r.getCountCategory("select count(id) from category")
	q := "select id,name from category"
	startDate, sort, query, queryCount := "", "id", "", ""
	limit := 5
	page := 1
	totalPage := math.Ceil(float64(count) / float64(limit))
	var offset float64 = (float64(page) - 1) * float64(limit)
	sortValue := getValueEnv("SORT_CATEGORY_VALUE", "desc")
	if request.Start.Valid {
		if !request.End.Valid {
			var star string = request.Start.Time.String()
			t := star[0:10]
			startDate = fmt.Sprintf("where created_at > '%s 00:00:00'", t)
			queryCount = fmt.Sprintf("select count(id) from category where created_at > '%s 00:00:00'", t)
			count, _ = r.getCountCategory(queryCount)

		} else {
			var star string = request.Start.Time.String()
			var end string = request.End.Time.String()
			t := star[0:10]
			t2 := end[0:10]
			startDate = fmt.Sprintf("where created_at > '%s 00:00:00' and created_at < '%s 00:00:00'", t, t2)
			queryCount = fmt.Sprintf("select count(id) from category where created_at > '%s 00:00:00' and created_at < '%s 00:00:00'", t, t2)
			count, _ = r.getCountCategory(queryCount)

		}
	}
	if request.Limit.Valid || request.Page.Valid {
		if request.Limit.Valid && !request.Page.Valid {
			limit = int(request.Limit.Float64)
			totalPage = math.Ceil(float64(count) / request.Limit.Float64)
			offset = float64((page - 1) * limit)
		}
		if !request.Limit.Valid && request.Page.Valid {
			page = int(request.Page.Float64)
			offset = float64((page - 1) * limit)
		}
		limit = int(request.Limit.Float64)
		page = int(request.Page.Float64)
		offset = float64((request.Page.Float64 - 1) * request.Limit.Float64)
	}
	if request.Sort.Valid {
		sort = fmt.Sprintf(request.Sort.String)
	}
	if request.SortValue.Valid {
		sortValue = fmt.Sprintf(request.SortValue.String)
	}

	query = fmt.Sprintf("%s %s Order by %s %s LIMIT %d OFFSET %v", q, startDate, sort, sortValue, limit, offset)

	return query, count, int(totalPage)
}

func (r *CategoryRepositoryImpl) getCountCategory(q string) (int, error) {
	count := domain.CategoryMeta{}
	err := r.DB.Get(&count.Total, q)
	if err != nil {
		return 0, err
	}
	return count.Total, nil
}

func getValueEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
