package gometer

import (
	"io/ioutil"
	"log"
)

func InitWithFile(jsonPath string) {

	jsonContent, err := ioutil.ReadFile(jsonPath)
	if err != nil {
		log.Printf("Read file err   #%v ", err)
	}

	LoadConf(jsonContent)
}

func InitWithJSON(jsonString string) {

	if jsonString == "" {
		log.Printf("JSON is empty. Func: InitWithJSON.")
	}

	LoadConf([]byte(jsonString))
}
