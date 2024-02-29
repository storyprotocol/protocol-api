package openapi

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type OpenchainResponse struct {
	Ok     bool
	Result Result
}

type Result struct {
	Event    interface{}
	Function map[string][]Functions
}

type Functions struct {
	Name     string
	Filtered bool
}

func NewOpenChainClient(requestUrl string) *OpenchainClient {
	return &OpenchainClient{
		requestURL:      requestUrl,
		cachedFunctions: map[string]string{},
	}
}

type OpenchainClient struct {
	//client *http.Client
	requestURL      string
	cachedFunctions map[string]string
}

func (ocClient *OpenchainClient) GetPermissionName(encodedName string) (string, error) {
	val, ok := ocClient.cachedFunctions[encodedName]
	if ok {
		fmt.Printf("Retrieving permission from cached functions")
		return val, nil
	}

	fmt.Printf("Getting permission from openchain api")
	requestUrl := fmt.Sprintf("%s?function=%s", ocClient.requestURL, encodedName)

	res, err := http.Get(requestUrl)
	if err != nil {
		fmt.Printf("error making http request: %s\n", err)
		return encodedName, err
	}

	decodedResponse := &OpenchainResponse{}
	err = json.NewDecoder(res.Body).Decode(decodedResponse)
	if err != nil {
		fmt.Printf("error decoding request: %s\n", err)
		return encodedName, err
	}

	if !decodedResponse.Ok {
		fmt.Printf("error in response: %s\n", err)
		return encodedName, err
	}

	ocClient.cachedFunctions[encodedName] = decodedResponse.Result.Function[encodedName][0].Name

	return decodedResponse.Result.Function[encodedName][0].Name, nil
}
