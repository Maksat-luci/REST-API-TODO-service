package todo

// User структура пользователя определяем его на самом верхнем уровне нашего приложения для удобства использования
type User struct {
	ID       int    `json:"-" db:"id"`
	Name     string `json:"name" binding:"required"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
