package main

import (
	f "./func"
	"fmt"
	"os"
	"strings"
)

func main() {
	var weiboUid string
	var isok float64

	isok = 1
	has_nextpage := true

	fmt.Print("please type uid: ")
	fmt.Scanln(&weiboUid)
	sinceId := ""

	for has_nextpage {
		pageContent := getPageContent(weiboUid, sinceId)

		pageMap := f.JsonToMap(pageContent)

		data := pageMap["data"].(map[string]interface{})["cards"]

		if isok != pageMap["ok"] {
			break
		}
		for _, v := range data.([]interface{}) {
			item := v.(map[string]interface{})
			mblog := item["mblog"].(map[string]interface{})
			id := mblog["id"]
			sinceId = id.(string)
			text := mblog["text"]
			fmt.Println(text)
		}

		fmt.Print("next page(y/n):")
		var type_res string
		fmt.Scanln(&type_res)
		if strings.ToLower(type_res) != "y" {
			has_nextpage = false
		}
	}

	fmt.Printf("Press any key to exit...")
	b := make([]byte, 1)
	os.Stdin.Read(b)
}

func getPageContent(weiboUid string, sinceId string) string {
	url := "https://m.weibo.cn/api/container/getIndex?uid=" + weiboUid + "&t=0&type=uid&value=" + weiboUid + "&containerid=107603" + weiboUid + "&since_id=" + sinceId
	return f.GetHttpResult(url)
}
