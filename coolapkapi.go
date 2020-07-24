package CoolapkApi

import (
	"encoding/json"
	"github.com/onism68/CoolapkApi/tools"
	"log"
)

//获取用户Uid
func GetUserUid(username string) (uid string, err error) {
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
	data, err := tools.CoolFetcher(url)
	if err != nil {
		return tools.CoolJsons{}, err
	}
	userdata = tools.CoolJsons{}
	err = json.Unmarshal(data, &username)
	if err != nil {
		return tools.CoolJsons{}, err
	}
	return userdata, nil
}

func GetTouTiao() (string, error) {
	url := "https://api.coolapk.com/v6/main/indexV8"
	data, err := tools.CoolFetcher(url)
	if err != nil {
		return "", err
	}
	return string(data), nil
}
