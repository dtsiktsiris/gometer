package lib

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func (r *Request) getRequestResult() map[string]interface{} {
	resp, err := r.resolve()
	if err != nil {
		log.Fatal(err)
	}
	body, readErr := io.ReadAll(resp.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}
	defer resp.Body.Close()

	//Declared an empty interface
	var result map[string]interface{}

	//Unmarshal or Decode the JSON to the interface.
	json.Unmarshal(body, &result)
	//fmt.Println("body:", result)

	result["statusCode"] = resp.StatusCode

	return result
}

func (r *Request) resolve() (*http.Response, error) {

	req, err := http.NewRequest(r.Method, r.Url, bytes.NewBufferString(r.Body))

	for k, v := range r.Header {
		//TODO check differences for .Add .Set and concatenacion
		req.Header.Add(k, v)
	}

	if err != nil {
		fmt.Println("Error in request resolve preparing request", err)
	}

	req.Header.Add("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		fmt.Println("Error in request resolve getting resp", err)
	}

	return resp, err

}
