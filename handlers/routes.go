package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//Route is the struct representation of a source object.  It contains a Url
type Route struct {
	URL string `json:"url"`
}

//GetAllRoutes is a method to return all source object
func GetAllRoutes(c *gin.Context) {
	routes := make([]Route, 2)
	routes[0] = Route{URL: "https://www.google.com"}
	routes[1] = Route{URL: "https://www.msn.com"}

	c.JSON(http.StatusOK, gin.H{
		"routes": routes,
	})
}
