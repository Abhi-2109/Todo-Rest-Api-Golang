package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
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
	fillDummyData()
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
	fillDummyData()
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

	expected := `{"Error":"Id not Found. It may be deleted or it has been not created till"}`
	if rr.Body.String()[:len(rr.Body.String())-1] != expected {
		t.Errorf("handler returned unexpected body: got %v want %v , Type is %T, %T",
			rr.Body.String(), expected, rr.Body.String(), expected)
	}

}
func TestUserIDByInvalidURL(t *testing.T) {
	fillDummyData()
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

	expected := `{"Error":"Invalid URL"}`
	if rr.Body.String()[:len(rr.Body.String())-1] != expected {
		t.Errorf("handler returned unexpected body: got %v want %v , Type is %T, %T",
			rr.Body.String(), expected, rr.Body.String(), expected)
	}
}
func TestPostUser(t *testing.T) {
	fillDummyData()
	//var jsonStr = []byte(`{"Name" : "Rohit"}`)
	data := url.Values{}
	data.Set("Name", "Rohit")
	fmt.Println(data)

	req, err := http.NewRequest("POST", "/user/", strings.NewReader(data.Encode()))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded;charset=UTF-8")
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
