package pkg1_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/zackarysantana/repo1/pkg1"
)

func TestHelloWorld(t *testing.T) {
	// Create a fake HTTP request
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()

	// Call the HelloWorld handler with the fake request and response recorder
	pkg1.HelloWorld(rr, req)

	// Check the status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body
	expected := "HELLOWORLD"
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}