package config

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

var DB *gorm.DB

func LoadConfig() {
	viper.AutomaticEnv()
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	viper.SetDefault("APP_PORT", "3000")
}

func ConnectDB() *gorm.DB {
	username := viper.GetString("DB_USERNAME")
	password := viper.GetString("DB_PASSWORD")
	host := viper.GetString("DB_HOST")
	port := viper.GetString("DB_PORT")
	database := viper.GetString("DB_DATABASE")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		username, password, host, port, database)

	var err error
	maxRetries := 10

	for i := 0; i < maxRetries; i++ {
		DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err == nil {
			fmt.Println("✅ Database connected!")
			return DB
		}
		fmt.Printf("⏳ Attempt %d: failed to connect to DB: %v\n", i+1, err)
		time.Sleep(3 * time.Second)
	}

	panic("🔥 Could not connect to database after retries")
}
