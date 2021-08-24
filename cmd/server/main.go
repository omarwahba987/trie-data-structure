package main

import (
	"github.com/gorilla/mux"
	"va_test_d/api"
	"va_test_d/httpHelper"
)

func main() {
	api.InitTrie()
	Router := mux.NewRouter()
	httpHelper := httpHelper.HttpHandlerRepo{}
	defer httpHelper.OpenServer(Router, "8081")
	Router.HandleFunc("/insert", httpHelper.HandleHttpRequest(api.InsertPerson)).Methods("POST")
	Router.HandleFunc("/visualize", httpHelper.HandleHttpRequest(api.Visualize)).Methods("GET")

}
