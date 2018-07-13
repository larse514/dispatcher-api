package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// SourceHandler is an interface to handle Source HTTP requests
type SourceHandler interface {
	GetAllSources(c *gin.Context)
	CreateRoute(c *gin.Context)
}

//Source is the struct representation of a source object
//it contains a name and a route object
type Source struct {
	Name   string  `json:"name"`
	Routes []Route `json:"routes"`
}

//SourceDTO is a struct representing a POST sources request
type SourceDTO struct {
	Routes             []Route `json:"routes"`
	WithSourceCreation bool    `json:"withSourceCreation"`
}

// HTTPSourceHandler struct implementation of SourceHandler for HTTP requests
type HTTPSourceHandler struct {
	Repository SourceRepository
}

//SourceRepository is an interface used by the sources.go handler to interact with storage
type SourceRepository interface {
	CreateRoute(source Source) error
	GetRoutes(source Source) (Source, error)
}

//GetAllSources is a method to return all source object
func (handler HTTPSourceHandler) GetAllSources(c *gin.Context) {
	sources := make([]Source, 2)
	routes := make([]Route, 1)
	routes[0] = Route{URL: "https://www.google.com"}
	sources[0] = Source{Name: "Service1", Routes: routes}
	sources[1] = Source{Name: "Service2", Routes: routes}

	c.JSON(http.StatusOK, gin.H{
		"sources": sources,
	})
}

//CreateRoute is a method to create a Route for a given Source
func (handler HTTPSourceHandler) CreateRoute(c *gin.Context) {
	sourceDTO := SourceDTO{}

	err := c.ShouldBind(&sourceDTO)

	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": "invalid request",
		})
		return
	}

	sourceName := c.Params.ByName("name")
	source := Source{Name: sourceName, Routes: sourceDTO.Routes}
	err = handler.Repository.CreateRoute(source)

	if err != nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{
			"message": "Error storing route information",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "created",
	})

}
