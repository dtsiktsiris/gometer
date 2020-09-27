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

func main() {
	var c reqs.Conf
	//load yaml file to Conf
	yamlPath := "../../requests.yaml"
	c.GetConf(yamlPath)
	keeper := make(map[string]string)

	re := regexp.MustCompile("\\${\\w+}")

	//we iterate through test sets
	for i := 0; i < len(c.TestSets); i++ {

		//check if there is dynamic variable which need to be setted
		//we do it with regex re and search for this form ${mplampla}
		if len(re.FindString(c.TestSets[i].Request.Url)) > 0 {
			splt := re.FindAllString(c.TestSets[i].Request.Url, -1)

			for _, s := range splt {
				//we replase ${mplampla} with keeper['mplampla']
				c.TestSets[i].Request.Url = strings.Replace(c.TestSets[i].Request.Url, s, keeper[s[2:len(s)-1]], -1)
			}
		}

		//we do request
		resp, err := c.TestSets[i].Request.Resolve()
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
		fmt.Println(result)
		//check if there is anything to keep
		if len(c.TestSets[i].Keep) > 0 {

			for k, v := range c.TestSets[i].Keep {
				//extractValue return value we want to keep
				//v is the path of this value
				keeper[k] = reqs.ExtractValue(result, v)
				fmt.Println("we keep: ", keeper[k])
			}

		}
		//assert happens here, currently only status
		reqs.Assert(c.TestSets[i].Expect, resp, result)
	}

}
