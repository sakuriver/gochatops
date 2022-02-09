package main

import (
	"bytes"
	"flag"
	"fmt"
	"net/http"
)

func main() {
	fmt.Printf("hello, world\n")
	flag.Parse()
	var flags = flag.Args()

	if flags[5] == "1" {
		postMessage(flags)
	} else {
		getMessage(flags)
	}
}

func postMessage(flags []string) {
	flag.Parse()
	var postMessage = "body=お疲れ様です。+時間帯ですが、%[1]s時から%[2]s時でいかがでしょうか。場所は%[3]sがいいかと思います"
	requestBody := fmt.Sprintf(postMessage, flags[0], flags[1], flags[2])
	var url = fmt.Sprintf("https://api.chatwork.com/v2/rooms/%[1]s/messages", flags[3])
	execChatWork(flags, "POST", url, requestBody)
}

func getMessage(flags []string) {
	flag.Parse()
	var url = fmt.Sprintf("https://api.chatwork.com/v2/rooms/%[1]s/messages/%[2]s", flags[3], flags[0])
	execChatWork(flags, "GET", url, "")
	println(url)
}

func execChatWork(flags []string, method string, url string, requestBody string) {

	req, err := http.NewRequest(method, url, bytes.NewBufferString(requestBody))
	req.Header.Add("X-ChatWorkToken", flags[4])
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	resp, _ := http.DefaultClient.Do(req)
	resp.Body.Close()
	if err != nil {
	}
	fmt.Println(resp.Request.Body)

}
