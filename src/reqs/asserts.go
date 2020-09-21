package reqs

import (
	"fmt"
	"net/http"
)

func Assert(e Expect, response *http.Response){
	if e.StatusCode == response.StatusCode {
		fmt.Println("status code: pass")
	} else {
		fmt.Println("assert fail")
	}
}
