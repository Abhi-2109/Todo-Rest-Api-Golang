package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type Todo struct {
	Id          int
	Name        string
	Description string
	Completed   bool
	UserId      int
	StartTime   time.Time
}
type User struct {
	Id   int
	Name string
}
type error struct {
	//Id        int
	Error string
}

var Todos []Todo
var Users []User

func homePage(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Welcome to the HomePage!")
	//fmt.Println("Endpoint Hit: homePage")
}
func errorHandler(writer http.ResponseWriter, errorName string) {
	var errorhere error
	errorhere.Error = errorName
	json.NewEncoder(writer).Encode(&errorhere)
}
func PostTodo(writer http.ResponseWriter, request *http.Request) {
	Name := request.FormValue("Name")
	Description := request.FormValue("Description")
	Completed, _ := strconv.ParseBool(request.FormValue("Completed"))
	UserId, _ := strconv.Atoi(request.FormValue("UserId"))
	flag := false
	newId := 0
	for _, i := range Users {
		if i.Id == UserId {
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
	var newTodo Todo
	newTodo.Id = newId + 1
	newTodo.Name = Name
	newTodo.Description = Description
	newTodo.Completed = Completed
	newTodo.UserId = UserId
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
	Name := request.FormValue("Name")
	newId := 1
	for _, i := range Users {
		if i.Id >= newId {
			newId = i.Id
		}
	}
	var newUser User
	newUser.Id = newId + 1
	newUser.Name = Name

	Users = append(Users, newUser)
	json.NewEncoder(writer).Encode(&newUser)

}
func GetAllUser(writer http.ResponseWriter, request *http.Request) {
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
			fmt.Println("error caught")
			errorHandler(writer, "Invalid URL")

		}
		return
	}

	if request.Method == "POST" {
		PostUser(writer, request)
	} else if request.Method == "GET" {
		GetAllUser(writer, request)
	} else {
		errorHandler(writer, "Invalid Request")
		// Invalid Request Function
	}
}
func GetSingleUser(writer http.ResponseWriter, request *http.Request, id int) {
	fmt.Println("get Single User")
	for _, singleUser := range Users {
		if singleUser.Id == id {
			json.NewEncoder(writer).Encode(&singleUser)

			return
		}
	}
	errorHandler(writer, "Id not Found. It may be deleted or it has been not created till")
}
func DeleteTodosUser(id int) {
	var leftTodos []Todo
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
	var allTodos []Todo

	for _, i := range Todos {
		if i.UserId == id {
			allTodos = append(allTodos, i)
		}
	}
	json.NewEncoder(writer).Encode(allTodos)
}
func usertodoHandler(writer http.ResponseWriter, request *http.Request) {
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
	newUser := User{1, "Abhishek"}

	Users = append(Users, newUser)
	newUser = User{2, "Anand"}
	Users = append(Users, newUser)

}
func main() {

	fillDummyData()

	handleRequest()
	// Specifying that it should listen at port :8081
	log.Fatal(http.ListenAndServe(":8081", nil))
}
