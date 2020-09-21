package reqs

import (
	"bytes"
	"net/http"
)

func (request *Request) Resolve() (*http.Response, error) {

	switch request.Method {
	case "GET" :
		resp, err := http.Get(request.Url)
		return resp, err
	case "POST" :
		resp, err := http.Post(request.Url,"application/json",bytes.NewBufferString("oti na ne"))
		return resp, err
	case "PUT" :
		req, err := http.NewRequest(http.MethodPut, request.Url,bytes.NewBufferString("oti na ne"))
		resp, err := http.DefaultClient.Do(req)
		return resp, err


	default:
		return nil, nil

	}

	//if err != nil {
	//	// handle error
	//}

}
