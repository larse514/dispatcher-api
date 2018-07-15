package main

import (
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/gin-gonic/gin"
	"github.com/larse514/aws-cloudformation-go"
	"github.com/larse514/dispatcher-api/assets"
	"github.com/larse514/dispatcher-api/handlers"
	"github.com/larse514/dispatcher-api/infrastructure"
	"github.com/larse514/dispatcher-api/repository"
)

const (
	tableNameKey = "TableName"
	regionKey    = "Region"
)

var ginLambda *ginadapter.GinLambda

// Handler is the main entry point for Lambda. Receives a proxy request and
// returns a proxy response
func Handler(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// stdout and stderr are sent to AWS CloudWatch Logs
	log.Printf("main handler")
	r := gin.Default()
	// Initialize a session in us-west-2 that the SDK will use to load
	// credentials from the shared credentials file ~/.aws/credentials.
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(os.Getenv(regionKey))},
	)

	if err != nil {
		fmt.Println("FATAL: Error creating dynamodb session:")
		fmt.Println(err.Error())
		os.Exit(1)
	}
	// Create DynamoDB client
	svc := dynamodb.New(sess)
	cloudformation := cloudformation.New(sess)

	tableName := os.Getenv(tableNameKey)
	dynamodb := repository.SourceDynamoDBRepository{Svc: svc, TableName: tableName}

	executor := cf.IaaSExecutor{Client: cloudformation}
	lambda := assets.AWSTemplate{}
	routercreator := infrastructure.LambdaRouterCreator{Executor: executor, Template: lambda}
	sourceHandler := handlers.HTTPSourceHandler{Dynamo: dynamodb, RouterCreator: routercreator}

	r.GET("/sources", sourceHandler.GetAllSources)
	r.GET("/sources/:name/routes", sourceHandler.GetRoutes)
	r.POST("/sources/:name/routes", sourceHandler.CreateRoute)

	r.GET("/ping", handlers.Ping)
	ginLambda = ginadapter.New(r)

	return ginLambda.Proxy(req)
}

func main() {
	lambda.Start(Handler)
}
