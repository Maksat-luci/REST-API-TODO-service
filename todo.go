package todo

// TodoList структура 
type TodoList struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
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
