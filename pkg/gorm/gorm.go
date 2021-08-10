package gorm

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/spf13/viper"
	"gostu/pkg/config"
	"log"
)

var Eloquent *gorm.DB

func Init() func() {
	var err error
	if learn := config.AppConfig.Sub("mysql.learn"); learn != nil {
		Eloquent, err = gorm.Open("mysql", mysqlConfig(learn))
		if err != nil {
			log.Fatal(err)
			return nil
		}
		dbConfig(Eloquent, learn)
	}

	return func() {
		if err = Eloquent.Close(); err != nil {
			log.Fatal(err)
		}
	}
}

func mysqlConfig(c * viper.Viper) string {
	//username := "wangxin"
	//password := "MD_wangxin3#4$"
	//host := "127.0.0.1"
	//port := "3306"
	//database := "learning"
	//
	//var buffer bytes.Buffer
	//buffer.WriteString(username)
	//buffer.WriteString(":")
	//buffer.WriteString(password)
	//buffer.WriteString("@tcp(")
	//buffer.WriteString(host)
	//buffer.WriteString(":")
	//buffer.WriteString(port)
	//buffer.WriteString(")/")
	//buffer.WriteString(database)
	//buffer.WriteString("?charset=utf8mb4&parseTime=True&loc=Local")
	//return buffer.String()

	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		c.GetString("username"),
		c.GetString("password"),
		c.GetString("host"),
		c.GetString("port"),
		c.GetString("database"),
	)
}

func dbConfig(db *gorm.DB, c *viper.Viper)  {
	db.DB().SetMaxIdleConns(c.GetInt("max_idle_conns"))
	db.DB().SetMaxOpenConns(c.GetInt("max_open_conns"))
	db.DB().SetConnMaxLifetime(c.GetDuration("conn_max_lifetime"))
	db.SingularTable(true)
	db.LogMode(config.AppConfig.GetBool("debug"))//设置为true之后控制台会输出对应的SQL语句
}
