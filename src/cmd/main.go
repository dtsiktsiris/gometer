package main

import (
	"fmt"
	"regexp"
	"strings"
	"sync"

	"../reqs"
)

type TestResult struct {
	Assertions []string
}

func setDynamicVariables(req *reqs.Request, keeper map[string]string) {

	re := regexp.MustCompile("\\${\\w+}")

	if len(re.FindString(req.Url)) > 0 {
		splt := re.FindAllString(req.Url, -1)

		for _, s := range splt {
			//we replase ${mplampla} with keeper['mplampla']
			req.Url = strings.Replace(req.Url, s, keeper[s[2:len(s)-1]], -1)
		}
	}
	//TODO we also need to check/replace for body and header
}

func main() {

	// yamlPath := "../../simple.yaml"
	yamlPath := "../../requests.yaml"

	var c reqs.Conf
	//load yaml file to Conf
	c.GetConf(yamlPath)
	keeper := make(map[string]string)

	fmt.Println("-----Before-----")

	// execute "before" requests
	for _, test := range c.Before {
		setDynamicVariables(&test.Request, keeper)

		//we do request
		result := test.Request.GetRequestResult()

		if len(test.Keep) > 0 {

			for k, v := range test.Keep {
				//extractValue return value we want to keep
				//v is the path to this value
				keeper[k] = reqs.ExtractValue(result, v)
				fmt.Println("we keep: ", keeper[k])
			}

		}
		//assert happens here
		reqs.Assert(test.Expect, result)
	}

	var wg sync.WaitGroup
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
	for _, test := range c.After {
		setDynamicVariables(&test.Request, keeper)

		//we do request
		result := test.Request.GetRequestResult()

		if len(test.Keep) > 0 {

			for k, v := range test.Keep {
				//extractValue return value we want to keep
				//v is the path to this value
				keeper[k] = reqs.ExtractValue(result, v)
				fmt.Println("we keep: ", keeper[k])
			}

		}
		//assert happens here
		reqs.Assert(test.Expect, result)
	}

}
func handleTests(tests []reqs.Test, keeper map[string]string, wg *sync.WaitGroup) {
	defer wg.Done()
	for testIndex, test := range tests {

		fmt.Println("*** Test: ", testIndex+1)

		//check if there is dynamic variable which need to be setted
		//we do it with regex re and search for this form ${mplampla}
		setDynamicVariables(&test.Request, keeper)

		//we do request
		result := test.Request.GetRequestResult()

		//check what to keep
		for k, v := range test.Keep {
			//extractValue return value we want to keep
			//v is the path to this value
			keeper[k] = reqs.ExtractValue(result, v)
			fmt.Println("we keep: ", keeper[k])
		}

		//assert happens here
		reqs.Assert(test.Expect, result)

	}

}
