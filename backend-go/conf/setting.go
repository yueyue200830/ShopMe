package conf

import (
	"fmt"
	"github.com/go-ini/ini"
)

var (
	Cfg *ini.File

	RunMode string
	HTTPPort int
	JwtSecret string
)

func init() {
	var err error
	Cfg, err = ini.Load("conf/app.ini")
	if err != nil {
		fmt.Println("Fail to load config: ", err)
	}

	loadBase()
	loadServer()
	loadApp()
}

func loadBase() {
	RunMode = Cfg.Section("").Key("RUN_MODE").MustString("debug")
}

func loadServer() {
	sec, err := Cfg.GetSection("server")
	if err != nil {
		fmt.Println("Fail to get server section in config: ", err)
	}
	HTTPPort = sec.Key("HTTP_PORT").MustInt(8000)
}

func loadApp() {
	sec, err := Cfg.GetSection("app")
	if err != nil {
		fmt.Println("Fail to get app section in config: ", err)
	}
	JwtSecret = sec.Key("JWT_SECRET").MustString("Test1Secret")
}

func GetDatabaseSetting() (dbType, con string) {
	sec, err := Cfg.GetSection("database")
	if err != nil {
		fmt.Println("Fail to get database section in config: ", err)
	}

	dbType = sec.Key("TYPE").MustString("mysql")
	dbName := sec.Key("DBNAME").String()
	user := sec.Key("USER").String()
	password := sec.Key("PASSWORD").String()
	host := sec.Key("HOST").MustString("127.0.0.1:3306")

	con = fmt.Sprintf("%s:%s@(%s)/%s?charset=utf8&parseTime=True&loc=Local", user, password, host, dbName)
	return dbType, con
}
