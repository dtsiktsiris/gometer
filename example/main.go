package main

import (
	"github.com/ditsikts/gometer"
)

func main() {

	jsonPath := "./testserver_example_test.json"

	gometer.InitWithFile(jsonPath)
}
