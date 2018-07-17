package infrastructure

import (
	"errors"
	"testing"

	"github.com/larse514/dispatcher-api/handlers"
)

type mockGoodExecutor struct {
}

//CreateStack is a general method to create aws cloudformation stacks
func (executor mockGoodExecutor) CreateStack(templateBody string, stackName string, parameterMap *map[string]string, tags *map[string]string) error {
	return nil
}
func (executor mockGoodExecutor) PauseUntilCreateFinished(stackName string) error {
	return nil
}
func (executor mockGoodExecutor) UpdateStack(templateBody string, stackName string, parameterMap *map[string]string, tags *map[string]string) error {
	return nil
}

func (executor mockGoodExecutor) PauseUntilUpdateFinished(stackName string) error {
	return nil
}

type mockBadExecutor struct {
}

//CreateStack is a general method to create aws cloudformation stacks
func (executor mockBadExecutor) CreateStack(templateBody string, stackName string, parameterMap *map[string]string, tags *map[string]string) error {
	return errors.New("")
}
func (executor mockBadExecutor) PauseUntilCreateFinished(stackName string) error {
	return errors.New("")
}
func (executor mockBadExecutor) UpdateStack(templateBody string, stackName string, parameterMap *map[string]string, tags *map[string]string) error {
	return errors.New("")
}

func (executor mockBadExecutor) PauseUntilUpdateFinished(stackName string) error {
	return errors.New("")
}

type mockGoodIaasAsset struct {
}

func (iaas mockGoodIaasAsset) GetRouterTemplate() (*string, error) {
	m := "ROUTERTEMPLATE"
	return &m, nil
}

func (iaas mockGoodIaasAsset) GetSourceTemplate() (*string, error) {
	m := "SOURCETEMPLATE"

	return &m, nil

}

type mockGoodTemplatePackager struct {
}

func (packager mockGoodTemplatePackager) UploadSourceCode() error {
	return nil
}

func TestLambdaRouterCreatorCreateRoutersWithSource(t *testing.T) {
	lambda := LambdaRouterCreator{Template: mockGoodIaasAsset{}, Executor: mockGoodExecutor{}}

	err := lambda.CreateRoutersWithSource(&handlers.Source{Routes: append(make([]handlers.Route, 0), handlers.Route{URL: "URL"})})

	if err != nil {
		t.Log("Error encounterd ", err.Error())
		t.Fail()
	}

}

func TestLambdaRouterCreatorCreateRoutersWithSourceBadExecutorReturnsError(t *testing.T) {
	lambda := LambdaRouterCreator{Template: mockGoodIaasAsset{}, Executor: mockBadExecutor{}}

	err := lambda.CreateRoutersWithSource(&handlers.Source{Routes: append(make([]handlers.Route, 0), handlers.Route{URL: "URL"})})

	if err == nil {
		t.Log("Error not encountered")
		t.Fail()
	}

}
