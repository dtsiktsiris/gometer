package gometer

import (
	"fmt"
	"regexp"
	"strings"
	"sync"
)

type TestResult struct {
	Assertions []string
}

func setDynamicVariables(req *Request, keeper map[string]string) {

	re := regexp.MustCompile("\\${\\w+}")

	if len(re.FindString(req.Url)) > 0 {
		splt := re.FindAllString(req.Url, -1)

		for _, s := range splt {
			//we replase ${mplampla} with keeper['mplampla']
			req.Url = strings.Replace(req.Url, s, keeper[s[2:len(s)-1]], -1)
		}
	}

	if len(req.Body) > 0 && len(re.FindString(req.Body)) > 0 {

		splt := re.FindAllString(req.Body, -1)

		for _, s := range splt {
			//we replase ${mplampla} with keeper['mplampla']
			req.Body = strings.Replace(req.Body, s, keeper[s[2:len(s)-1]], -1)
		}
	}

	for k, v := range req.Header {
		if len(re.FindString(v)) > 0 {
			splt := re.FindAllString(v, -1)

			for _, s := range splt {
				req.Header[k] = strings.Replace(v, s, keeper[s[2:len(s)-1]], -1)
			}
		}
	}
}

func Handle(c *Conf) {
	keeper := make(map[string]string)

	var wg sync.WaitGroup

	fmt.Println("-----Before-----")

	// execute "before" requests
	wg.Add(1)
	handleTests(c.Before, keeper, &wg)
	wg.Wait()

	//we iterate through test sets
	for i := 0; i < len(c.TestSets); i++ {

		fmt.Println("-----Test Set", i+1, "Begin-----")

		for retry := 0; retry < c.TestSets[i].Retries; retry++ {
			wg.Add(1)
			go handleTests(c.TestSets[i].Tests, keeper, &wg)

		}

		wg.Wait()
	}

	fmt.Println("-----After-----")
	// execute "after" requests
	wg.Add(1)
	handleTests(c.After, keeper, &wg)
	wg.Wait()

}

func handleTests(tests []Test, keeper map[string]string, wg *sync.WaitGroup) {
	defer wg.Done()
	for testIndex, test := range tests {

		var assRes []string
		assRes = append(assRes, fmt.Sprintf("*** Test: %v", testIndex+1))

		//check if there is dynamic variable which need to be setted
		//we do it with regex re and search for this form ${mplampla}
		setDynamicVariables(&test.Request, keeper)

		//we do request
		result := test.Request.GetRequestResult()

		//check what to keep
		for k, v := range test.Keep {
			// fmt.Println(k.Path)
			//extractValue return value we want to keep
			//v is the path to this value
			keeper[k] = ExtractValue(result, v)
			// fmt.Println("we keep: ", keeper[k.VariableName])
		}
		//assert happens here
		assRes = Assert(test.Expect, result, assRes)

		for _, s := range assRes {
			fmt.Println(s)
		}

	}

}
