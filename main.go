package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/KulwinderSingh07/POW-Blockchain/controllers"
	"github.com/KulwinderSingh07/POW-Blockchain/model"
	"github.com/KulwinderSingh07/POW-Blockchain/routes"
	"github.com/joho/godotenv"
)

var NewBlock model.Block

var Blockchain []model.Block

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	controllers.Blockinitalizer()
	log.Fatal(run())
}

func run() error {
	mux := routes.CreateMuxRoutes()
	httpPort := os.Getenv("PORT")
	log.Println("Http server Listening on port :", httpPort)
	s := &http.Server{
		Addr:           ":" + httpPort,
		Handler:        mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	if err := s.ListenAndServe(); err != nil {
		return err
	}
	return nil
}
