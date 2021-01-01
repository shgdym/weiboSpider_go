package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	var url string
	weibo_uid := "" // 要爬取的 微博UID
	url = "https://m.weibo.cn/api/container/getIndex?uid=" + weibo_uid + "&t=0&type=uid&value=" + weibo_uid + "&containerid=107603" + weibo_uid

	pageContent := getHttpResult(url)
	pageMap := JSONToMap(pageContent)

	data := pageMap["data"].(map[string]interface{})["cards"]

	for _, v := range data.([]interface{}) {
		item := v.(map[string]interface{})
		text := item["mblog"].(map[string]interface{})["text"]

		fmt.Println(text)
	}

	select {}
}

func getHttpResult(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("error")
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	return string(body)
}

func JSONToMap(str string) map[string]interface{} {
	var tempMap map[string]interface{}
	err := json.Unmarshal([]byte(str), &tempMap)
	if err != nil {
		panic(err)
	}
	return tempMap
}
