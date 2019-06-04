package main

import (
	"encoding/json"
	"fmt"
	"github.com/onism68/CoolapkApi/Coolapk"
)

func main() {
	s := "酷安小编"
	uid := Coolapk.GetUid(s)
	url := "https://api.coolapk.com/v6/user/space?uid=" + uid
	b := Coolapk.CoolFetcher(url)
	b1 := Coolapk.CoolJsons{}
	json.Unmarshal(b, &b1)
	fmt.Println(b1)
}