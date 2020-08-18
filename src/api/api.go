package api

import(
	"encoding/json"
	"net/http"
	"strings"

	"github.com/francobottoni/go-rest-api/src/models"
	"github.com/francobottoni/go-rest-api/src/helpers"
	"github.com/gorilla/mux"
)

type Data struct{
	Success bool `json: "success"`
	Data []models.Todo `json: "data"`
	Errors []string `json: "errors"`
}

func CreateTodo(w http.ResponseWriter, req *http.Request){
	bodyTodo, success := helpers.DecodeBody(req)
	if success != true {
		http.Error(w, "could not decode body", http.StatusBadRequest)
		return
	}

	var data Data = Data{Errors: make([]string,0)}
	bodyTodo.Description = strings.TrimSpace(bodyTodo.Description)
	if !helpers.IsValidDescription(bodyTodo.Description){
		data.Success = false
		data.Errors = append(data.Errors, "invalid description")

		json, _ := json.Marshal(data)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(json)
		return 
	}

	todo, success := models.Insert(bodyTodo.Description)
	if success != true {
		data.Errors = append(data.Errors, "could not create todo")
	}


	data.Success = success
	data.Data = append(data.Data, todo)

	json, _ := json.Marshal(data)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(json)

}

func GetTodo(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id := vars["id"]

	var data Data

	var todo models.Todo
	var success bool
	todo, success = models.Get(id)
	if success != true {
		data.Success = false
		data.Errors = append(data.Errors, "todo not found")

		json, _ := json.Marshal(data)
		
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(json)
		return
	}

	data.Success = true
	data.Data = append(data.Data, todo)

	json, _ := json.Marshal(data)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(json) 
}