package CoolapkApi

import (
	"encoding/json"
	"fmt"
	"github.com/onism68/CoolapkApi/Coolapk"
)

//获取用户Uid
func GetUserUid(username string) {
	b := Coolapk.GetUid(username)
	fmt.Println(b)
}

//获取用户信息
func GetUserData(username string) {
	uid := Coolapk.GetUid(username)
	//fmt.Println(uid)
	url := "https://api.coolapk.com/v6/user/space?uid=" + uid
	b := Coolapk.CoolFetcher(url)
	b1 := Coolapk.CoolJsons{}
	json.Unmarshal(b, &b1)
	fmt.Println(b1)
}
