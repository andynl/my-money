package main

import (
	"github.com/andynl/my-money/controller"
	"log"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"os"
)

func main() {
	app := fiber.New()

	app.Use(recover.New())

	transactionController := controller.TransactionController{}

	transactionController.Route(app)

	port := os.Getenv("PORT")
	log.Fatal(app.Listen(":" + port))
}
