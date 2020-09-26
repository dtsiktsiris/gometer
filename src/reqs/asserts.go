package reqs

import (
	"fmt"
	"net/http"
)

func Assert(e Expect, response *http.Response, body map[string]interface{}) {

	if e.StatusCode == response.StatusCode {
		fmt.Println("status code: pass")
	} else {
		fmt.Println("assert fail")
	}
	for k, v := range e.Assertions {
 		if v == ExtractValue(body,k){
 			fmt.Println("we match")
		}
		//fmt.Println(k, v)
	}
}
