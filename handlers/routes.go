package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//Route is the struct representation of a source object.  It contains a Url
type Route struct {
	URL string `json:"url"`
}

// HTTPRouteHandler struct implementation of RouteHandler for HTTP requests
type HTTPRouteHandler struct {
	repository SourceRepository
}

// RouteHandler is an interface to handle Source HTTP requests
type RouteHandler interface {
	GetAllRoutes(c *gin.Context)
}

//GetAllRoutes is a method to return all source object
func (handler HTTPRouteHandler) GetAllRoutes(c *gin.Context) {
	routes := make([]Route, 2)
	routes[0] = Route{URL: "https://www.google.com"}
	routes[1] = Route{URL: "https://www.msn.com"}

	c.JSON(http.StatusOK, gin.H{
		"routes": routes,
	})
}
