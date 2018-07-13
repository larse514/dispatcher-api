package handlers

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetAllRoutesHTTPStatusOK(t *testing.T) {
	r := getRouter()
	handler := HTTPRouteHandler{}
	r.GET("/routes", handler.GetAllRoutes)

	req, _ := http.NewRequest("GET", "/routes", nil)

	testHTTPResponse(t, r, req, func(w *httptest.ResponseRecorder) bool {
		statusOK := w.Code == http.StatusOK

		_, err := ioutil.ReadAll(w.Body)
		if err != nil {
			t.Log("Error parsing body")
			t.Fail()
		}
		return statusOK
	})

}

func TestGetAllRoutes2AreReturned(t *testing.T) {
	r := getRouter()
	handler := HTTPRouteHandler{}
	r.GET("/routes", handler.GetAllRoutes)

	req, _ := http.NewRequest("GET", "/routes", nil)

	testHTTPResponse(t, r, req, func(w *httptest.ResponseRecorder) bool {
		body, err := ioutil.ReadAll(w.Body)
		expected := `{"routes":[{"url":"https://www.google.com"},{"url":"https://www.msn.com"}]}`
		actual := string(body)
		sourcesOk := err == nil && actual == expected

		if err != nil {
			t.Log("Error parsing body")
			t.Fail()
		}
		if actual != expected {
			t.Log("expected ", expected, " got ", actual)
		}
		return sourcesOk
	})

}
