package reqs

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v3"
)

type Conf struct {
	Variables map[string]interface{} `yaml:"variables"`
	Before    []TestSet              `yaml:"before"`
	TestSets  []TestSet              `yaml:"test_sets"`
	After     []TestSet              `yaml:"after"`
}

type TestSet struct {
	Request Request           `yaml:"request"`
	Expect  Expect            `yaml:"expect"`
	Keep    map[string]string `yaml:"keep"`
}

type Request struct {
	Method string `yaml:"method"`
	Url    string `yaml:"url"`
	Body   string `yaml:"body"`
}

type Expect struct {
	StatusCode int               `yaml:"statusCode"`
	Assertions map[string]string `yaml:"assertions"`
}

func (c *Conf) GetConf(yamlPath string) *Conf {

	yamlFile, err := ioutil.ReadFile(yamlPath)
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return c
}
