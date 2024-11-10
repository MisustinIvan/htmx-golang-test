package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

var templates *Templates

func main() {
	app := fiber.New()
	templates = newTemplates()

	app.Get("/index.js", func(c *fiber.Ctx) error {
		return c.SendFile("./index.js")
	})

	app.Get("/", func(c *fiber.Ctx) error {
		err := templates.Render("index", map[string]string{"Counter": fmt.Sprintf("%d", counter)}, c)
		log.Printf("viewed index.html")
		return err
	})

	app.Post("/increment", handleIncrement)
	app.Post("/decrement", handleDecrement)

	app.Post("/name_input", handleName)

	app.Get("/get_temperature", handleTemperature)

	app.Listen(":6969")
}
