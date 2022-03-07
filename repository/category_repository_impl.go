package repository

import (
	"context"
	"errors"
	"fmt"
	"rest_api/model"
	"rest_api/model/domain"
	"strings"

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
	query := "INSERT INTO category (name,description, created_at) VALUES (:name,:description, now())"

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
	query := "SELECT id,name, description FROM category WHERE id=?"
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
	//true/false
	affectedId, _ := rs.RowsAffected()
	if err != nil {
		return 0, err
	}
	if int(affectedId) == 0 {
		return 0, errors.New("data not found")
	}
	return categoryId, nil
}
func (r *CategoryRepositoryImpl) Update(ctx context.Context, category domain.Category) (*domain.Category, error) {
	var args []interface{}
	q, qArgs := r.generateFieldQuery(category)
	// query := fmt.Sprintf("SELECT id, name, description FROM category %s ORDER BY ? asc LIMIT ? OFFSET ? ", q)

	args = append(args, qArgs...)
	args = append(args, category.ID)
	query := fmt.Sprintf("UPDATE category %s, updated_at=now() WHERE id=?", q)
	fmt.Println("query : ", query)
	fmt.Println("args : ", args)

	rs := r.DB.MustExec(query, args...)

	insertId, err := rs.RowsAffected()
	fmt.Println("inserted id ", insertId)
	if err != nil {
		return nil, err
	}
	if int(insertId) == 0 {
		return nil, errors.New("data not found")
	}
	category.ID = int(insertId)

	return &category, nil
}

func (r *CategoryRepositoryImpl) FindData(ctx context.Context, filter domain.CategoryFilter, paginate model.PaginateParams) ([]*domain.Category, *model.PaginateParams, error) {
	var args []interface{}
	q, qArgs := r.generateWhereQuery(filter)

	args = append(args, qArgs...)
	args = append(args, filter.Sort, paginate.Limit, paginate.Offset)

	query := fmt.Sprintf("SELECT id, name, description FROM category %s ORDER BY ? asc LIMIT ? OFFSET ? ", q)
	fmt.Println("query : ", query)
	fmt.Println("args : ", args)
	categories := []*domain.Category{}
	err := r.DB.Select(&categories, query, args...)
	if err != nil {
		return nil, nil, err
	}
	count, err := r.getCountCategory(filter)
	if err != nil {
		return nil, nil, err
	}
	meta := model.PaginateParams{
		Total: count,
	}
	return categories, &meta, nil
}

func (r *CategoryRepositoryImpl) generateWhereQuery(filter domain.CategoryFilter) (q string, args []interface{}) {
	var condition []string
	if filter.StartDate.Valid {
		condition = append(condition, "created_at > ?")
		args = append(args, filter.StartDate.Time.Format("2006-02-01"))
	}
	if filter.EndDate.Valid {
		condition = append(condition, "created_at < ?")
		args = append(args, filter.EndDate.Time.Format("2006-01-02"))
	}
	if filter.Name.Valid {
		condition = append(condition, "name LIKE ?")
		args = append(args, fmt.Sprintf("%%%s%%", filter.Name.String))
	}
	if len(condition) > 0 {
		q = fmt.Sprintf("WHERE %s", strings.Join(condition, " and "))
	}
	return
}

func (r *CategoryRepositoryImpl) generateFieldQuery(cat domain.Category) (q string, args []interface{}) {
	var condition []string
	if cat.Name != "" {
		condition = append(condition, "name=?")
		args = append(args, cat.Name)
	}
	if cat.Description != "" {
		condition = append(condition, "description=?")
		args = append(args, cat.Description)
	}
	if len(condition) > 0 {
		q = fmt.Sprintf("SET %s", strings.Join(condition, " , "))
	}
	return
}

func (r *CategoryRepositoryImpl) getCountCategory(filter domain.CategoryFilter) (int, error) {
	var args []interface{}
	var count int
	q, qArgs := r.generateWhereQuery(filter)
	args = append(args, qArgs...)
	fmt.Printf("service %+v \n", args...)

	query := fmt.Sprintf("SELECT count(id) as total from category %s", q)
	err := r.DB.Get(&count, query, args...)
	if err != nil {
		return 0, err
	}
	return count, nil
}
