package todo

import (
	"errors"
	"fmt"
)

// TodoList структура
type TodoList struct {
	ID          int    `json:"id" db:"id"`
	Title       string `json:"title" db:"title" binding:"required"`
	Description string `json:"description" db:"description"`
}

// UsersList ...
type UsersList struct {
	ID     int
	UserID int
	ListID int
}

// TodoItem ...
type TodoItem struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Done        bool   `json:"done"`
}

// ListsItem ...
type ListsItem struct {
	ID     int
	ListID int
	ItemID int
}

type UpdateListInput struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
}
func (i UpdateListInput) Validate() error {
	fmt.Println(i.Description, "asd", i.Title)
	if i.Title == nil && i.Description == nil {
		return errors.Nels
		w("update structure has no values")
	}
	return nil
}