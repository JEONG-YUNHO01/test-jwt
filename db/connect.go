package db

import (
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect() *gorm.DB {
	USER := os.Getenv("DBUSER")       // DB 유저명
	PASS := os.Getenv("DBPASS")       // DB 유저의 패스워드
	PROTOCOL := "tcp(localhost:3306)" // 개발환경이므로 localhost의 3306포트로 설정한다.
	DBNAME := os.Getenv("DBNAME")     // 사용할 DB 명을 입력

	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(CONNECT), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}

	return db
}
