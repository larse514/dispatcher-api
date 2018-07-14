package repository

import (
	"fmt"
	"sync"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
	"github.com/larse514/dispatcher-api/handlers"
	uuid "github.com/satori/go.uuid"
)

type dynamoSource struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Route string `json:"route"`
}

// SourceRepositoryInMemory struct containing in memory list
type SourceRepositoryInMemory struct {
	Sources map[string][]handlers.Route
	Lock    *sync.Mutex
}

// SourceDynamoDBRepository struct containing in dynamodb connection
type SourceDynamoDBRepository struct {
	Svc       dynamodbiface.DynamoDBAPI
	TableName string
}

// GetSource method to get a slice of routes for a Source
func (repo SourceDynamoDBRepository) GetSource(source handlers.Source) (handlers.Source, error) {
	filt := expression.Name("name").Equal(expression.Value(source.Name))
	proj := expression.NamesList(expression.Name("id"), expression.Name("name"), expression.Name("route"))
	expr, err := expression.NewBuilder().WithFilter(filt).WithProjection(proj).Build()

	if err != nil {
		fmt.Println("Got error building expression:")
		fmt.Println(err.Error())
		return handlers.Source{}, err
	}

	params := &dynamodb.ScanInput{
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		FilterExpression:          expr.Filter(),
		ProjectionExpression:      expr.Projection(),
		TableName:                 aws.String(repo.TableName),
	}

	// Make the DynamoDB Query API call
	result, err := repo.Svc.Scan(params)

	if err != nil {
		fmt.Println("Query API call failed:")
		fmt.Println((err.Error()))
		return handlers.Source{}, err
	}

	fmt.Println("Query returned ", result)

	for _, i := range result.Items {
		route := handlers.Route{}
		err = dynamodbattribute.UnmarshalMap(i, &route)
		if err != nil {
			fmt.Println("Got error unmarshalling:")
			fmt.Println(err.Error())
			return handlers.Source{}, err
		}
		source.Routes = append(source.Routes, route)
	}
	return source, nil
}

// GetAllSources method to get a slice of Sources
func (repo SourceDynamoDBRepository) GetAllSources() ([]handlers.Source, error) {
	return nil, nil
}

// CreateRoute method to create a route
func (repo SourceDynamoDBRepository) CreateRoute(source handlers.Source) error {
	for _, route := range source.Routes {
		u1 := uuid.Must(uuid.NewV4())
		attribute, err := dynamodbattribute.MarshalMap(dynamoSource{ID: u1.String(), Name: source.Name, Route: route.URL})

		if err != nil {
			fmt.Println("Got error marshalling map:")
			fmt.Println(err.Error())
		}

		// Create item in table
		input := &dynamodb.PutItemInput{
			Item:      attribute,
			TableName: aws.String(repo.TableName),
		}

		//put item
		_, err = repo.Svc.PutItem(input)

		if err != nil {
			fmt.Println("Got error calling PutItem:")
			fmt.Println(err.Error())
			return err
		}

		fmt.Println("Successfully added Source")
	}

	return nil
}

// CreateRoute method to create a route
func (repo SourceRepositoryInMemory) CreateRoute(source handlers.Source) error {
	repo.Lock.Lock()
	defer repo.Lock.Unlock()

	repo.Sources[source.Name] = append(repo.Sources[source.Name], source.Routes...)
	return nil
}

// GetSource method to get a slice of routes for a Source
func (repo SourceRepositoryInMemory) GetSource(source handlers.Source) (handlers.Source, error) {
	repo.Lock.Lock()
	defer repo.Lock.Unlock()
	source.Routes = repo.Sources[source.Name]
	return source, nil
}

// GetAllSources method to get all sources
func (repo SourceRepositoryInMemory) GetAllSources() ([]handlers.Source, error) {
	repo.Lock.Lock()
	defer repo.Lock.Unlock()
	sources := make([]handlers.Source, 0)
	for key, value := range repo.Sources {
		fmt.Println("Key:", key, "Value:", value)
		sources = append(sources, handlers.Source{Name: key, Routes: value})
	}

	return sources, nil
}
