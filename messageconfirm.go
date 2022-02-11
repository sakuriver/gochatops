package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
)

var apiUrl = "https://api.chatwork.com/"
var requestBodyByte []byte

func main() {
	fmt.Printf("hello, world\n")
	flag.Parse()
	var flags = flag.Args()

	if flags[5] == "1" {
		postMessage(flags)
	} else if flags[5] == "2" {
		getMessage(flags)
	} else {
		getMembersFromRoom(flags)

	}
}

// リアル打ち合わせお問い合わせ
func postMessage(flags []string) {
	flag.Parse()
	var postMessage = "body=お疲れ様です。+時間帯ですが、%[1]s時から%[2]s時でいかがでしょうか。場所は%[3]sがいいかと思います"
	requestBody := fmt.Sprintf(postMessage, flags[0], flags[1], flags[2])
	var url = fmt.Sprintf("%[1]sv2/rooms/%[2]s/messages", apiUrl, flags[3])
	execChatWork(flags, "POST", url, requestBody)
}

// メッセージを取得
func getMessage(flags []string) {
	flag.Parse()
	var url = fmt.Sprintf("%[1]sv2/rooms/%[2]s/messages/%[3]s", apiUrl, flags[3], flags[0])
	execChatWork(flags, "GET", url, "")
	println(url)
}

// 日報を連絡するところにいるメンバー一覧を取得する
func getMembersFromRoom(flags []string) {
	println("日報の連絡所のメンバー一覧を取得します(daily report room member get)")
	flag.Parse()
	var url = fmt.Sprintf("%[1]sv2/rooms/%[2]s/members", apiUrl, flags[3])
	execChatWork(flags, "GET", url, "")
	println(url)

	members := []Member{}
	json.Unmarshal(requestBodyByte, &members)
	println(fmt.Sprintf("%[1]d 人いることを確認しました", len(members)))
	// 部屋にいるメンバーを取得する
	for _, member := range members {
		println(fmt.Sprintf("アカウントId: %[1]d アカウント画像 %[2]s", member.AccountId, member.AvatarImageUrl))
	}

}

type Member struct {
	AccountId      int    `json:"account_id"`
	AvatarImageUrl string `json:"avatar_image_url"`
}

func execChatWork(flags []string, method string, url string, requestBody string) {

	var req = &http.Request{}
	if method == "GET" {
		req, _ = http.NewRequest(method, url, nil)
	} else {
		req, _ = http.NewRequest(method, url, bytes.NewBufferString(requestBody))
	}

	req.Header.Add("X-ChatWorkToken", flags[4])
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	resp, _ := http.DefaultClient.Do(req)
	if resp != nil {
		println(resp.Body)
		requestBodyByte, _ = io.ReadAll(resp.Body)
		resp.Body.Close()
	}

}
