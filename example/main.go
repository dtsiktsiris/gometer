package main

import (
	"../../gometer"
)

func main() {

	jsonPath := "./testserver_example_test.json"
	//load json file to Conf
	gometer.GetConf(jsonPath)
}
