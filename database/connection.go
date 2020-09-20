package database

import (
	"RandomChatting_Server/config"
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
)

type connectionMethod interface {
	Connect()
}

// DB - 데이터베이스 전역변수
var DB *gorm.DB

// Connect - 데이터베이스 구조 생성, 연결 하는 메서드
func Connect() {
	dbConf := config.Config.DB

	connectOptions := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		dbConf.Host,
		dbConf.Port,
		dbConf.Username,
		dbConf.Name,
		dbConf.Password)

	db, err := gorm.Open("postgres", connectOptions)

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(
		&Member{},
	)

	DB = db

	log.Print("[DATABASE] 연결 완료")
}
