package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHelloFunc(t *testing.T) {
	expected := "hello world"
	request, _ := http.NewRequest(http.MethodGet, "/hello", nil)
	w := httptest.NewRecorder()

	hello(w, request)

	response := w.Result()
	body, _ := ioutil.ReadAll(response.Body)
	actualResult := string(body)

	if actualResult != expected {
		t.Errorf("expected %s but got it %s", expected, actualResult)
	}
}
