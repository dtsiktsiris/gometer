package main

import (
	"../../gometer"
)

func main() {

	jsonPath := "./testserver_example_test.json"

	gometer.InitWithFile(jsonPath)
}
