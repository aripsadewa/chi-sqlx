package domain

type Category struct {
	ID   int    `db:"id"`
	Name string `db:"name"`
}

type CategoryMeta struct {
	Total     float64 `db:"total" json:"total"`
	Page      int     `db:"page" json:"page"`
	TotalPage float64 `db:"total_page" json:"total_page"`
}
