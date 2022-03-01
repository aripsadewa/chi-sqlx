package domain

import (
	"gopkg.in/guregu/null.v4"
)

type Category struct {
	ID          int         `db:"id"`
	Name        string      `db:"name"`
	Description null.String `db:"description"`
}

type CategoryFilter struct {
	StartDate null.Time   `json:"start_date"`
	EndDate   null.Time   `json:"end_date"`
	Name      null.String `json:"name"`
	Sort      string      `json:"sort"`
	SortValue string      `json:"sort_value"`
}
