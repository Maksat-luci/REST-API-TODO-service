package todo

import (
	"context"
	"net/http"
	"time"
)

// Server структура  хендлер
type Server struct {
	httpServer *http.Server
}
// Run метод нашей структуры которая настраивает нашу структуру и запускает бесконечный цикл(горутины) для прослушивания URL адресов
func (s *Server) Run(port string) error {
	// создаём обьект структуры сервер настраиваем его и готовим к прослушиванию
	s.httpServer = &http.Server{
		Addr:           ":" + port,
		MaxHeaderBytes: 1 << 20, // 1 mb
		ReadTimeout:    10 * time.Second,
		WriteTimeout: 10 *time.Second,
	}
	// прослушиваем 
	return s.httpServer.ListenAndServe()
}

// GracefullShutdown метод который позволяет безопасно завершать работу приложения
func (s *Server) GracefullShutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
