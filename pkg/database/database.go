package database

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var (
	Eloquent *sql.DB
)


func Init() func() {
	var err error
	dsn := mysqlConfig()
	Eloquent, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	err = Eloquent.Ping()
	if err != nil {
		log.Fatal(err)
		return nil
	}

	return func() {
		err = Eloquent.Close()
		if err != nil {
			log.Fatal(err)
		}
	}

}

func mysqlConfig() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			"wangxin",
				"MD_wangxin3#4$",
				"127.0.0.1",
				"3306",
				"learning",
		)
}

