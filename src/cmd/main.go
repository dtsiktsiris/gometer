package main

import (
	"../reqs"
)

func main() {

	//jsonPath := "../../simple.json"
	jsonPath := "../../requests.json"

	//load yaml file to Conf
	reqs.GetConf(jsonPath)
}
