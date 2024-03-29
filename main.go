package main

import (
	"log"
	

	"github.com/DoomGuy1818/gofiber-test/database"
	"github.com/gofiber/fiber/v2"
	
	
)

func welcome(c *fiber.Ctx) error {
	return c.SendString("welcome to my awesome API")
}


func main() {
	database.ConnectDb()
	app := fiber.New()

	app.Get("/api", welcome)

	log.Fatal(app.Listen(":3000"))
}
