package main

import (
	"../reqs"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strings"
)

func main() {
	var c reqs.Conf
	c.GetConf()
	keeper := make(map[string]string)

	re := regexp.MustCompile("\\${\\w+}")

	for i := 0; i < len(c.TestSets); i++ {

		if len(re.FindString(c.TestSets[i].Request.Url)) > 0 {
			splt := re.FindAllString(c.TestSets[i].Request.Url, -1)

			for _, s := range splt {
				c.TestSets[i].Request.Url = strings.Replace(c.TestSets[i].Request.Url, s, keeper[s[2:len(s)-1]], -1)
			}
		}

		resp, err := c.TestSets[i].Request.Resolve()
		if err != nil {
			// handle error
		}

		if len(c.TestSets[i].Keep) > 0 {

			body, readErr := ioutil.ReadAll(resp.Body)
			if readErr != nil {
				log.Fatal(readErr)
			}
			defer resp.Body.Close()

			//Declared an empty interface
			var result map[string]interface{}

			//Unmarshal or Decode the JSON to the interface.
			json.Unmarshal(body, &result)
			//fmt.Println(result)

			for k, v := range c.TestSets[i].Keep {
				keeper[k] = reqs.ExtractValue(result, v)
				fmt.Println("we keep: ", keeper[k])
			}

		}

		reqs.Assert(c.TestSets[i].Expect, resp)
	}

}
