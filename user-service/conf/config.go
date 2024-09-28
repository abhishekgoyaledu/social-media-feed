package conf

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/social-media/user-service/dao"
	"github.com/social-media/user-service/database"
)

const (
	envVar = "USER_SERVICE_CONF"
)

var DB *gorm.DB

type RedisConfig struct {
	Host           string `json:"host"`
	Port           int    `json:"port"`
	MaxConn        int    `json:"max_conn"`
	ConnectTimeout int    `json:"connect_timeout"`
	ReadTimeout    int    `json:"read_timeout"`
	WriteTimeout   int    `json:"write_timeout"`
}

var Config = &struct {
	InternalHostAndPort string                     `json:"internalHostAndPort"`
	MySQL               *database.MysqlMasterSlave `json:"mysql"`
	Redis               RedisConfig                `json:"redis"`
}{}

func init() {
	LoadJSONEnvPathOrPanic(envVar, Config)
	connectDatabase()
	fmt.Print("Config loaded...")
}

func connectDatabase() {
	dsn := Config.MySQL.Master.Dsn
	dialector := gorm.Dialector(mysql.Open(dsn))
	db, err := gorm.Open(dialector)

	//defer db.Close()
	if err != nil {
		panic("Connection Failed to Open")
	}
	log.Println("Connection Established")

	db.AutoMigrate(&dao.UserProfile{})

	DB = db
}
