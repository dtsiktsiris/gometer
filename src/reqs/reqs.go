package reqs

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func (request *Request) GetRequestResult() map[string]interface{} {
	resp, err := request.Resolve()
	if err != nil {
		log.Fatal(err)
	}
	body, readErr := ioutil.ReadAll(resp.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}
	defer resp.Body.Close()

	//Declared an empty interface
	var result map[string]interface{}

	//Unmarshal or Decode the JSON to the interface.
	json.Unmarshal(body, &result)
	fmt.Println("body:", result)

	result["statusCode"] = resp.StatusCode

	return result
}

func (request *Request) Resolve() (*http.Response, error) {

	switch request.Method {
	case "GET":
		resp, err := http.Get(request.Url)
		return resp, err
	case "POST":
		req, err := http.NewRequest(http.MethodPost, request.Url, bytes.NewBufferString(request.Body))
		req.Header.Add("Content-Type", "application/json")
		resp, err := http.DefaultClient.Do(req)
		return resp, err
	case "PUT":
		req, err := http.NewRequest(http.MethodPut, request.Url, bytes.NewBufferString(request.Body))
		resp, err := http.DefaultClient.Do(req)
		return resp, err
	default:
		return nil, nil

	}

	//if err != nil {
	//	// handle error
	//}

}
