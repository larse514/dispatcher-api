package handlers

import (
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
	Routes             []Route `json:"routes"`
	WithSourceCreation bool    `json:"withSourceCreation"`
}

// HTTPSourceHandler struct implementation of SourceHandler for HTTP requests
type HTTPSourceHandler struct {
	Repository    SourceRepository
	Dynamo        SourceRepository
	RouterCreator RouterCreator
}

//SourceRepository is an interface used by the sources.go handler to interact with storage
type SourceRepository interface {
	CreateRoute(source Source) error
	GetSource(source Source) (Source, error)
	GetAllSources() ([]Source, error)
}

//RouterCreator is an interface to define the method to create a list of Routers
type RouterCreator interface {
	CreateRouters(source *Source) error
	CreateRoutersWithSource(source *Source) error
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
	sources, err := handler.Repository.GetAllSources()

	if err != nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{
			"message": "Error retrieving route information",
		})
		return
	}

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
