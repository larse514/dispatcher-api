# Dispatch API   [![CircleCI](https://circleci.com/gh/larse514/dispatcher-api.svg?style=svg)](https://circleci.com/gh/larse514/dispatcher-api) [![Go Report Card](https://goreportcard.com/badge/github.com/larse514/dispatcher-api)](https://goreportcard.com/report/github.com/larse514/dispatcher-api)

API responsible for maintaining Source to Route information

API's include:

### 1. POST /sources/{sourcename}/routes

**Body:**</br>
```json
{
	"route": {
		"url": "String"
	},
	"withSourceCreation": "boolean"
}
```

**Description:**</br>
Store route based on sourcename and create Lambda Route function that will bind to SQS queue.  if withSourceCreation is set to true, then create the source as well.  For now this will only support SQS (could include SNS in the future).  


### 2. GET /sources

**Description:**</br>
Retrieve all sources that are maintained by dispatcher

### 3. GET /sources/{sid}/routes
Retrieve all routes for a given source


## Deployment
### Application code
The Dispatcher API is built using [aws lambda go](https://github.com/aws/aws-lambda-go) for integration with [aws lambda](https://aws.amazon.com/lambda/) and the [golang gin library](https://gin-gonic.github.io/gin/) 
### Build Code
``` shell
$ make
```
This will execute the clean, dependencies, test, build, and package tasks

### Deploy
The Dispatcher API uses the [Serverless Application Model](https://github.com/awslabs/serverless-application-model) for deployment.  The API is defined in the sam.yml file.  This includes 2 steps:

1. Package assets

``` shell
$ aws cloudformation package \
    --template-file $(SAM_FILE) \
    --output-template-file $(SAM_OUTPUT) \
    --s3-bucket $(DEPLOYMENT_BUCKET) 

```

2. Deploy assets
``` shell
$ aws cloudformation deploy \
    --template-file $(SAM_OUTPUT) \
    --stack-name $(STACK_NAME) \
    --capabilities CAPABILITY_IAM \
    --parameter-overrides Environment=$(ENV)

```
Note that these can be found in the Makefile under the package and deploy tasks