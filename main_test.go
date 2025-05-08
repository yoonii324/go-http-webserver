package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

// func TestIndexHandler(t *testing.T) { // path / test
// 	assert := assert.New(t)

// 	res := httptest.NewRecorder()
// 	req := httptest.NewRequest("GET", "/", nil)

// 	mux := MakeWebHandler()
// 	mux.ServeHTTP(res, req)

// 	assert.Equal(http.StatusOK, res.Code) // check status code
// 	data, _ := io.ReadAll(res.Body)       // by reading the data
// 	assert.Equal("Hello World", string(data))
// }

// func TestBarHandler(t *testing.T) {
// 	assert := assert.New(t)

// 	res := httptest.NewRecorder()
// 	req := httptest.NewRequest("GET", "/bar", nil) // path /bar test

// 	mux := MakeWebHandler()
// 	mux.ServeHTTP(res, req)

// 	assert.Equal(http.StatusOK, res.Code)
// 	data, _ := io.ReadAll(res.Body)
// 	assert.Equal("Hello Bar", string(data))
// }

func TestJsonHandler(t *testing.T) {
	assert := assert.New(t)

	res := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/student", nil) // path /student test

	mux := MakeWebHandler()
	mux.ServeHTTP(res, req)

	assert.Equal(http.StatusOK, res.Code) // check status code
	student := new(Student)
	err := json.NewDecoder(res.Body).Decode(student) // translate the result
	assert.Nil(err)                                  // check the result
	assert.Equal("aaa", student.Name)
	assert.Equal(16, student.Age)
	assert.Equal(87, student.Score)
}
