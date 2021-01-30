package gometer

import (
	"fmt"
)

func Assert(e Expect, result map[string]interface{}, assertionResults []string) []string {

	if e.StatusCode == result["statusCode"] {
		assertionResults = append(assertionResults, fmt.Sprintf("expect ( %v ) be equal to ( %v ) : PASS", result["statusCode"], e.StatusCode))
	} else {
		assertionResults = append(assertionResults, fmt.Sprintf("expect ( %v ) be equal to ( %v ) : FAIL", result["statusCode"], e.StatusCode))
	}
	for k, v := range e.Assertions {
		respValue := ExtractValue(result, k)
		if v == respValue {
			assertionResults = append(assertionResults, fmt.Sprintf("expect ( %v ) be equal to ( %v ) : PASS", respValue, v))
		} else {

			assertionResults = append(assertionResults, fmt.Sprintf("expect ( %v ) be equal to ( %v ) : FAIL", respValue, v))
		}
	}

	return assertionResults
}
