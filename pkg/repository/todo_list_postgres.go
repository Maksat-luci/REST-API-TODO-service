package repository

import (
	"fmt"

	todo "github.com/Maksat-luci/REST-API-TODO-service"
	"github.com/jmoiron/sqlx"
)

type TodoListPostgres struct {
	db *sqlx.DB
}

func NewTodoListPostgres(db *sqlx.DB) *TodoListPostgres {
	return &TodoListPostgres{db: db}
}

func (r *TodoListPostgres) Create(userID int, list todo.TodoList) (int, error) {
	// запускаем транзакцию
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var id int
	//  сооздаём sql запрос
	createListQuery := fmt.Sprintf("INSERT INTO %s (title, description) VALUES ($1, $2) RETURNING id", todoListsTable)
	// делаю запрос в базу данных далее сохраняем данные в переменную и используем эти данные для последующего запроса в базу данных 
	row := tx.QueryRow(createListQuery, list.Title, list.Description)
	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}
	// формируем второй запрос в базу данных 
	createUsersListQuery := fmt.Sprintf("INSERT INTO %s (user_id, list_id) VALUES ($1, $2)", usersListsTable)
	//отправляем второй запрос с данными первого запроса в одну логическую операцию
	_, err = tx.Exec(createUsersListQuery, userID, id)
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	return id , tx.Commit()
}

func (r *TodoListPostgres) GetAll(userId int) ([]todo.TodoList, error) {
	var lists []todo.TodoList
	// выполняем sql inner join в базу данных 
	query := fmt.Sprintf("SELECT tl.id, tl.title ,tl.description FROM %s tl INNER JOIN %s ul on tl.id = ul.list_id WHERE ul.user_id = $1",
	todoListsTable, usersListsTable)
	// получаем и отправляем данные
	err := r.db.Select(&lists, query, userId)
	
	return lists, err
}

func (r *TodoListPostgres) GetById(userId, listId int) (todo.TodoList, error) {
	var list todo.TodoList
	
	query := fmt.Sprintf(`SELECT tl.id, tl.title, tl.description FROM %s tl 
		INNER JOIN %s ul on tl.id = ul.list_id WHERE ul.user_id = $1 AND ul.list_id = $2`,
	todoListsTable, usersListsTable)

	err := r.db.Get(&list, query, userId, listId)

	return list, err
}