package CoolapkApi

import (
	"encoding/json"
	"github.com/onism68/CoolapkApi/tools"
	"log"
)

//获取用户Uid
func GetUserUid(username string) (uid string, err error){
	uid, err = tools.GetUid(username)
	if err != nil {
		log.Println(err)
		return "", err
	}
	return uid, nil
}

//获取用户信息
func GetUserData(username string) (userdata tools.CoolJsons, err error) {
	uid, err := tools.GetUid(username)
	if err != nil {
		log.Println(err)
		return tools.CoolJsons{}, err
	}
	//fmt.Println(uid)
	url := "https://api.coolapk.com/v6/user/space?uid=" + uid
	b := tools.CoolFetcher(url)
	b1 := tools.CoolJsons{}
	json.Unmarshal(b, &b1)
	return b1, nil
}
