package config

import (
	"flag"
	"fmt"
	"gorm.io/gorm"
	"os"
	"restfull-api-pjbl-2025/db/seeders"
)

func SeedFlag(db *gorm.DB) {
	seedToken := flag.Bool("token", false, "token")

	flag.Parse()

	if *seedToken {
		seeders.TokenSeeder()
		fmt.Println("Seeder Token berhasil dijalankan!")
		os.Exit(0)
	}
}
