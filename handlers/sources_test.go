package handlers

import (
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type mockGoodRepository struct {
}
type mockBadRepository struct {
}

func (mockGoodRepository mockGoodRepository) CreateRoute(source Source) error {
	return nil
}
func (mockGoodRepository mockGoodRepository) GetSource(source Source) (Source, error) {
	return Source{Routes: append(make([]Route, 0), Route{URL: "https://www.google.com"})}, nil
}
func (mockGoodRepository mockGoodRepository) GetAllSources() ([]Source, error) {
	return append(make([]Source, 1), Source{Name: "MOCKNAME"}), nil
}

func (mockBadRepository mockBadRepository) CreateRoute(source Source) error {
	return errors.New("THIS IS AN ERROR")
}
func (mockBadRepository mockBadRepository) GetSource(source Source) (Source, error) {
	return Source{}, errors.New("THIS IS AN ERROR")
}
func (mockBadRepository mockBadRepository) GetAllSources() ([]Source, error) {
	return make([]Source, 0), errors.New("THIS IS AN ERROR")
}
func TestCreateRouteStatusCreated(t *testing.T) {
	r := getRouter()

	handler := HTTPSourceHandler{Dynamo: mockGoodRepository{}}

	r.POST("/sources/name/routes", handler.CreateRoute)

	req, _ := http.NewRequest("POST", "/sources/name/routes",
		strings.NewReader(`{"route": {"url": "String"},"withSourceCreation": "boolean"}`))

	testHTTPResponse(t, r, req, func(w *httptest.ResponseRecorder) bool {

		_, err := ioutil.ReadAll(w.Body)
		if err != nil {
			t.Log("Error parsing body")
			t.Fail()
		}
		if w.Code != http.StatusCreated {
			t.Log("incorrect status, expected ", http.StatusCreated, " got ", w.Code)
			t.Fail()
		}
		return w.Code == http.StatusCreated
	})

}
func TestCreateRouteInvalidRequestUnprocessableEntity(t *testing.T) {
	r := getRouter()

	handler := HTTPSourceHandler{Dynamo: mockGoodRepository{}}

	r.POST("/sources/name/routes", handler.CreateRoute)

	req, _ := http.NewRequest("POST", "/sources/name/routes", nil)

	testHTTPResponse(t, r, req, func(w *httptest.ResponseRecorder) bool {

		_, err := ioutil.ReadAll(w.Body)
		if err != nil {
			t.Log("Error parsing body")
			t.Fail()
		}
		if w.Code != http.StatusUnprocessableEntity {
			t.Log("incorrect status, expected ", http.StatusUnprocessableEntity, " got ", w.Code)
			t.Fail()
		}
		return w.Code == http.StatusUnprocessableEntity
	})

}
func TestCreateRouteInvalidRequestDatabaseFailsReturnsServiceUnavailable(t *testing.T) {
	r := getRouter()

	handler := HTTPSourceHandler{Dynamo: mockBadRepository{}}

	r.POST("/sources/name/routes", handler.CreateRoute)

	req, _ := http.NewRequest("POST", "/sources/name/routes",
		strings.NewReader(`{"route": {"url": "String"},"withSourceCreation": "boolean"}`))

	testHTTPResponse(t, r, req, func(w *httptest.ResponseRecorder) bool {

		_, err := ioutil.ReadAll(w.Body)
		if err != nil {
			t.Log("Error parsing body")
			t.Fail()
		}
		if w.Code != http.StatusServiceUnavailable {
			t.Log("incorrect status, expected ", http.StatusServiceUnavailable, " got ", w.Code)
			t.Fail()
		}
		return w.Code == http.StatusServiceUnavailable
	})

}
func TestGetAllHTTPStatusOK(t *testing.T) {
	r := getRouter()
	handler := HTTPSourceHandler{Repository: mockGoodRepository{}}

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
func TestGetAllDBErrorReturnsHTTPStatusServiceUnavailable(t *testing.T) {
	r := getRouter()
	handler := HTTPSourceHandler{Repository: mockBadRepository{}}

	r.GET("/sources", handler.GetAllSources)

	req, _ := http.NewRequest("GET", "/sources", nil)

	testHTTPResponse(t, r, req, func(w *httptest.ResponseRecorder) bool {
		statusUnavailable := w.Code == http.StatusServiceUnavailable

		_, err := ioutil.ReadAll(w.Body)
		if err != nil {
			t.Log("Error parsing body")
			t.Fail()
		}
		return statusUnavailable
	})

}
func TestGetAll2SourcesAreReturned(t *testing.T) {
	r := getRouter()
	handler := HTTPSourceHandler{Repository: mockGoodRepository{}}
	r.GET("/sources", handler.GetAllSources)

	req, _ := http.NewRequest("GET", "/sources", nil)

	testHTTPResponse(t, r, req, func(w *httptest.ResponseRecorder) bool {
		body, err := ioutil.ReadAll(w.Body)
		expected := `{"sources":[{"name":"","routes":null},{"name":"MOCKNAME","routes":null}]}`
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

func TestGetRoutesOkHttpStatusOK(t *testing.T) {
	r := getRouter()
	handler := HTTPSourceHandler{Dynamo: mockGoodRepository{}}
	r.GET("/sources/somename/routes", handler.GetRoutes)

	req, _ := http.NewRequest("GET", "/sources/somename/routes", nil)

	testHTTPResponse(t, r, req, func(w *httptest.ResponseRecorder) bool {
		_, err := ioutil.ReadAll(w.Body)
		expected := http.StatusOK
		actual := w.Code
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

func TestGetRoutesOkReturnsSameRoutes(t *testing.T) {
	r := getRouter()
	handler := HTTPSourceHandler{Dynamo: mockGoodRepository{}}
	r.GET("/sources/somename/routes", handler.GetRoutes)

	req, _ := http.NewRequest("GET", "/sources/somename/routes", nil)

	testHTTPResponse(t, r, req, func(w *httptest.ResponseRecorder) bool {
		body, err := ioutil.ReadAll(w.Body)
		expected := `{"routes":[{"url":"https://www.google.com"}]}`
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
