package infrastructure

import (
	"github.com/larse514/aws-cloudformation-go"
	"github.com/larse514/dispatcher-api/handlers"
)

const (
	queueNameKey    = "QueueName"
	tagKey          = "dispatcher"
	tagValue        = "dispatcher-lambda"
	lambdaS3Code    = "s3://lambda.bucket.moodle"
	lambdaSourceKey = "LambdaSource"
)

// //RouterCreator is an interface to define the method to create a list of Routers
// type RouterCreator interface {
// 	CreateRouters(routes []Route) error
// 	CreateRoutersWithSource(source Source) error
// }

//LambdaRouterCreator is the struct implementation of the sources.RouterCreator
type LambdaRouterCreator struct {
	Executor cf.Executor
	Template IaasAsset
	// Packager TemplatePackager
}

//IaasAsset is an interface to retrieve a required IaaS template
type IaasAsset interface {
	GetRouterTemplate() (*string, error)
	GetSourceTemplate() (*string, error)
	// GetLambdaSource() (*string, error)
}

//TemplatePackager is an interface to store lambda source code
// type TemplatePackager interface {
// 	UploadSourceCode() error
// }

//CreateRouters creates Lambda Routers based on provided clise of Routes
func (lambda LambdaRouterCreator) CreateRouters(source *handlers.Source) error {
	template, err := lambda.Template.GetRouterTemplate()
	if err != nil {
		return err
	}
	return lambda.Executor.CreateStack(*template, source.Name, nil, nil)
}

//CreateRoutersWithSource creates Lambda Routers and SQS Queue based on provided clise of Routes
func (lambda LambdaRouterCreator) CreateRoutersWithSource(source *handlers.Source) error {
	template, err := lambda.Template.GetSourceTemplate()

	if err != nil {
		return err
	}

	return lambda.Executor.CreateStack(*template, source.Name, createTemplateParameters(source.Name), createTags())
}

//helper function to create a map of Cloudformation Parameters
func createTemplateParameters(name string) *map[string]string {
	m := map[string]string{
		queueNameKey:    name,
		lambdaSourceKey: lambdaS3Code,
	}
	return &m
}

//helper function to create tags
func createTags() *map[string]string {
	m := map[string]string{
		tagKey: tagValue,
	}
	return &m
}
