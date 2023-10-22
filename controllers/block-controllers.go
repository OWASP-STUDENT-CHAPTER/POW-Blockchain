package controllers

import (
	"encoding/json"
	"io"
	"net/http"
)

func HandleGetBlockchain(res http.ResponseWriter, req *http.Request) {
	data, err := json.Marshal("hello")
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	io.WriteString(res, string(data))
}

func HandleWriteBlock(res http.ResponseWriter, req *http.Request) {
	data, err := json.Marshal("hello ji")
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	io.WriteString(res, string(data))
}
