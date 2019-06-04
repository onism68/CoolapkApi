package main

import (
	"coolapk/tools"
	"encoding/json"
	"fmt"
)

func main() {
	s := "酷安小编"
	uid := tools.GetUid(s)
	url := "https://api.coolapk.com/v6/user/space?uid=" + uid
	b := tools.CoolFetcher(url)
	b1 := tools.CoolJsons{}
	json.Unmarshal(b, &b1)
	fmt.Println(b1)
}