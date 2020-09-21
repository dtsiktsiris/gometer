package main

import (
	"fmt"
	"regexp"
)

func main() {
	//keeped := map[string]string{"url":"localhost", "port":"8080"}
	str := "http://${url}:${port}/"

	re := regexp.MustCompile("\\${\\w+}")
	sp := re.FindString(str)
	fmt.Println(sp[2 : len(sp)-1])
}
