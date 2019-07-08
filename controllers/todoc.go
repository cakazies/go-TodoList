package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/local/TaskListGo/models"
	util "github.com/local/TaskListGo/util"
)

var Hello = func(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello World from controller")
	response := util.MetaMsg(true, "Hello World")
	util.Respond(w, response)
}

var ListToDos = func(w http.ResponseWriter, r *http.Request) {
	data := models.GetToDo()
	response := util.MetaMsg(true, "Success")
	response["data"] = data
	util.Respond(w, response)
}

var PertoDocs = func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	toDoId, err := strconv.Atoi(params["toDold"])
	if err != nil {
		util.Respond(w, util.MetaMsg(false, "Params is invalid"))
		return
	}

	// uint for positif absolute
	data := models.SingleToDo(uint(toDoId))

	response := util.MetaMsg(true, "Success")
	response["data"] = data
	util.Respond(w, response)
}

var CreateToDocs = func(w http.ResponseWriter, r *http.Request) {
	todo := &models.Todo{}
	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&todo)
	if err != nil {
		panic(err)
	}
	response := todo.Create()
	util.Respond(w, response)
}

var ActionToDocs = func(w http.ResponseWriter, r *http.Request) {
	todo := &models.Todo{}
	params := mux.Vars(r)

	toDoId, err := strconv.Atoi(params["toDold"])
	if err != nil {
		util.Respond(w, util.MetaMsg(false, "Params is invalid"))
		return
	}
	todo.ID = uint(toDoId)
	response := todo.ActionToDo()
	util.Respond(w, response)
}

var EditToDocs = func(w http.ResponseWriter, r *http.Request) {
	todo := &models.Todo{}
	params := mux.Vars(r)
	toDoId, err := strconv.Atoi(params["toDold"])
	log.Println("params id : ", toDoId)
	if err != nil {
		util.MetaMsg(false, "Param is invalid")
		return
	}
	todo.ID = uint(toDoId)
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&todo)
	if err != nil {
		panic(err)
	}
	response := todo.EditToDo()
	util.Respond(w, response)
}

var DeleteToDocs = func(w http.ResponseWriter, r *http.Request) {
	todo := &models.Todo{}
	params := mux.Vars(r)
	toDoId, err := strconv.Atoi(params["toDold"])
	if err != nil {
		util.MetaMsg(false, "Param is invalid")
		return
	}
	todo.ID = uint(toDoId)
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&todo)
	if err != nil {
		panic(err)
	}
	response := todo.DeleteToDo()
	util.Respond(w, response)
}

var Register = func(w http.ResponseWriter, r *http.Request) {
	account := &models.Account{}

	err := json.NewDecoder(r.Body).Decode(account)
	if err != nil {
		fmt.Println(err)
		util.Respond(w, util.MetaMsg(false, "Invalid request"))
		return
	}

	response := account.CreateAccount()
	util.Respond(w, response)
}

var Login = func(w http.ResponseWriter, r *http.Request) {
	account := &models.Account{}

	err := json.NewDecoder(r.Body).Decode(account)
	if err != nil {
		fmt.Println(err)
		util.Respond(w, util.MetaMsg(false, "Invalid request"))
		return
	}

	response := account.Login()
	util.Respond(w, response)
}
