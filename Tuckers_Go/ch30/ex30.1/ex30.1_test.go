package main

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestJsonHandler1(t *testing.T) {
	ass := assert.New(t)

	res := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/students", nil)

	mux := MakeWebHandler()
	mux.ServeHTTP(res, req)

	ass.Equal(http.StatusOK, res.Code)
	var list []Student
	err := json.NewDecoder(res.Body).Decode(&list)
	ass.Nil(err)
	ass.Equal(2, len(list))
	ass.Equal("aaa", list[0].Name)
	ass.Equal("bbb", list[1].Name)
}

func TestJsonHandler2(t *testing.T) {
	ass := assert.New(t)

	var student Student
	mux := MakeWebHandler()
	res := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/students/1", nil)

	mux.ServeHTTP(res, req)
	ass.Equal(http.StatusOK, res.Code)
	err := json.NewDecoder(res.Body).Decode(&student)
	ass.Nil(err)
	ass.Equal("aaa", student.Name)

	res = httptest.NewRecorder()
	req = httptest.NewRequest("GET", "/students/2", nil)
	mux.ServeHTTP(res, req)
	ass.Equal(http.StatusOK, res.Code)
	err = json.NewDecoder(res.Body).Decode(&student)
	ass.Nil(err)
	ass.Equal("bbb", student.Name)
}

func TestJsonHandler3(t *testing.T) {
	ass := assert.New(t)

	var student Student
	mux := MakeWebHandler()
	res := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/students", strings.NewReader(`{"Id":0,"Name":"ccc","Age":15,"Score": 78}`))

	mux.ServeHTTP(res, req)
	ass.Equal(http.StatusCreated, res.Code)

	res = httptest.NewRecorder()
	req = httptest.NewRequest("GET", "/students/3", nil)
	mux.ServeHTTP(res, req)
	ass.Equal(http.StatusOK, res.Code)
	err := json.NewDecoder(res.Body).Decode(&student)
	ass.Nil(err)
	ass.Equal("ccc", student.Name)
}

func TestJsonHandler4(t *testing.T) {
	ass := assert.New(t)

	mux := MakeWebHandler()
	res := httptest.NewRecorder()
	req := httptest.NewRequest("DELETE", "/students/1", nil)

	mux.ServeHTTP(res, req)
	ass.Equal(http.StatusOK, res.Code)

	res = httptest.NewRecorder()
	req = httptest.NewRequest("GET", "/students", nil)
	mux.ServeHTTP(res, req)

	ass.Equal(http.StatusOK, res.Code)
	var list Students
	err := json.NewDecoder(res.Body).Decode(&list)
	ass.Nil(err)
	ass.Equal(1, len(list))
	ass.Equal("bbb", list[0].Name)
}
