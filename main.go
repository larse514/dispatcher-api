package main

import (
	"log"
	"sync"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/gin-gonic/gin"
	"github.com/larse514/dispatcher-api/handlers"
	"github.com/larse514/dispatcher-api/repository"
)

var ginLambda *ginadapter.GinLambda

// Handler is the main entry point for Lambda. Receives a proxy request and
// returns a proxy response
func Handler(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// stdout and stderr are sent to AWS CloudWatch Logs
	log.Printf("main handler")
	r := gin.Default()
	repo := repository.SourceRepositoryInMemory{Sources: map[string][]handlers.Route{}, Lock: new(sync.Mutex)}
	sourceHandler := handlers.HTTPSourceHandler{Repository: repo}
	routeHandler := handlers.HTTPRouteHandler{}

	r.GET("/sources", sourceHandler.GetAllSources)
	r.GET("/sources/:name/routes", routeHandler.GetAllRoutes)
	r.POST("/sources/:name/routes", sourceHandler.CreateRoute)

	r.GET("/ping", handlers.Ping)
	ginLambda = ginadapter.New(r)

	return ginLambda.Proxy(req)
}

func main() {
	lambda.Start(Handler)
}
