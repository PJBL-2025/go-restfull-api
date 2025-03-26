package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"restfull-api-pjbl-2025/config"
	"restfull-api-pjbl-2025/router"
)

func main() {
	config.LoadConfig()
	db := config.ConnectDB()

	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}

	err = sqlDB.Ping()
	if err != nil {
		panic(err)
	}

	config.SeedFlag(db)
	chatController := config.DependencyInjection(config.DB)

	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:5173",
		AllowMethods:     "GET,POST,PUT,DELETE,OPTIONS,PATCH",
		AllowHeaders:     "Origin, Content-Type, Accept,Authorization",
		AllowCredentials: true,
	}))

	router.SetUpRoutes(app, chatController)

	fmt.Println("success")

	err = app.Listen(":3000")
	if err != nil {
		panic(err)
	}

}
