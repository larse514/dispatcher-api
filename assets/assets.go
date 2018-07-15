package assets

const (
	sourceTemplate = "cfn-transformed-template.yml"
	routerTemplate = "sam.yml"
	lambdaSource   = "index.js"
)

// GetRouterTemplate() (string, error)
// GetSourceTemplate() (string, error)
// GetLambdaSource() (string, error)

//AWSTemplate is the implementation of IaasAsset interface
type AWSTemplate struct {
}

//GetRouterTemplate method to retrieve Router Template
func (template AWSTemplate) GetRouterTemplate() (*string, error) {
	return getAsset(routerTemplate)
}

//GetSourceTemplate method to retrieve Source Template
func (template AWSTemplate) GetSourceTemplate() (*string, error) {
	return getAsset(sourceTemplate)
}

// //GetLambdaSource method to retrieve Lamda Source Code
// func (template AWSTemplate) GetLambdaSource() (*string, error) {
// 	return getAsset(lambdaSource)
// }

//getAsset retrives []byte of file and converts it to a string
//todo- this can probably be factored out again
func getAsset(template string) (*string, error) {
	//todo- think through if it's worth factoring out into an interface for mocking?
	data, err := Asset(template)
	if err != nil {
		println("Error reading file ", template)
		return nil, err
	}
	s := string(data[:])
	return &s, nil
}
