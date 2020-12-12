package reqs

import (
	"fmt"
	"strings"
)

func ExtractValue(respBody map[string]interface{}, path string) string {
	var keep string
	pathKeys := strings.Split(path, " ")
	temp := respBody
	for z := 0; z < len(pathKeys); z++ {
		if z == len(pathKeys)-1 {
			keep = fmt.Sprintf("%v", temp[pathKeys[z]])
		} else {
			temp = temp[pathKeys[z]].(map[string]interface{})
		}
		//fmt.Println(temp)
	}

	return keep
}
