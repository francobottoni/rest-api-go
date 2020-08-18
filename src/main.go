package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/francobottoni/go-rest-api/src/api"
)

func main(){
	var port string = "8080"

	router := mux.NewRouter()
	apiRouter := router.PathPrefix("/api/").Subrouter()
	apiRouter.HandleFunc("/todos",api.CreateTodo).Methods("POST")
	apiRouter.HandleFunc("/todos/{id}",api.GetTodo).Methods("GET")

	fmt.Printf("Server running at port %s", port)
	http.ListenAndServe(":" +port, router)
}