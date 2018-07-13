package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//Source is the struct representation of a source object
//it contains a name and a route object
type Source struct {
	Name  string `json:"name"`
	Route Route  `json:"route"`
}

//GetAllSources is a method to return all source object
func GetAllSources(c *gin.Context) {
	sources := make([]Source, 2)
	sources[0] = Source{Name: "Service1", Route: Route{URL: "https://www.google.com"}}
	sources[1] = Source{Name: "Service2", Route: Route{URL: "https://www.google.com"}}

	c.JSON(http.StatusOK, gin.H{
		"sources": sources,
	})
}
