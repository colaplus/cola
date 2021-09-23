package common

type Category struct {
	CategoryId   uint64 `json:"category_id" db:"category_id"`
	CategoryName string `json:"category_name" db:"category_name"`
}
