package main

import (
	"fmt"
	_ "github.com/restream/reindexer/v3/bindings/cproto"
	"github.com/spf13/viper"
	"github.com/valyala/fasthttp"
	"log"
	"resenje.org/logging"
	reindexer_db "service-for-example/internal/api_db/postgres_db"
	"service-for-example/internal/handler"
	"service-for-example/internal/service"
)

func main() {
	viper.SetConfigFile("config/config.yaml")
	if err := viper.ReadInConfig(); err != nil {
		logging.Info(fmt.Errorf("config: %v", err))
	}

	//установление связи с бд
	//проверка бд

	logging.Info("Connection to reindexer DB successful!")

	userApiDB := reindexer_db.NewUserApiDB()       //<-сюда прокинуть бд
	chatApiDB := reindexer_db.NewChatApiDB()       //<-сюда прокинуть бд
	messageApiDB := reindexer_db.NewMessageApiDB() //<-сюда прокинуть бд
	activeApiDB := reindexer_db.NewActiveApiDB()   //<-сюда прокинуть бд
	generalApiDB := reindexer_db.NewGeneralApiDB() //<-сюда прокинуть бд

	s := service.NewService(userApiDB, chatApiDB, messageApiDB, activeApiDB, generalApiDB)

	h := handler.NewHandler(s)

	server := fasthttp.Server{
		Handler: h.InitRoutes,
	}

	if err := server.ListenAndServe(viper.GetString("http.port")); err != nil {
		log.Fatalf("start server: %v", err)
	}
}
