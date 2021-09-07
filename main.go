package main

import (
	"github.com/andynl/my-money/controller"
	"github.com/andynl/my-money/model"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	//db, err := config.NewMysqlDatabase()

	app := fiber.New()
	app.Use(recover.New())
	app.Use(compress.New())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(model.WebResponse{
			Code: 200,
			Status: "OK",
			Data: fiber.Map{"message": "My Money - UangKu API"},
		})
	})

	transactionController := controller.TransactionController{}

	transactionController.Route(app)

	app.Get("*", func(c *fiber.Ctx) error {
		return c.JSON(model.WebResponse{
			Code: 404,
			Status: "Failed",
			Data: fiber.Map{"message": "Oops Error"},
		})
	})

	port := os.Getenv("PORT")
	log.Fatal(app.Listen(":" + port))
}
