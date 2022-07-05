package todo

// User структура пользователя определяем его на самом верхнем уровне нашего приложения для удобства использования
type User struct {
	ID       int    `json:"-"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
}
