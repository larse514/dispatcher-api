package repository

import (
	"errors"
	"testing"

	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"

	"github.com/larse514/dispatcher-api/handlers"
)

// Define a mock to return a basic success
type mockGoodDynamoDbClient struct {
	dynamodbiface.DynamoDBAPI
}
type mockBadDynamoDbClient struct {
	dynamodbiface.DynamoDBAPI
}

func (client mockGoodDynamoDbClient) Scan(*dynamodb.ScanInput) (*dynamodb.ScanOutput, error) {
	int64Ptr := int64(1)
	return &dynamodb.ScanOutput{Count: &int64Ptr}, nil
}

func (client mockBadDynamoDbClient) Scan(*dynamodb.ScanInput) (*dynamodb.ScanOutput, error) {
	int64Ptr := int64(0)
	return &dynamodb.ScanOutput{Count: &int64Ptr}, nil
}
func (client mockGoodDynamoDbClient) PutItem(*dynamodb.PutItemInput) (*dynamodb.PutItemOutput, error) {
	return &dynamodb.PutItemOutput{}, nil
}
func (client mockBadDynamoDbClient) PutItem(*dynamodb.PutItemInput) (*dynamodb.PutItemOutput, error) {
	return &dynamodb.PutItemOutput{}, errors.New("An error occured")
}

func TestSourceDynamoDBRepositoryScanNoError(t *testing.T) {
	repo := SourceDynamoDBRepository{Svc: mockGoodDynamoDbClient{}}

	_, err := repo.GetSource(handlers.Source{Name: "Service1", Routes: nil})

	if err != nil {
		t.Log("Error returned when none was expected ", err)
		t.Fail()
	}
}

func TestSourceDynamoDBRepositoryGetRouteForSourceNoCountReturnsNotFoundError(t *testing.T) {
	repo := SourceDynamoDBRepository{Svc: mockBadDynamoDbClient{}}

	_, err := repo.GetRouteForSource("service", "route")
	_, ok := err.(handlers.NotFoundError)
	if !ok {
		t.Log("Error not of correct type NotFoundError ", err)
		t.Fail()
	}
}

func TestSourceDynamoDBRepositoryGetRouteForSource1CountReturnsNoError(t *testing.T) {
	repo := SourceDynamoDBRepository{Svc: mockGoodDynamoDbClient{}}

	_, err := repo.GetRouteForSource("service", "route")
	if err != nil {
		t.Log("Error returned when none was expected ", err)
		t.Fail()
	}
}
func TestSourceDynamoDBRepositoryAddOneSourceNoError(t *testing.T) {
	repo := SourceDynamoDBRepository{Svc: mockGoodDynamoDbClient{}}
	sources := make([]handlers.Source, 2)
	routes := make([]handlers.Route, 1)
	routes[0] = handlers.Route{URL: "https://www.google.com"}
	sources[0] = handlers.Source{Name: "Service1", Routes: routes}

	err := repo.CreateRoute(sources[0])

	if err != nil {
		t.Log("Error returned when none was expected ", err)
		t.Fail()
	}
}
func TestSourceDynamoDBRepositoryAddOneSourceWithErrorReturnsError(t *testing.T) {
	repo := SourceDynamoDBRepository{Svc: mockBadDynamoDbClient{}}
	sources := make([]handlers.Source, 2)
	routes := make([]handlers.Route, 1)
	routes[0] = handlers.Route{URL: "https://www.google.com"}
	sources[0] = handlers.Source{Name: "Service1", Routes: routes}

	err := repo.CreateRoute(sources[0])

	if err == nil {
		t.Log("Error not returned when one expected")
		t.Fail()
	}
}
