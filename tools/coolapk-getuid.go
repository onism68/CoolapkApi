package tools

import (
	"encoding/json"
	"fmt"
	"strings"
)

type UserData struct {
	Data []Data1 `json:"data"`
}

type Data1 struct {
	Username string `json:"username"`
	Uid      string `json:"uid"`
}

//获取用户名的Uid
func GetUid(s string) (uid string, err error) {
	url0 := "https://api.coolapk.com/v6/search?type=user&searchValue=*&page=1"
	url := strings.Replace(url0, "*", s, 1)
	b, err := CoolFetcher(url)
	if err != nil {
		return "", nil
	}
	b1 := UserData{}
	err = json.Unmarshal(b, &b1)
	if err != nil {
		return "", fmt.Errorf("出错了%s", err)
	}
	//fmt.Println(b1)
	if len(b1.Data) > 0 {
		uid := b1.Data[0].Uid
		return uid, nil
	} else {
		return "", fmt.Errorf("出错了%s", err)
	}

}
