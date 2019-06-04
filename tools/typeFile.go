package tools

//struct里的变量大写才能外部可见！！！

type CoolJsons struct {
	Data CoolUserData `json:"data"`

}

type CoolUserData struct {
	Uid int `json:"uid"` //Uid
	Username string `json:"username"`
	Level int `json:"level"` //等级
	Province string `json:"province"` //省份
	City string `json:"city"` //城市
	Follow int `json:"follow"` //关注
	Fans int `json:"fans"` //粉丝
	Feed int `json:"feed"` //动态数
	Gender int `json:"gender"` //性别
	Avatarstatus int `json:"avatarstatus"`
	Regdate int64 `json:"regdate"`
	Logintime int64 `json:"logintime"`
	Astro string `json:"astro"`
	Weibo string `json:"weibo"`
	Blog string `json:"blog"`
	Bio string `json:"bio"`

}
