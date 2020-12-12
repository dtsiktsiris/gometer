package reqs

import (
	"fmt"
	"strconv"
	"strings"
)

func ExtractValue(respBody map[string]interface{}, path string) string {
	var keep string
	pathKeys := strings.Split(path, " ")
	temp := respBody
	for z := 0; z < len(pathKeys); z++ {

		//fmt.Println(temp, pathKeys[z])
		if strings.Contains(pathKeys[z], "[") {
			//index := pathKeys[z][1:len(pathKeys[z])-1]

			splitKey := strings.Split(pathKeys[z], "[")
			//fmt.Println("splitkey", splitKey)
			index, err := strconv.Atoi(splitKey[1][:len(splitKey)-1])
			if err != nil {
				fmt.Println("Error converting string to int", err)
			}
			tempArray := temp[splitKey[0]].([]interface{})

			temp = tempArray[index].(map[string]interface{})
		} else if z == len(pathKeys)-1 {
			keep = fmt.Sprintf("%v", temp[pathKeys[z]])
		} else {
			temp = temp[pathKeys[z]].(map[string]interface{})
		}
	}

	return keep
}
