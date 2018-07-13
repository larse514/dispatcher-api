package handlers

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetAllHTTPStatusOK(t *testing.T) {
	r := getRouter()
	r.GET("/sources", GetAllSources)

	req, _ := http.NewRequest("GET", "/sources", nil)

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

func TestGetAll2SourcesAreReturned(t *testing.T) {
	r := getRouter()
	r.GET("/sources", GetAllSources)

	req, _ := http.NewRequest("GET", "/sources", nil)

	testHTTPResponse(t, r, req, func(w *httptest.ResponseRecorder) bool {
		body, err := ioutil.ReadAll(w.Body)
		expected := `{"sources":[{"name":"Service1","route":{"url":"https://www.google.com"}},{"name":"Service2","route":{"url":"https://www.google.com"}}]}`
		actual := string(body)
		sourcesOk := err == nil && actual == expected

		if err != nil {
			t.Log("Error parsing body")
			t.Fail()
		}
		return sourcesOk
	})

}
