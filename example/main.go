package main

import (
	"github.com/dtsiktsiris/gometer"
)

func main() {

	// jsonFilePath := os.Args[1]
	// gometer.InitWithFile(jsonFilePath)

	gometer.InitWithJSONString(getJSONString())
}

func getJSONString() string {
	json := `{
		"variables": {
		  "baseUrl": "www.localhost.com",
		  "port": "8080"
		},
		"before": [
		  {
			"name": "test 1",
			"request": {
			  "method": "GET",
			  "url": "http://localhost:8080/"
			},
			"expect": {
			  "statusCode": 200,
			  "assertions": {
				"persons[0] firstname": "John",
				"enviroments[0]": "develop"
			  }
			},
			"keep": {
			  "name": "persons[0] lastname",
			  "local": "host"
			}
		  }
		],
		"test_sets": [
		  {
			"name": "test set 1",
			"retries": 1,
			"tests": [
			  {
				"name": "test 1",
				"request": {
				  "method": "GET",
				  "url": "http://${local}:8080/"
				},
				"expect": {
				  "statusCode": 202
				},
				"keep": {
				  "loc": "host"
				}
			  }
			]
		  }
		]
	  }`
	return json
}
