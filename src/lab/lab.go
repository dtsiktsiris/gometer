package main

import (
	"encoding/json"
	"fmt"
)
func main() {
	respBody := []byte(`{ "host" : "localhost" ,"person" : [{"firstname": "Jack", "lastname" :"Doe"},{"firstname": "John", "lastname" :"Snow"}]}`)

	var result map[string]interface{}

	//Unmarshal or Decode the JSON to the interface.
	json.Unmarshal(respBody, &result)

	persons := result["person"].([]interface{})

	firs:= persons[1].(map[string]interface{})
		fmt.Println(firs["firstname"])



}
