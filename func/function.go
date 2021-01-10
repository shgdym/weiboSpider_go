package f

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func GetHttpResult(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("error")
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	return string(body)
}

func JsonToMap(str string) map[string]interface{} {
	var tempMap map[string]interface{}
	err := json.Unmarshal([]byte(str), &tempMap)
	if err != nil {
		panic(err)
	}
	return tempMap
}
