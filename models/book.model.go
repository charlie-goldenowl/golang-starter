package models

type Book struct {
	BaseModel
	Title  string `json:"title"`
	Author string `json:"author"`
}
type CreateBookInput struct {
	Title  string `json:"title" binding:"required"`
	Author string `json:"author" binding:"required"`
}
type UpdateBookInput struct {
	Title  string `json:"title"`
	Author string `json:"author"`
}
