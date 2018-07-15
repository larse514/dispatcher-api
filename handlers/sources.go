package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// SourceHandler is an interface to handle Source HTTP requests
type SourceHandler interface {
	GetAllSources(c *gin.Context)
	CreateRoute(c *gin.Context)
	GetRoutes(c *gin.Context)
}

//Source is the struct representation of a source object
//it contains a name and a route object
type Source struct {
	Name   string  `json:"name"`
	Routes []Route `json:"routes"`
}

//Route is the struct representation of a source object.  It contains a Url
type Route struct {
	URL string `json:"url"`
}

//SourceDTO is a struct representing a POST sources request
type SourceDTO struct {
	Route              Route `json:"route"`
	WithSourceCreation bool  `json:"withSourceCreation"`
}

// HTTPSourceHandler struct implementation of SourceHandler for HTTP requests
type HTTPSourceHandler struct {
	Dynamo        SourceRepository
	RouterCreator RouterCreator
}

//SourceRepository is an interface used by the sources.go handler to interact with storage
type SourceRepository interface {
	CreateRoute(source Source) error
	GetSource(source Source) (Source, error)
	GetAllSources() ([]Source, error)
	GetRouteForSource(sourceName string, routeName string) (Route, error)
}

//RouterCreator is an interface to define the method to create a list of Routers
type RouterCreator interface {
	CreateRouters(source *Source) error
	CreateRoutersWithSource(source *Source) error
}

//NotFoundError is an error corresponding to a resource not found
type NotFoundError struct {
	Resource string
}

func (e NotFoundError) Error() string {
	return fmt.Sprintf("Resource %s not found", e.Resource)

}

//GetRoutes is a function to return a source based on the requested Source Name
func (handler HTTPSourceHandler) GetRoutes(c *gin.Context) {
	sourceName := c.Params.ByName("name")

	log.Printf("DEBUG: received name: %s", sourceName)

	source, err := handler.Dynamo.GetSource(Source{Name: sourceName})

	if err != nil {
		log.Println("ERROR: received error from Repository: ", err)

		c.JSON(http.StatusServiceUnavailable, gin.H{
			"message": "Error retrieving source",
		})
		return
	}
	log.Println("DEBUG: returning routes from source: ", source)
	c.JSON(http.StatusOK, gin.H{

		"routes": source.Routes,
	})
}

//GetAllSources is a method to return all source object
func (handler HTTPSourceHandler) GetAllSources(c *gin.Context) {
	sources := append(make([]Source, 0), Source{Name: "SOMENAME", Routes: append(make([]Route, 0), Route{URL: "SOMEURL"})})

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
	routes := append(make([]Route, 0), sourceDTO.Route)
	source := Source{Name: sourceName, Routes: routes}

	_, err = handler.Dynamo.GetRouteForSource(sourceName, sourceDTO.Route.URL)

	_, ok := err.(NotFoundError)

	if ok {
		c.JSON(http.StatusConflict, gin.H{
			"message": "Route exists for Source",
		})
		return
	}

	err = handler.Dynamo.CreateRoute(source)

	if err != nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{
			"message": "Error storing route information",
		})
		return
	}

	if sourceDTO.WithSourceCreation {
		log.Println("DEBUG: Source Creation set to true, creating Sources")
		err = handler.RouterCreator.CreateRoutersWithSource(&source)
		if err != nil {
			c.JSON(http.StatusServiceUnavailable, gin.H{
				"message": "Error creating Routers and Source",
			})
			return
		}
	} else {
		log.Println("DEBUG: Creating Routers")

		err = handler.RouterCreator.CreateRouters(&source)
	}

	if err != nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{
			"message": "Error creating Routers",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "created",
	})

}
