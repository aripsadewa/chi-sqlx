package repository

import (
	"context"
	"rest_api/model/domain"
	"rest_api/web"

	"gorm.io/gorm"
)

type CategoryRepositoryImpl struct {
	DB *gorm.DB
}

func NewCategoryRepository(DB *gorm.DB) CategoryRepository {
	return &CategoryRepositoryImpl{
		DB: DB,
	}
}

func (r *CategoryRepositoryImpl) Save(ctx context.Context, category domain.Category) (*domain.Category, error) {
	err := r.DB.Create(&category).Error
	if err != nil {
		return nil, err
	}

	return &category, nil
}

func (r *CategoryRepositoryImpl) FindById(ctx context.Context, categoryId int) (*domain.Category, error) {
	rs := domain.Category{}
	err := r.DB.Where("id = ?", categoryId).First(&rs).Error
	if err != nil {
		return nil, err
	}
	return &rs, nil
}

func (r *CategoryRepositoryImpl) Delete(ctx context.Context, categoryId int) (int, error) {
	rs := domain.Category{}
	err := r.DB.Delete(&rs, categoryId).Error
	if err != nil {
		return 0, err
	}
	return categoryId, nil
}

func (r *CategoryRepositoryImpl) Update(ctx context.Context, category domain.Category) (*domain.Category, error) {
	err := r.DB.Model(&category).Where("id", category.ID).Updates(&category).Error
	if err != nil {
		return nil, err
	}

	return &category, nil
}

func (r *CategoryRepositoryImpl) FindData(ctx context.Context, filter domain.CategoryFilter, paginate *web.PaginateMetaData) ([]*domain.Category, error) {
	categories := []*domain.Category{}

	// query := r.DB.Debug()
	query := r.generateWhereQuery(filter).Debug()
	err := query.Limit(int(paginate.Limit)).Offset(paginate.Offset).Find(&categories).Error
	if err != nil {
		return nil, err
	}
	return categories, nil
	// var args []interface{}
	// q, qArgs := r.generateWhereQuery(filter)

	// args = append(args, qArgs...)
	// args = append(args, filter.Sort, paginate.Limit, paginate.Offset)

	// query := fmt.Sprintf("SELECT id, name, description FROM category %s ORDER BY ? asc LIMIT ? OFFSET ? ", q)
	// categories := []*domain.Category{}
	// err := r.DB.Select(&categories, query, args...)
	// if err != nil {
	// 	return nil, err
	// }
	// return categories, nil
}

func (r *CategoryRepositoryImpl) generateWhereQuery(filter domain.CategoryFilter) *gorm.DB {
	query := r.DB
	if filter.Name.Valid {
		query = query.Where("name LIKE ?", "%"+filter.Name.String+"%")
	}
	if filter.StartDate.Valid {
		query = query.Where("created_at > ?", filter.StartDate.Time)
	}
	if filter.EndDate.Valid {
		query = query.Where("created_at < ?", filter.EndDate.Time)
	}

	// if len(condition) > 0 {
	// 	q = fmt.Sprintf("WHERE %s", strings.Join(condition, " and "))
	// }
	return query
}

// func (r *CategoryGormRepositoryImpl) generateFieldQuery(cat domain.Category) (q string, args []interface{}) {
// 	var condition []string
// 	if cat.Name != "" {
// 		condition = append(condition, "name=?")
// 		args = append(args, cat.Name)
// 	}
// 	if cat.Description.Valid {
// 		condition = append(condition, "description=?")
// 		args = append(args, cat.Description)
// 	}
// 	if len(condition) > 0 {
// 		q = fmt.Sprintf("SET %s", strings.Join(condition, " , "))
// 	}
// 	return
// }

func (r *CategoryRepositoryImpl) GetCountCategory(filter domain.CategoryFilter) (int64, error) {
	var count int64
	query := r.generateWhereQuery(filter).Debug()
	err := query.Model(&domain.Category{}).Count(&count).Error
	if err != nil {
		return 0, err
	}

	return count, nil
}
