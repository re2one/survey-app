package persistence

import (
	//"fmt"
	"log"
	"github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	//"config"
)

func NewDB() *gorm.DB {
	DBMS := "mysql"
	mySqlConfig := &mysql.Config{
		User:                 "survey_admin", //config.C.Database.User,
		Passwd:               "1337", //config.C.Database.Password,
		Net:                  "tcp", //config.C.Database.Net,
		Addr:                 "172.17.0.1:3306", //config.C.Database.Addr,
		DBName:               "SurveyAppDB", // config.C.Database.DBName,
		AllowNativePasswords: true, //config.C.Database.AllowNativePasswords,
		Params: map[string]string{
			"parseTime": "true", //config.C.Database.Params.ParseTime,
		},
	}

	db, err := gorm.Open(DBMS, mySqlConfig.FormatDSN())

	if err != nil {
		log.Fatalln(err)
	}

	return db
}