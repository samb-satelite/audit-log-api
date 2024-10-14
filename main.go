package main

import (
	"audit-log/src/interfaces/rest"
	"fmt"
	"log"
	"os"

	"audit-log/src/infrastructure/database/mysql"
	"audit-log/src/infrastructure/rabbitmq"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("Error loading .env file")
	}

	mysql.InitMySQLReadDB()
	mysql.InitMySQLWriteDB()

	if err := rabbitmq.InitRabbitMQ(); err != nil {
		fmt.Print("Error initializing RabbitMQ:", err)
	}

	app := echo.New()

	port := os.Getenv("PORT")
	if port == "" {
		port = "4000"
	}

	// Register routes
	rest.RegisterRoutes(app)

	log.Println("Connected to Application")
	app.Logger.Fatal(app.Start(":" + port))
}
