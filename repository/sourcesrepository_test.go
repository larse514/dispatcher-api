package repository

import (
	"errors"
	"sync"
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
	return &dynamodb.ScanOutput{}, nil
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
func TestSourceRepositoryInMemoryAddOneSourceGetOneSource(t *testing.T) {
	repo := SourceRepositoryInMemory{Sources: map[string][]handlers.Route{}, Lock: new(sync.Mutex)}
	sources := make([]handlers.Source, 2)
	routes := make([]handlers.Route, 1)
	routes[0] = handlers.Route{URL: "https://www.google.com"}
	sources[0] = handlers.Source{Name: "Service1", Routes: routes}

	repo.CreateRoute(sources[0])
	source, err := repo.GetSource(handlers.Source{Name: sources[0].Name})

	if err != nil {
		t.Log("Error returned when none was expected ", err)
		t.Fail()
	}

	excepted := routes
	actual := source.Routes

	if actual[0] != excepted[0] {
		t.Log("expected ", excepted[0], " actual ", actual[0])
		t.Fail()
	}
}

func TestSourceRepositoryInMemoryAddOneSourceGetAllOneSource(t *testing.T) {
	repo := SourceRepositoryInMemory{Sources: map[string][]handlers.Route{}, Lock: new(sync.Mutex)}
	sources := make([]handlers.Source, 2)
	routes := make([]handlers.Route, 1)
	routes[0] = handlers.Route{URL: "https://www.google.com"}
	sources[0] = handlers.Source{Name: "Service1", Routes: routes}

	repo.CreateRoute(sources[0])
	sources, err := repo.GetAllSources()

	if err != nil {
		t.Log("Error returned when none was expected ", err)
		t.Fail()
	}

	if len(sources) != 1 {
		t.Log("sources is invalid ", sources)
		t.Fail()
	}
}
func TestSourceRepositoryInMemoryAddTwoSourceGetAllTwoSource(t *testing.T) {
	repo := SourceRepositoryInMemory{Sources: map[string][]handlers.Route{}, Lock: new(sync.Mutex)}
	sources := make([]handlers.Source, 2)
	routes := make([]handlers.Route, 1)
	routes[0] = handlers.Route{URL: "https://www.google.com"}
	sources[0] = handlers.Source{Name: "Service1", Routes: routes}
	sources[1] = handlers.Source{Name: "Service2", Routes: routes}

	repo.CreateRoute(sources[0])
	repo.CreateRoute(sources[1])

	sources, err := repo.GetAllSources()

	if err != nil {
		t.Log("Error returned when none was expected ", err)
		t.Fail()
	}

	if len(sources) != 2 {
		t.Log("sources is invalid ", sources)
		t.Fail()
	}
}

func TestSourceRepositoryInMemoryAddTwoSourceSameNameGetAllOneSource(t *testing.T) {
	repo := SourceRepositoryInMemory{Sources: map[string][]handlers.Route{}, Lock: new(sync.Mutex)}
	sources := make([]handlers.Source, 2)
	routes := make([]handlers.Route, 1)
	routes[0] = handlers.Route{URL: "https://www.google.com"}
	sources[0] = handlers.Source{Name: "Service1", Routes: routes}
	sources[1] = handlers.Source{Name: "Service1", Routes: routes}

	repo.CreateRoute(sources[0])
	repo.CreateRoute(sources[1])

	sources, err := repo.GetAllSources()

	if err != nil {
		t.Log("Error returned when none was expected ", err)
		t.Fail()
	}

	if len(sources) != 1 {
		t.Log("sources is invalid ", sources)
		t.Fail()
	}
}
func TestSourceRepositoryInMemoryAddTwoSourceGetTwoSource(t *testing.T) {
	repo := SourceRepositoryInMemory{Sources: map[string][]handlers.Route{}, Lock: new(sync.Mutex)}
	sources := make([]handlers.Source, 2)
	routes := make([]handlers.Route, 2)
	routes[0] = handlers.Route{URL: "https://www.google.com"}
	sources[0] = handlers.Source{Name: "Service1", Routes: routes}
	routes[1] = handlers.Route{URL: "https://www.msn.com"}
	sources[1] = handlers.Source{Name: "Service1", Routes: routes}

	repo.CreateRoute(sources[0])
	repo.CreateRoute(sources[1])

	source, err := repo.GetSource(handlers.Source{Name: sources[0].Name})

	if err != nil {
		t.Log("Error returned when none was expected ", err)
		t.Fail()
	}

	excepted := routes
	actual := source.Routes

	if actual[0] != excepted[0] {
		t.Log("expected ", excepted[0], " actual ", actual[0])
		t.Fail()
	}
	if actual[1] != excepted[1] {
		t.Log("expected ", excepted[0], " actual ", actual[0])
		t.Fail()
	}
}

func TestSourceRepositoryInMemoryAddTwoDifferentSourceGetTwoDifferentSource(t *testing.T) {
	repo := SourceRepositoryInMemory{Sources: map[string][]handlers.Route{}, Lock: new(sync.Mutex)}
	sources := make([]handlers.Source, 2)
	routes := make([]handlers.Route, 1)
	routes[0] = handlers.Route{URL: "https://www.google.com"}
	sources[0] = handlers.Source{Name: "Service1", Routes: routes}
	sources[1] = handlers.Source{Name: "Service2", Routes: routes}

	repo.CreateRoute(sources[0])
	repo.CreateRoute(sources[1])

	source, err := repo.GetSource(handlers.Source{Name: sources[0].Name})

	if err != nil {
		t.Log("Error returned when none was expected ", err)
		t.Fail()
	}

	excepted := routes
	actual := source.Routes

	if actual[0] != excepted[0] {
		t.Log("expected ", excepted[0], " actual ", actual[0])
		t.Fail()
	}

	source, err = repo.GetSource(handlers.Source{Name: sources[1].Name})

	if err != nil {
		t.Log("Error returned when none was expected ", err)
		t.Fail()
	}

	excepted = routes
	actual = source.Routes

	if actual[0] != excepted[0] {
		t.Log("expected ", excepted[0], " actual ", actual[0])
		t.Fail()
	}
}
