package main

import (
	"log"
	"net/http"
	"os"
	"staffplan-reports/controller"
	"staffplan-reports/dao"
	"staffplan-reports/service"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	db := dao.Connection()

	router := mux.NewRouter()
	merchandiserRepo := dao.NewMerchandiserRepo(db)
	merchandiserUseCase := service.NewMerchandiserUseCase(merchandiserRepo)
	merchansdiserRouter := router.PathPrefix("/report/merchandiser").Subrouter()
	merchandiserController := controller.MerchandiserController{
		UseCase: merchandiserUseCase,
	}

	merchansdiserRouter.HandleFunc("/route/stats", merchandiserController.RouteStats)

	var address string

	if port, ok := os.LookupEnv("PORT"); ok && port != "" {
		address = ":" + port

	} else {
		address += ":4000"
	}

	server := &http.Server{
		Addr:         address,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		Handler:      router,
	}

	println("server starting on", address)

	if err := server.ListenAndServe(); err != nil {
		db.Close()
		log.Fatal(err.Error())
	}
}
