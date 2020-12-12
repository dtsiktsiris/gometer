package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	//respBody := []byte(`{ "host" : "localhost" ,"person" : [{"firstname": "Jack", "lastname" :"Doe"},{"firstname": "John", "lastname" :"Snow"}]}`)
	//
	//var result map[string]interface{}
	//
	////Unmarshal or Decode the JSON to the interface.
	//json.Unmarshal(respBody, &result)
	//
	//persons := result["person"].([]interface{})
	//
	//firs:= persons[1].(map[string]interface{})
	//	fmt.Println(firs["firstname"])

	str := "persons[32]"

	splitKey := strings.Split(str,"[")
str := splitKey[1][:len(splitKey)-1]
	index := strconv.Atoi(str)
	getvalue := str[1:len(str)-1]
	fmt.Println(getvalue)

}
