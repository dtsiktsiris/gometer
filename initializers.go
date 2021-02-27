package gometer

import (
	"os"
	"log"
)

func InitWithFile(jsonPath string) {

	jsonContent, err := os.ReadFile(jsonPath)
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
