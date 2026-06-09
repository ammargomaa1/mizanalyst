package pagination

import (
	"math"

	"gorm.io/gorm"
)

// PaginationMeta holds pagination metadata for API responses.
type PaginationMeta struct {
	CurrentPage int   `json:"current_page"`
	PerPage     int   `json:"per_page"`
	Total       int64 `json:"total"`
	LastPage    int   `json:"last_page"`
}

// PaginatedResult wraps data with pagination metadata.
type PaginatedResult struct {
	Data interface{}    `json:"data"`
	Meta PaginationMeta `json:"meta"`
}

// Paginate applies pagination to a GORM query scope. It counts total records,
// calculates the offset, and applies Limit/Offset to the query.
// Returns the scoped *gorm.DB (with limit/offset applied) and the PaginationMeta.
func Paginate(db *gorm.DB, page, perPage int) (*gorm.DB, PaginationMeta) {
	if page < 1 {
		page = 1
	}
	if perPage < 1 {
		perPage = 10
	}
	if perPage > 100 {
		perPage = 100
	}

	var total int64
	db.Count(&total)

	lastPage := int(math.Ceil(float64(total) / float64(perPage)))
	if lastPage < 1 {
		lastPage = 1
	}

	offset := (page - 1) * perPage

	meta := PaginationMeta{
		CurrentPage: page,
		PerPage:     perPage,
		Total:       total,
		LastPage:    lastPage,
	}

	return db.Offset(offset).Limit(perPage), meta
}
