package handlers

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

type mockGoodRepository struct {
}

func (mockGoodRepository mockGoodRepository) CreateRoute(source Source) error {
	return nil
}

func TestCreateRouteStatusCreated(t *testing.T) {
	// r := getRouter()
	// r.POST("/sources/name/routes", CreateRoute)

	// req, _ := http.NewRequest("POST", "/sources/name/routes", nil)

	// testHTTPResponse(t, r, req, func(w *httptest.ResponseRecorder) bool {

	// 	_, err := ioutil.ReadAll(w.Body)
	// 	if err != nil {
	// 		t.Log("Error parsing body")
	// 		t.Fail()
	// 	}
	// 	if w.Code != http.StatusCreated {
	// 		t.Log("incorrect status, expected ", http.StatusCreated, " got ", w.Code)
	// 		t.Fail()
	// 	}
	// 	return w.Code == http.StatusCreated
	// })

}
func TestGetAllHTTPStatusOK(t *testing.T) {
	r := getRouter()
	handler := HTTPSourceHandler{}

	r.GET("/sources", handler.GetAllSources)

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
	handler := HTTPSourceHandler{}
	r.GET("/sources", handler.GetAllSources)

	req, _ := http.NewRequest("GET", "/sources", nil)

	testHTTPResponse(t, r, req, func(w *httptest.ResponseRecorder) bool {
		body, err := ioutil.ReadAll(w.Body)
		expected := `{"sources":[{"name":"Service1","routes":[{"url":"https://www.google.com"}]},{"name":"Service2","routes":[{"url":"https://www.google.com"}]}]}`
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
