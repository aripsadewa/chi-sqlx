package domain

import "time"

type Category struct {
	ID        int       `db:"id"`
	Name      string    `db:"name"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
}

type CategoryMeta struct {
	Sort      string  `json:"sort"`
	SortValue string  `json:"sort_value"`
	Limit     float64 `db:"limit" json:"limit"`
	Total     float64 `db:"total" json:"total"`
	Page      int     `db:"page" json:"page"`
	TotalPage float64 `db:"total_page" json:"total_page"`
}
