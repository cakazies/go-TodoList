package controllers

import (
	"fmt"
	"net/http"

	util "github.com/local/TaskListGo/util"
)

var Hello = func(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello World from controller")
	response := util.MetaMsg(true, "Hello World")
	util.Respond(w, response)
}

var ListToDos = func(w http.ResponseWriter, r *http.Request) {
	fmt.Println("List Todo")
	response := util.MetaMsg(true, "List Todo")
	util.Respond(w, response)
}

var PertoDocs = func(w http.ResponseWriter, r *http.Request) {
	fmt.Println("get Pertodo")
	response := util.MetaMsg(true, "get Pertodo")
	util.Respond(w, response)
}

var CreateToDocs = func(w http.ResponseWriter, r *http.Request) {
	fmt.Println("CreateToDocs")
	response := util.MetaMsg(true, "CreateToDocs")
	util.Respond(w, response)
}

var ActionToDocs = func(w http.ResponseWriter, r *http.Request) {
	fmt.Println("ActionToDocs")
	response := util.MetaMsg(true, "ActionToDocs")
	util.Respond(w, response)
}

var EditToDocs = func(w http.ResponseWriter, r *http.Request) {
	fmt.Println("EditToDocs")
	response := util.MetaMsg(true, "EditToDocs")
	util.Respond(w, response)
}

var DeleteToDocs = func(w http.ResponseWriter, r *http.Request) {
	fmt.Println("DeleteToDocs")
	response := util.MetaMsg(true, "DeleteToDocs")
	util.Respond(w, response)
}

var UserRegister = func(w http.ResponseWriter, r *http.Request) {
	fmt.Println("UserRegister")
	response := util.MetaMsg(true, "UserRegister")
	util.Respond(w, response)
}

var UserLogin = func(w http.ResponseWriter, r *http.Request) {
	fmt.Println("UserLogin")
	response := util.MetaMsg(true, "UserLogin")
	util.Respond(w, response)
}

//TODO
//1. Add all function required that is called from router
