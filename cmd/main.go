package main

import (
	"fmt"
	"log"

	todo "github.com/Maksat-luci/REST-API-TODO-service"
	"github.com/Maksat-luci/REST-API-TODO-service/pkg/handler"
	"github.com/Maksat-luci/REST-API-TODO-service/pkg/repository"
	"github.com/Maksat-luci/REST-API-TODO-service/pkg/service"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())
	}
	fmt.Println(viper.GetString("db.sslmode"))
	//создаём обьект структуры конфиг и передаём в функцию конструктора после чего получаем, ОБЬЕКТ ДЛЯ РАБОТЫ С ПОСТГРЕСОМ
	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: viper.GetString("dc.password"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  "disable",
	})
	if err != nil {
		log.Fatal("failed to initialize db:", err.Error())
	}

	//иницилизируем обьекты слоев и проводим dependency injection
	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)
	// создаём обьект хендлер с методом InitRoutes для прослушивания наших url адресов
	// создаём наш сервер с методом Run который конфигурирует наш сервер и начинает слушать URLы которык мы передали вторым аргументом
	srv := new(todo.Server)
	// начинаем слушать url адреса
	if err := srv.Run(viper.GetString("8000"), handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}

}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
