package config

import (
	"flag"
	"fmt"
	"os"
	"restfull-api-pjbl-2025/helper"
)

func SeedFlag() {
	seedToken := flag.Bool("token", false, "token")

	flag.Parse()

	if *seedToken {
		helper.TokenSeeder()
		fmt.Println("Seeder Token berhasil dijalankan!")
		os.Exit(0)
	}
}
