package reqs

import (
	"fmt"
	"net/http"
)

func Assert(e Expect, response *http.Response, body map[string]interface{}) {
	fmt.Print("expect (", response.StatusCode, ") be equal to (", e.StatusCode, ") : ")
	if e.StatusCode == response.StatusCode {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
	for k, v := range e.Assertions {
		respValue := ExtractValue(body, k)
		if v == respValue {
			fmt.Println("expect (", respValue, ") be equal to (", v, ") : PASS")
		} else {

			fmt.Println("expect (", respValue, ") be equal to (", v, ") : FAIL")
		}
	}
}
