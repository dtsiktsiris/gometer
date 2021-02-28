package gometer

import (
	"log"
	"os"
)

func InitWithFile(jsonFilePath string) {

	jsonContent, err := os.ReadFile(jsonFilePath)
	if err != nil {
		log.Printf("Read file err   #%v ", err)
	}

	LoadConf(jsonContent)
}

func InitWithJSONString(jsonString string) {

	if jsonString == "" {
		log.Printf("JSON is empty. Func: InitWithJSON.")
	}

	LoadConf([]byte(jsonString))
}
