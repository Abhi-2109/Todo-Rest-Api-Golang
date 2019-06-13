package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetUsers(t *testing.T) {
	fillDummyData()
	req, err := http.NewRequest("GET", "/user/", nil)

	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(userHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	//expected := `{"Id":2,"Name":"Anand"}`

	expected := `[{"Id":1,"Name":"Abhishek"},{"Id":2,"Name":"Anand"}]`
	if rr.Body.String()[:len(rr.Body.String())-1] != expected {
		t.Errorf("handler returned unexpected body: got %v want %v , Type is %T, %T",
			rr.Body.String(), expected, rr.Body.String(), expected)
	}
}
func TestGetSingleUser(t *testing.T) {
	//fillDummyData()
	req, err := http.NewRequest("GET", "/user/1", nil)

	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(userHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	//expected := `{"Id":2,"Name":"Anand"}`

	expected := `{"Id":1,"Name":"Abhishek"}`
	if rr.Body.String()[:len(rr.Body.String())-1] != expected {
		t.Errorf("handler returned unexpected body: got %v want %v , Type is %T, %T",
			rr.Body.String(), expected, rr.Body.String(), expected)
	}
}
func TestUserIDByNotFound(t *testing.T) {
	//fillDummyData()
	req, err := http.NewRequest("GET", "/user/5", nil)

	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(userHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	//expected := `{"Id":2,"Name":"Anand"}`

	expected := `{"Errorname":"Id not Found. It may be deleted or it has been not created till"}`
	if rr.Body.String()[:len(rr.Body.String())-1] != expected {
		t.Errorf("handler returned unexpected body: got %v want %v , Type is %T, %T",
			rr.Body.String(), expected, rr.Body.String(), expected)
	}

}
func TestUserIDByInvalidURL(t *testing.T) {
	//fillDummyData()
	req, err := http.NewRequest("GET", "/user/5s", nil)

	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(userHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	//expected := `{"Id":2,"Name":"Anand"}`

	expected := `{"Errorname":"Invalid URL"}`
	if rr.Body.String()[:len(rr.Body.String())-1] != expected {
		t.Errorf("handler returned unexpected body: got %v want %v , Type is %T, %T",
			rr.Body.String(), expected, rr.Body.String(), expected)
	}
}
func TestPostUser(t *testing.T) {
	//fillDummyData()
	var jsonStr = []byte(`{"Name" : "Rohit"}`)
	//data := url.Values{}
	//data.Set("Name", "Rohit")
	//fmt.Println(data)

	req, err := http.NewRequest("POST", "/user/", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	//req.Header.Set("Content-Type", "application/x-www-form-urlencoded;charset=UTF-8")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(userHandler)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	expected := `{"Id":3,"Name":"Rohit"}`
	if rr.Body.String()[:len(rr.Body.String())-1] != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestPostTodo(t *testing.T) {
	var jsonStr = []byte(`{"Name" : "Todo","UserId" : 2}`)

	req, err := http.NewRequest("POST", "/todo/", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(todoHandler)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	expected := `{"Id":3,"Name":"Todo","Desciption":"","Completed":false,"UserId":2,"StartTime":"0001-01-01T00:00:00Z"}`
	if rr.Body.String()[:len(rr.Body.String())-1] != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
func TestGetTodoIDByNOtFound(t *testing.T) {
	req, err := http.NewRequest("GET", "/todo/5", nil)

	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(todoHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	//expected := `{"Id":2,"Name":"Anand"}`

	expected := `{"Errorname":"Id not Found. It may be deleted or it has been not created till"}`
	if rr.Body.String()[:len(rr.Body.String())-1] != expected {
		t.Errorf("handler returned unexpected body: got %v want %v , Type is %T, %T",
			rr.Body.String(), expected, rr.Body.String(), expected)
	}
}
func TestGetTodoIDByInvalidURL(t *testing.T) {
	req, err := http.NewRequest("GET", "/todo/5s", nil)

	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(todoHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	//expected := `{"Id":2,"Name":"Anand"}`

	expected := `{"Errorname":"Invalid URL"}`
	if rr.Body.String()[:len(rr.Body.String())-1] != expected {
		t.Errorf("handler returned unexpected body: got %v want %v , Type is %T, %T",
			rr.Body.String(), expected, rr.Body.String(), expected)
	}
}
func TestGetTodo(t *testing.T) {
	//fillDummyData()
	req, err := http.NewRequest("GET", "/todo/", nil)

	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(todoHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	expected := `[{"Id":1,"Name":"Anandhjvwdghjkjkj","Desciption":"","Completed":false,"UserId":1,"StartTime":"0001-01-01T00:00:00Z"},{"Id":2,"Name":"Man","Desciption":"","Completed":false,"UserId":2,"StartTime":"0001-01-01T00:00:00Z"},{"Id":3,"Name":"Todo","Desciption":"","Completed":false,"UserId":2,"StartTime":"0001-01-01T00:00:00Z"}]`
	if rr.Body.String()[:len(rr.Body.String())-1] != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
func TestGetSingleTodo(t *testing.T) {
	req, err := http.NewRequest("GET", "/todo/1", nil)

	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(todoHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	expected := `{"Id":1,"Name":"Anandhjvwdghjkjkj","Desciption":"","Completed":false,"UserId":1,"StartTime":"0001-01-01T00:00:00Z"}`
	if rr.Body.String()[:len(rr.Body.String())-1] != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestGetUserTodoByUserID(t *testing.T) {
	req, err := http.NewRequest("GET", "/usertodo/1", nil)

	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(usertodoHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	expected := `[{"Id":1,"Name":"Anandhjvwdghjkjkj","Desciption":"","Completed":false,"UserId":1,"StartTime":"0001-01-01T00:00:00Z"}]`
	if rr.Body.String()[:len(rr.Body.String())-1] != expected {
		t.Errorf("handler returned unexpected body: got %v want %v , Type is %T, %T",
			rr.Body.String(), expected, rr.Body.String(), expected)
	}

}

func TestGetUserTodoByNoUserID(t *testing.T) {
	req, err := http.NewRequest("GET", "/usertodo/", nil)

	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(usertodoHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	expected := `{"Errorname":"Specify Id in URL"}`
	if rr.Body.String()[:len(rr.Body.String())-1] != expected {
		t.Errorf("handler returned unexpected body: got %v want %v , Type is %T, %T",
			rr.Body.String(), expected, rr.Body.String(), expected)
	}
}

func TestGetUserTodoByInvalidUserID(t *testing.T) {
	req, err := http.NewRequest("GET", "/usertodo/7", nil)

	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(usertodoHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	expected := `{"Errorname":"Id is either deleted or not created"}`
	if rr.Body.String()[:len(rr.Body.String())-1] != expected {
		t.Errorf("handler returned unexpected body: got %v want %v , Type is %T, %T",
			rr.Body.String(), expected, rr.Body.String(), expected)
	}

}
func TestGetUserTodoByInvalidUrl(t *testing.T) {
	req, err := http.NewRequest("GET", "/usertodo/5s", nil)

	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(usertodoHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	expected := `{"Errorname":"Invalid URL"}`
	if rr.Body.String()[:len(rr.Body.String())-1] != expected {
		t.Errorf("handler returned unexpected body: got %v want %v , Type is %T, %T",
			rr.Body.String(), expected, rr.Body.String(), expected)
	}
}
func TestPostTodoByInvalidUserID(t *testing.T) {
	var jsonStr = []byte(`{"Name" : "Todo","UserId" : 8}`)

	req, err := http.NewRequest("POST", "/todo/", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(todoHandler)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	expected := `{"Errorname":"User Id is Invalid"}`
	if rr.Body.String()[:len(rr.Body.String())-1] != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestUserInvalidMethod(t *testing.T) {
	req, err := http.NewRequest("PUT", "/user/", nil)

	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(userHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	//expected := `{"Id":2,"Name":"Anand"}`

	expected := `{"Errorname":"Invalid Request"}`
	if rr.Body.String()[:len(rr.Body.String())-1] != expected {
		t.Errorf("handler returned unexpected body: got %v want %v , Type is %T, %T",
			rr.Body.String(), expected, rr.Body.String(), expected)
	}

}
func TestTodoInvalidMethod(t *testing.T) {
	req, err := http.NewRequest("PUT", "/todo/", nil)

	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(todoHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	//expected := `{"Id":2,"Name":"Anand"}`

	expected := `{"Errorname":"Invalid Request"}`
	if rr.Body.String()[:len(rr.Body.String())-1] != expected {
		t.Errorf("handler returned unexpected body: got %v want %v , Type is %T, %T",
			rr.Body.String(), expected, rr.Body.String(), expected)
	}

}
func TestSingleUserInvalidMethod(t *testing.T) {
	req, err := http.NewRequest("PUT", "/user/1", nil)

	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(userHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	//expected := `{"Id":2,"Name":"Anand"}`

	expected := `{"Errorname":"Invalid Request"}`
	if rr.Body.String()[:len(rr.Body.String())-1] != expected {
		t.Errorf("handler returned unexpected body: got %v want %v , Type is %T, %T",
			rr.Body.String(), expected, rr.Body.String(), expected)
	}

}
func TestSingleTodoInvalidMethod(t *testing.T) {
	req, err := http.NewRequest("PUT", "/todo/1", nil)

	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(todoHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	//expected := `{"Id":2,"Name":"Anand"}`

	expected := `{"Errorname":"Invalid Request"}`
	if rr.Body.String()[:len(rr.Body.String())-1] != expected {
		t.Errorf("handler returned unexpected body: got %v want %v , Type is %T, %T",
			rr.Body.String(), expected, rr.Body.String(), expected)
	}
}
func TestDeleteTodo(t *testing.T) {
	req, err := http.NewRequest("DELETE", "/todo/1", nil)

	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(todoHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	expected := `{"Id":1,"Name":"Anandhjvwdghjkjkj","Desciption":"","Completed":false,"UserId":1,"StartTime":"0001-01-01T00:00:00Z"}`
	if rr.Body.String()[:len(rr.Body.String())-1] != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
func TestDeleteUser(t *testing.T) {
	req, err := http.NewRequest("DELETE", "/user/1", nil)

	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(userHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	expected := `{"Id":1,"Name":"Abhishek"}`
	if rr.Body.String()[:len(rr.Body.String())-1] != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
func TestDeleteUserIDNotFOund(t *testing.T) {
	req, err := http.NewRequest("DELETE", "/user/8", nil)

	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(userHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	expected := `{"Errorname":"Id does not exist"}`
	if rr.Body.String()[:len(rr.Body.String())-1] != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
func TestDeleteUserInvalidURL(t *testing.T) {
	req, err := http.NewRequest("DELETE", "/user/8s", nil)

	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(userHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	expected := `{"Errorname":"Invalid URL"}`
	if rr.Body.String()[:len(rr.Body.String())-1] != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
func TestDeleteTodoInvalidURL(t *testing.T) {
	req, err := http.NewRequest("DELETE", "/todo/8s", nil)

	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(todoHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	expected := `{"Errorname":"Invalid URL"}`
	if rr.Body.String()[:len(rr.Body.String())-1] != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
func TestDeleteTodoIDNotFound(t *testing.T) {
	req, err := http.NewRequest("DELETE", "/todo/8", nil)

	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(todoHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	expected := `{"Errorname":"Id does not exist"}`
	if rr.Body.String()[:len(rr.Body.String())-1] != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestHomePage(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)

	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(homePage)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	expected := `Welcome to the HomePage!`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v , Type is %T, %T",
			rr.Body.String(), expected, rr.Body.String(), expected)
	}
}
func TestHandleRequest(t *testing.T) {
	req, err := http.NewRequest("POST", "/todo/1", nil)

	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	http.DefaultServeMux.ServeHTTP(rr, req)

	t.Errorf(rr.Body.String())
}
