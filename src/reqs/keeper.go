package reqs

import (
	"fmt"
	"strings"
)

func ExtractValue(respBody map[string]interface{}, path string) string {
	var keep string
	pathKeys := strings.Split(path, " ")
	var temp map[string]interface{}
	for z := 0; z < len(pathKeys); z++ {
		if len(pathKeys) == 1 {
			keep = fmt.Sprintf("%v", respBody[pathKeys[z]])
		} else if z == 0 {
			temp = respBody[pathKeys[z]].(map[string]interface{})
		} else if z == len(pathKeys)-1 {
			keep = fmt.Sprintf("%v", temp[pathKeys[z]])
		} else {
			temp = temp[pathKeys[z]].(map[string]interface{})
		}
		//fmt.Println(temp)
	}

	return keep
}
