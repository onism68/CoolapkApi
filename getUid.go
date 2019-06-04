package main

import (
	"fmt"
	"github.com/onism68/CoolapkApi/Coolapk"
)


func main() {
	b := Coolapk.GetUid("酷安小编")
	fmt.Println(b)
}
