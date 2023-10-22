package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/KulwinderSingh07/POW-Blockchain/model"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

var NewBlock model.Block

var Blockchain []model.Block

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	log.Fatal(run())
}

func run() error {
	mux := makeMuxRouter()
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

func testing(res http.ResponseWriter, req *http.Request) {
	data, err := json.Marshal("hello")
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	io.WriteString(res, string(data))
}

func makeMuxRouter() http.Handler {
	muxRouter := mux.NewRouter()
	muxRouter.HandleFunc("/", testing).Methods("GET")
	muxRouter.HandleFunc("/", testing).Methods("POST")
	return muxRouter
}
