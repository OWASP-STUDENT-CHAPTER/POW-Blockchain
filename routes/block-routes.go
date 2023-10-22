package routes

import (
	"net/http"

	"github.com/KulwinderSingh07/POW-Blockchain/controllers"
	"github.com/gorilla/mux"
)

func CreateMuxRoutes() http.Handler {
	muxRouter := mux.NewRouter()
	muxRouter.HandleFunc("/", controllers.HandleGetBlockchain).Methods("GET")
	muxRouter.HandleFunc("/", controllers.HandleWriteBlock).Methods("POST")
	return muxRouter
}
