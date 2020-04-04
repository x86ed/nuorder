package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHubHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/hub", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(hub)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	if len(rr.Body.String()) < 10 { //shamefully lazy test here
		t.Errorf("handler returned unexpected body: got %v",
			rr.Body.String())
	}
}
