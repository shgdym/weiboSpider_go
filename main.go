package main

import (
	f "./func"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	var weiboUid string
	var isok float64

	fmt.Print("please type uid: ")
	fmt.Scanln(&weiboUid)
	sinceId := ""

	isok = 1
	get_nextpage := true

	// 去除html标签
	re, _ := regexp.Compile("\\<[\\S\\s]+?\\>")

	for get_nextpage {
		// 获取api返回信息 默认sinceId为空 即返回第一页
		pageContent := getPageContent(weiboUid, sinceId)
		pageMap := f.JsonToMap(pageContent)

		// 异常
		if isok != pageMap["ok"] {
			break
		}

		// 解析数据
		data := pageMap["data"].(map[string]interface{})["cards"]
		for _, v := range data.([]interface{}) {
			item := v.(map[string]interface{})
			mblog := item["mblog"].(map[string]interface{})
			text := re.ReplaceAllString(mblog["text"].(string), "") // 微博内容
			fmt.Println(text)
			sinceId = mblog["id"].(string)
		}

		// 是否获取下一页
		fmt.Print("next page(y/n):")
		var type_res string
		fmt.Scanln(&type_res)
		if strings.ToLower(type_res) != "y" {
			get_nextpage = false
		}
	}

	fmt.Printf("Press any key to exit...")
	b := make([]byte, 1)
	os.Stdin.Read(b)
}

// 拼接url, 进行网络请求
func getPageContent(weiboUid string, sinceId string) string {
	url := "https://m.weibo.cn/api/container/getIndex?uid=" + weiboUid + "&t=0&type=uid&value=" + weiboUid + "&containerid=107603" + weiboUid + "&since_id=" + sinceId
	return f.GetHttpResult(url)
}
