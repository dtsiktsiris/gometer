package reqs

import (
	"fmt"
)

func Assert(e Expect, result map[string]interface{}) {
	fmt.Print("expect (", result["statusCode"], ") be equal to (", e.StatusCode, ") : ")
	if e.StatusCode == result["statusCode"] {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
	for k, v := range e.Assertions {
		respValue := ExtractValue(result, k)
		if v == respValue {
			fmt.Println("expect (", respValue, ") be equal to (", v, ") : PASS")
		} else {

			fmt.Println("expect (", respValue, ") be equal to (", v, ") : FAIL")
		}
	}
}
