package main

import (
	"log"

	todo "github.com/Maksat-luci/REST-API-TODO-service"
	"github.com/Maksat-luci/REST-API-TODO-service/pkg/handler"
	"github.com/Maksat-luci/REST-API-TODO-service/pkg/repository"
	"github.com/Maksat-luci/REST-API-TODO-service/pkg/service"
	"github.com/spf13/viper"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())
	}

	//иницилизируем обьекты слоев и проводим dependency injection
	repos := repository.NewRepository()
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