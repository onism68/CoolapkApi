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
	b := CoolFetcher(url)
	b1 := UserData{}
	json.Unmarshal(b, &b1)
	//fmt.Println(b1)
	if len(b1.Data) > 0 {
		uid := b1.Data[0].Uid
		return uid, nil
	} else {
		return "",fmt.Errorf("出错了%s",err)
	}

}
