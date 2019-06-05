package tools

import (
	"crypto/md5"
	"encoding/base64"
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

var token1 string = "token://com.coolapk.market/c67ef5943784d09750dcfbb31020f0ab?"

var package_name string = "com.coolapk.market"

func base64Encode(data string) (string){

	return base64.StdEncoding.EncodeToString([]byte(data))
}

//func base64Decode(data string) string {
////	d, err :=  base64.StdEncoding.DecodeString(data)
////	if err!= nil {
////		fmt.Println(err)
////	}
////	return string(d)
////}


//func reverseStr(str string) (string){
//	var revS string
//	strlen := len(str)
//	for i := 0 ; i < strlen ; i++ {
//		revS = revS + fmt.Sprintf("%c", str[strlen-i-1])
//	}
//	return revS
//}

func md5_hex_digest (str string ) (string){


	data := []byte(str)
	has := md5.Sum(data)
	md5str1 := fmt.Sprintf("%x", has) //将[]byte转成16进制

	return md5str1
}

func CoolApkToken () string {
	a1 :=rand.Int()%100000000
	var device_id string =  strconv.Itoa(a1) + "-0000-0000-0000-000000000000"
	timestamp := time.Now().Unix() + 100
	timestamp_md5 := md5_hex_digest(strconv.FormatInt(timestamp,10))
	salt := token1 + timestamp_md5 + "$" + device_id + "&" + package_name
	saltEncoded := base64Encode(salt)
	saltMd5 := md5_hex_digest(saltEncoded)
	return saltMd5 + device_id + fmt.Sprintf("0x%x",timestamp)

}

