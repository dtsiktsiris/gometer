package gometer

import (
	"fmt"
	"sync"
)

func handleTests(tests []Test, keeper map[string]string, wg *sync.WaitGroup) {
	defer wg.Done()
	for testIndex, test := range tests {

		var assRes []string
		assRes = append(assRes, fmt.Sprintf("*** Test: %v", testIndex+1))

		//check if there is dynamic variable which need to be setted
		//we do it with regex re and search for this form ${mplampla}
		setDynamicVariables(&test.Request, keeper)

		//we do request
		result := test.Request.getRequestResult()

		//check what to keep
		for k, v := range test.Keep {
			//extractValue return value we want to keep
			//v is the path to this value
			keeper[k] = extractValue(result, v)
			// fmt.Println("we keep: ", keeper[k.VariableName])
		}
		//assert happens here
		assRes = assert(test.Expect, result, assRes)

		for _, s := range assRes {
			fmt.Println(s)
		}

	}

}
