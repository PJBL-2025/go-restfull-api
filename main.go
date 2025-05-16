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
	snapClient := config.InitMidtrans()
	db := config.ConnectDB()

	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}

	err = sqlDB.Ping()
	if err != nil {
		panic(err)
	}

	config.SeedFlag()
	chatController, checkoutController := config.DependencyInjection(config.DB, snapClient)

	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:5500, http://127.0.0.1:5500, http://localhost:5173",
		AllowMethods:     "GET,POST,PUT,DELETE,OPTIONS,PATCH",
		AllowHeaders:     "Origin, Content-Type, Accept,Authorization",
		AllowCredentials: true,
	}))

	router.SetUpRoutes(app, chatController, checkoutController)

	fmt.Println("success")

	err = app.Listen(":3000")
	if err != nil {
		panic(err)
	}

}
