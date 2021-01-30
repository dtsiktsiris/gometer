package main

import (
	"../../gometer"
)

func main() {

	jsonPath := "./requests.json"
	//load yaml file to Conf
	gometer.GetConf(jsonPath)
}
