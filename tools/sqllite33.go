package tools

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

func ConnectDB() *sql.DB {
	Db, err := sql.Open("sqlite3","./coolapk.db")
	Check("数据库打开失败！",err)
	return Db
}

//创建数据表
func SqlCreateTable (Db *sql.DB) {

	//defer Db.Close()
	sqlStmt := `CREATE TABLE CoolApkUser(
	ID 	INTEGER PRIMARY KEY AUTOINCREMENT,
	UID				INT    	NOT NULL ,
	USERNAME       TEXT    	NOT NULL,
	USERGROUPNAME  TEXT    	NOT NULL ,
	PROVINCE		TEXT  	NOT NULL ,
	CITY			TEXT 	NOT NULL ,
	FLLOW			INT    	NOT NULL ,
	FANS			INT   	NOT NULL ,
	FEED          	INT   	NOT NULL ,
	GENDER			INT 	NOT NULL
	);`


	_, err := Db.Exec(sqlStmt)
	Check("创建数据表出错！", err )
}

//入库
func SqlInsert(Db *sql.DB,CoolUserData CoolJsons) {
	//defer Db.Close()
	stmt, err := Db.Prepare("INSERT INTO CoolApkUser(UID, USERNAME, USERGROUPNAME, PROVINCE, CITY, FLLOW, FANS, FEED, GENDER) VALUES (?,?,?,?,?,?,?,?,?)")
	_, err = stmt.Exec(CoolUserData.Data.Uid, CoolUserData.Data.Username, CoolUserData.Data.Level,
		CoolUserData.Data.Province, CoolUserData.Data.City, CoolUserData.Data.Follow, CoolUserData.Data.Fans,
		CoolUserData.Data.Feed, CoolUserData.Data.Gender)
	Check("写入数据出错！", err)

}




