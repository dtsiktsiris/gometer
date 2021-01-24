package main

import (
	"../reqs"
)

func main() {

	//jsonPath := "../../simple.json"
	jsonPath := "../../requests.json"

	var c reqs.Conf
	//load yaml file to Conf
	c.GetConf(jsonPath)
	reqs.Handle(&c)
}
