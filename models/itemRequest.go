package models

// ItemRequest xx
type ItemRequest struct {
	Title       string `binding:"required"`
	Description string `binding:"required"`
	Picture     string `binding:"required"`
}
