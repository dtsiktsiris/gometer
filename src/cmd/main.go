package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strings"

	"../reqs"
)

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

	//we iterate through test sets
	for i := 0; i < len(c.TestSets[0].Tests); i++ {

		// go func(){

		// }

		fmt.Println("-----Test Set Begin-----")

		//check if there is dynamic variable which need to be setted
		//we do it with regex re and search for this form ${mplampla}
		setDynamicVariables(&c.TestSets[0].Tests[i].Request, keeper)

		//we do request
		resp, err := c.TestSets[0].Tests[i].Request.Resolve()
		if err != nil {
			// handle error
		}
		body, readErr := ioutil.ReadAll(resp.Body)
		if readErr != nil {
			log.Fatal(readErr)
		}
		defer resp.Body.Close()

		//Declared an empty interface
		var result map[string]interface{}

		//Unmarshal or Decode the JSON to the interface.
		json.Unmarshal(body, &result)
		fmt.Println("body:", result)
		//check if there is anything to keep
		if len(c.TestSets[0].Tests[i].Keep) > 0 {

			for k, v := range c.TestSets[0].Tests[i].Keep {
				//extractValue return value we want to keep
				//v is the path to this value
				keeper[k] = reqs.ExtractValue(result, v)
				fmt.Println("we keep: ", keeper[k])
			}

		}
		//assert happens here
		reqs.Assert(c.TestSets[0].Tests[i].Expect, resp, result)
	}

}
