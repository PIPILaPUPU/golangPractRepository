package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCafeWhenOk(t *testing.T) {
	handler := http.HandlerFunc(mainHandle)

	requests := []string{
		"/cafe?count=2&city=moscow",
		"/cafe?city=tula",
		"/cafe?city=moscow&search=ложка",
	}
	for _, v := range requests {
		response := httptest.NewRecorder()
		req := httptest.NewRequest("GET", v, nil)

		handler.ServeHTTP(response, req)

		assert.Equal(t, http.StatusOK, response.Code)
		// пока сравнивать не будем, а просто выведем ответы
		// удалите потом этот вывод
		fmt.Println(response.Body.String())
	}
}

func TestCafeWhenErr(t *testing.T) {
	handler := http.HandlerFunc(mainHandle)

	req := []struct {
		Address string
		Status  int
		message string
	}{
		{"/cafe", http.StatusBadRequest, "unknown city"},
		{"/cafe?city=omsk", http.StatusBadRequest, "unknown city"},
		{"/cafe?city=tula&count=na", http.StatusBadRequest, "incorrect count"},
	}

	for _, item := range req {
		response := httptest.NewRecorder()
		req := httptest.NewRequest("GET", item.Address, nil)

		handler.ServeHTTP(response, req)

		assert.Equal(t, item.Status, response.Code)
		assert.Equal(t, item.message, strings.TrimSpace(response.Body.String()))
		fmt.Println(response.Body.String())
	}
}
