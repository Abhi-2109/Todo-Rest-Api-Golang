package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/abhi2109/todo_API/data"
)

var Todos data.TodoArray
var Users data.UserArray

func homePage(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Welcome to the HomePage!")
	//fmt.Println("Endpoint Hit: homePage")
}
func errorHandler(writer http.ResponseWriter, errorName string) {
	var errorhere data.Error
	errorhere.Errorname = errorName
	json.NewEncoder(writer).Encode(errorhere)
}
func PostTodo(writer http.ResponseWriter, request *http.Request) {
	b, err := ioutil.ReadAll(request.Body)
	if err != nil {
		errorHandler(writer, "there")
		return
	}
	defer request.Body.Close()
	var newTodo data.Todo
	erro := newTodo.UnmarshalJSON(b)
	if erro != nil {
		errorHandler(writer, "Error in Processing the Data")
		return
	}
	flag := false
	newId := 0
	for _, i := range Users {
		if i.Id == newTodo.UserId {
			flag = true
			break
		}
	}
	if flag == false {
		errorHandler(writer, "User Id is Invalid")
		return
	}
	for _, i := range Todos {
		if i.Id >= newId {
			newId = i.Id
		}
	}
	newTodo.Id = newId + 1
	newTodo.StartTime = time.Now()
	Todos = append(Todos, newTodo)
	json.NewEncoder(writer).Encode(&newTodo)

}
func GetAllTodo(writer http.ResponseWriter, request *http.Request) {
	json.NewEncoder(writer).Encode(Todos)
}
func todoHandler(writer http.ResponseWriter, request *http.Request) {

	urlPath := request.URL.Path
	path := strings.Split(urlPath, "/")
	if path[len(path)-1] != "" {
		id, err := strconv.Atoi(path[len(path)-1])
		if err == nil {
			singleTodoHandler(writer, request, id)

		} else {
			errorHandler(writer, "Invalid URL")
		}
		return
	}

	if request.Method == "POST" {
		PostTodo(writer, request)

	} else if request.Method == "GET" {
		GetAllTodo(writer, request)
	} else {
		errorHandler(writer, "Invalid Request")
	}

}
func GetSingleTodo(writer http.ResponseWriter, request *http.Request, id int) {
	fmt.Println("Get Single Todo")
	errorHandler(writer, "Single Todo")
	for _, singleTodo := range Todos {
		if singleTodo.Id == id {
			json.NewEncoder(writer).Encode(singleTodo)
			return
		}
	}
	fmt.Println("Error in Deccoding")
	errorHandler(writer, "Id not Found. It may be deleted or it has been not created till")
}

func DeleteTodo(writer http.ResponseWriter, request *http.Request, id int) {
	for ids, i := range Todos {
		if i.Id == id {
			Todos = append(Todos[:ids], Todos[ids+1:]...)
			json.NewEncoder(writer).Encode(i)
			return
		}
	}
	errorHandler(writer, "Id does not exist")
}

func singleTodoHandler(writer http.ResponseWriter, request *http.Request, id int) {
	if request.Method == "GET" {
		GetSingleTodo(writer, request, id)
	} else if request.Method == "DELETE" {
		DeleteTodo(writer, request, id)
	} else {
		errorHandler(writer, "Invalid Request")
		// Invalid Request Function
	}

}

