package main

import (
	"bytes"
	"flag"
	"fmt"
	"net/http"
)

func main() {
	fmt.Printf("hello, world\n")
	postTest()
}

func postTest() {
	flag.Parse()
	var flags = flag.Args()
	var postMessage = "body=お疲れ様です。+時間帯ですが、%[1]s時から%[2]s時でいかがでしょうか。場所は%[3]sがいいかと思います"
	requestBody := fmt.Sprintf(postMessage, flags[0], flags[1], flags[2])

	req, err := http.NewRequest("POST", fmt.Sprintf("https://api.chatwork.com/v2/rooms/%[1]s/messages", flags[3]), bytes.NewBufferString(requestBody))
	req.Header.Add("X-ChatWorkToken", flags[4])
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	resp, _ := http.DefaultClient.Do(req)
	resp.Body.Close()
	if err != nil {
	}
	fmt.Println(resp)
}
