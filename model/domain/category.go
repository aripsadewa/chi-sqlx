package domain

import (
	"time"
)

type Category struct {
	ID        int       `db:"id"`
	Name      string    `db:"name"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
}

// type CategoryMeta struct {
// 	StartDate null.Time   `json:"start_date"`
// 	EndtDate  null.Time   `json:"end_date"`
// 	Sort      null.String `json:"sort"`
// 	SortValue null.String `json:"sort_value"`
// 	Limit     null.Float  `db:"limit" json:"limit"`
// 	Total     null.Float  `db:"total" json:"total"`
// 	Page      null.Float  `db:"page" json:"page"`
// 	TotalPage null.Float  `db:"total_page" json:"total_page"`
// }

type CategoryMeta struct {
	Page  float64 `validate:"number" json:"page"`
	Limit float64 `validate:"number" json:"limit"`
	// Start     null.Time   `validate:"datetime" schema:"start"`
	// End       null.Time   `validate:"datetime" schema:"end"`
	// Sort      null.String `json:"sort" schema:"sort"`
	// SortValue null.String `json:"sort_value" schema:"sort_value"`
	Total     int `db:"total" json:"total"`
	TotalPage int `db:"total_page" json:"total_page"`
}
