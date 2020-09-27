package main
import (
	"../reqs"
	"fmt"
)
func main() {
	//keeped := map[string]string{"url":"localhost", "port":"8080"}
	//str := "http://${url}:${port}/"
	//
	//re := regexp.MustCompile("\\${\\w+}")
	//sp := re.FindString(str)
	//fmt.Println(sp[2 : len(sp)-1])

	var c reqs.Conf
	//load yaml file to Conf
	c.GetConf()
	fmt.Println(c.TestSets[0].Request.Body)
}
