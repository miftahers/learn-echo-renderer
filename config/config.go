package config

import (
	"fmt"
	"learn-echo-renderer/model"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Config struct {
	DBUser      string
	DBPassword  string
	DBPort      string
	DBHost      string
	DBName      string
	APIPort     string
	TokenSecret string
}

var DB *gorm.DB
var Cfg *Config

func InitConfig() {
	cfg := &Config{}

	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	err = viper.Unmarshal(&cfg)
	if err != nil {
		panic(err)
	}

	Cfg = cfg

}

func InitDB() {
	DBUser := Cfg.DBUser
	DBPassword := Cfg.DBPassword
	DBHost := Cfg.DBHost
	DBPort := Cfg.DBPort
	DBName := Cfg.DBName

	connectionString := fmt.
		Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
			DBUser,
			DBPassword,
			DBHost,
			DBPort,
			DBName)

	db, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	DB = db

	err = DB.AutoMigrate(
		&model.User{},
	)

	if err != nil {
		panic(err)
	}
}
