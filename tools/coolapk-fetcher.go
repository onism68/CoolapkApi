package Coolapk

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func Check(reason string, err error) {
	if err != nil {
		fmt.Println(reason, err)
	}
}

//伪装请求CoolApk的api
func CoolFetcher(url string) []byte {
	token := CoolApkToken()
	client := &http.Client{}
	//生成要访问的url
	//url := "https://api.coolapk.com/v6/main/indexV8?page=1&firstItem=9773448"
	//url = "https://api.coolapk.com/v6/user/space?uid=2688489"//12202
	//提交请求
	reqest, err := http.NewRequest("GET", url, nil)

	//增加header选项
	reqest.Header.Add("user-agent", "Dalvik/2.1.0 (Linux; U; Android 6.0.1; HM2014011 Build/MOB31E) (#Build; Xiaomi; HM2014011; cm_HM2014011-userdebug 6.0.1 MOB31E bdb9819d76 test-keys; 6.0.1) +CoolMarket/8.7")
	reqest.Header.Add("x-requested-with", "XMLHttpRequest")
	reqest.Header.Add("x-app-id", "com.coolapk.market")
	reqest.Header.Add("x-app-token", token)
	reqest.Header.Add("x-app-version", "9.2.1")
	reqest.Header.Add("X-App-Code", "1905221")
	reqest.Header.Add("Host", "api.coolapk.com")
	reqest.Header.Add("Cookie", "uid=10002; username=阿酷; token=97448d8cqTEkx-P4WZ1zc01EHOfd7Q_6E0CCMfg-rqEHb4UCiTzwovA1adXgtvztY6cIX6SDVP_UdtIu4mnsu1z9ElvsNuiSmezAUEi524SuZZYurEgFPwUIZQn5khg5Lo_BMkszvnax-8cM4DF5jSRFSBXurmK3erlcH49zxCzu_VhnA9gSs-VJ_142zToHkVrAyzdY")

	//622540
	if err != nil {
		panic(err)
	}
	//处理返回结果
	response, err := client.Do(reqest)
	Check("打开网页出错",err)
	defer response.Body.Close()

	b, err := ioutil.ReadAll(response.Body)
	Check("读取网页body出错！", err)
	return b
}