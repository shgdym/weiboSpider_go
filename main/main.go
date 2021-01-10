package main

import (
	f "./func"
	"fmt"
)

func main() {
	var url string
	var weiboUid string

	fmt.Print("请输入微博uid：")
	fmt.Scanln(&weiboUid)
	url = "https://m.weibo.cn/api/container/getIndex?uid=" + weiboUid + "&t=0&type=uid&value=" + weiboUid + "&containerid=107603" + weiboUid

	pageContent := f.GetHttpResult(url)
	pageMap := f.JsonToMap(pageContent)

	data := pageMap["data"].(map[string]interface{})["cards"]

	for _, v := range data.([]interface{}) {
		item := v.(map[string]interface{})
		text := item["mblog"].(map[string]interface{})["text"]

		fmt.Println(text)
	}

	select {}
}
