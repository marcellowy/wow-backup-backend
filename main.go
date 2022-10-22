package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"

	"github.com/marcellowy/wow-backup-backend/config"
	"github.com/marcellowy/wow-backup-backend/router"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/spf13/viper"

	"github.com/gin-gonic/gin"
)

func initViper() {
	config.Viper = viper.New()
	config.Viper.SetConfigType("yaml")
	fmt.Println(config.File)
	b, err := os.ReadFile(config.File)
	if err != nil {
		panic(err)
	}
	if err = config.Viper.ReadConfig(bytes.NewBuffer(b)); err != nil {
		panic(err)
	}
}

func connectMySQL() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		config.Viper.GetString("database.user"),
		config.Viper.GetString("database.password"),
		config.Viper.GetString("database.host"),
		config.Viper.GetInt("database.port"),
		config.Viper.GetString("database.name"),
	)

	var err error
	if config.Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{}); err != nil {
		panic(err)
	}
}

func init() {
	flag.StringVar(&config.File, "c", "", "指定启动配置文件")
	flag.Parse()
	if config.File == "" {
		panic("配置文件不能为空")
	}
	initViper()    // 初始化配置
	connectMySQL() // 连接数据库
}

func main() {
	gin.SetMode(config.Viper.GetString("mode"))
	engin := gin.Default()
	router.Router(engin)
	engin.Run(config.Viper.GetString("listen")) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