func PostUser(writer http.ResponseWriter, request *http.Request) {

	b, err := ioutil.ReadAll(request.Body)
	if err != nil {
		errorHandler(writer, "there")
		return
	}
	defer request.Body.Close()
	var newUser data.User
	error := newUser.UnmarshalJSON(b)
	if error != nil {
		errorHandler(writer, "Error in Processing the Data")
	}
	//json.Unmarshal(b, &newUser)
	newId := 1
	for _, i := range Users {
		if i.Id >= newId {
			newId = i.Id
		}
	}
	newUser.Id = newId + 1

	Users = append(Users, newUser)
	json.NewEncoder(writer).Encode(&newUser)

}
func GetAllUser(writer http.ResponseWriter, request *http.Request) {
	writer.WriteHeader(http.StatusOK)
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(Users)

}
func userHandler(writer http.ResponseWriter, request *http.Request) {

	urlPath := request.URL.Path
	path := strings.Split(urlPath, "/")
	if path[len(path)-1] != "" {
		id, err := strconv.Atoi(path[len(path)-1])
		if err == nil {
			singleUserHandler(writer, request, id)

		} else {
			errorHandler(writer, "Invalid URL")

		}
		return
	} else {

		if request.Method == "POST" {
			PostUser(writer, request)
		} else if request.Method == "GET" {
			//errorHandler(writer, "This is shit")
			GetAllUser(writer, request)
		} else {
			errorHandler(writer, "Invalid Request")
			// Invalid Request Function
		}
	}
}
func GetSingleUser(writer http.ResponseWriter, request *http.Request, id int) {
	for _, singleUser := range Users {
		if singleUser.Id == id {
			json.NewEncoder(writer).Encode(&singleUser)

			return
		}
	}
	errorHandler(writer, "Id not Found. It may be deleted or it has been not created till")
}
func DeleteTodosUser(id int) {
	var leftTodos data.TodoArray
	for _, i := range Todos {
		if i.UserId != id {
			leftTodos = append(leftTodos, i)
		}
	}
	Todos = leftTodos
}
func DeleteUser(writer http.ResponseWriter, request *http.Request, id int) {
	for ids, i := range Users {
		if i.Id == id {
			DeleteTodosUser(id)
			Users = append(Users[:ids], Users[ids+1:]...)
			json.NewEncoder(writer).Encode(i)
			return
		}
	}
	errorHandler(writer, "Id does not exist")
}
func singleUserHandler(writer http.ResponseWriter, request *http.Request, id int) {
	if request.Method == "GET" {
		GetSingleUser(writer, request, id)

	} else if request.Method == "DELETE" {
		DeleteUser(writer, request, id)
	} else {
		errorHandler(writer, "Invalid Request")
		// Invalid Request Function
	}
}

func getUserAllTodo(writer http.ResponseWriter, request *http.Request, id int) {
	flag := false
	for _, i := range Users {
		if i.Id == id {
			flag = true
			break
		}
	}
	if flag == false {
		errorHandler(writer, "Id is either deleted or not created")
		return
	}
	var allTodos data.TodoArray

	for _, i := range Todos {
		if i.UserId == id {
			allTodos = append(allTodos, i)
		}
	}
	json.NewEncoder(writer).Encode(allTodos)
}
func usertodoHandler(writer http.ResponseWriter, request *http.Request) {

	fmt.Println("User Todo handler")
	urlPath := request.URL.Path
	path := strings.Split(urlPath, "/")
	if path[len(path)-1] != "" {
		id, err := strconv.Atoi(path[len(path)-1])
		if err == nil {
			getUserAllTodo(writer, request, id)
		} else {
			fmt.Println("error caught")
			errorHandler(writer, "Invalid URL")

		}
		return
	} else {
		errorHandler(writer, "Specify Id in URL")
	}
}

func handleRequest() {
	/* This function is to handle all the web request  */

	http.HandleFunc("/", homePage)
	http.HandleFunc("/todo/", todoHandler)
	http.HandleFunc("/user/", userHandler)
	http.HandleFunc("/usertodo/", usertodoHandler)

}
func fillDummyData() {
	newUser := data.User{1, "Abhishek"}

	Users = append(Users, newUser)
	newUser = data.User{2, "Anand"}
	Users = append(Users, newUser)

}
func main() {

	fillDummyData()

	handleRequest()
	// Specifying that it should listen at port :8081
	log.Fatal(http.ListenAndServe(":8081", nil))
}
