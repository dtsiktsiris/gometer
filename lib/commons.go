package lib

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func setDynamicVariables(req *Request, keeper map[string]string) {

	re := regexp.MustCompile("\\${\\w+}")

	if len(re.FindString(req.Url)) > 0 {
		splt := re.FindAllString(req.Url, -1)

		for _, s := range splt {
			//we replase ${blabla} with keeper['blabla']
			req.Url = strings.Replace(req.Url, s, keeper[s[2:len(s)-1]], -1)
		}
	}

	if len(req.Body) > 0 && len(re.FindString(req.Body)) > 0 {

		splt := re.FindAllString(req.Body, -1)

		for _, s := range splt {
			//we replase ${mplampla} with keeper['mplampla']
			req.Body = strings.Replace(req.Body, s, keeper[s[2:len(s)-1]], -1)
		}
	}

	for k, v := range req.Header {
		if len(re.FindString(v)) > 0 {
			splt := re.FindAllString(v, -1)

			for _, s := range splt {
				req.Header[k] = strings.Replace(v, s, keeper[s[2:len(s)-1]], -1)
			}
		}
	}
}

func extractValue(respBody map[string]interface{}, path string) string {
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
			if z == len(pathKeys)-1 {
				keep = fmt.Sprintf("%v", tempArray[index])
			} else {
				temp = tempArray[index].(map[string]interface{})
			}
		} else {
			if z == len(pathKeys)-1 {
				keep = fmt.Sprintf("%v", temp[pathKeys[z]])
			} else {
				temp = temp[pathKeys[z]].(map[string]interface{})
			}
		}
	}

	return keep
}
