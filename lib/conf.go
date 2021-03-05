package lib

import (
	"encoding/json"
	"log"
)

type Conf struct {
	Variables map[string]interface{} `json:"variables"`
	Before    []Test                 `json:"before"`
	TestSets  []TestSet              `json:"test_sets"`
	After     []Test                 `json:"after"`
}

type TestSet struct {
	Name    string `json:"name"`
	Retries int    `json:"retries"`
	Tests   []Test `json:"tests"`
}

type Test struct {
	Name    string            `json:"name"`
	Request Request           `json:"request"`
	Expect  Expect            `json:"expect"`
	Keep    map[string]string `json:"keep"`
}

type Request struct {
	Method string            `json:"method"`
	Url    string            `json:"url"`
	Body   string            `json:"body"`
	Header map[string]string `json:"header"`
}

type Expect struct {
	StatusCode int               `json:"statusCode"`
	Assertions map[string]string `json:"assertions"`
}

func LoadConf(jsonContent []byte) {

	var c Conf

	err := json.Unmarshal(jsonContent, &c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	//fmt.Printf("%+v",c)
	handle(&c)
}
