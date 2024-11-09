package main

import (
	"fmt"
	"log"
	"math/rand"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	templates := newTemplates()

	counter := 69

	app.Get("/index.js", func(c *fiber.Ctx) error {
		return c.SendFile("./index.js")
	})

	app.Get("/style.css", func(c *fiber.Ctx) error {
		return c.SendFile("./style.css")
	})

	app.Get("/", func(c *fiber.Ctx) error {
		err := templates.Render("index", map[string]string{"Counter": fmt.Sprintf("%d", counter)}, c)
		log.Printf("viewed index.html")
		return err
	})

	app.Post("/increment", func(c *fiber.Ctx) error {
		counter += 1
		err := c.SendString(fmt.Sprintf("%d", counter))
		log.Printf("incremented")
		return err
	})

	app.Post("/decrement", func(c *fiber.Ctx) error {
		counter -= 1
		err := c.SendString(fmt.Sprintf("%d", counter))
		log.Printf("decremented")
		return err
	})

	app.Post("/name_input", func(c *fiber.Ctx) error {
		name := c.FormValue("name", "unknown_name")
		log.Printf("name: %s", name)
		err := templates.Render("name_sent", map[string]string{"Name": name}, c)
		if err != nil {
			return err
		}
		return c.SendStatus(fiber.StatusOK)
	})

	app.Get("/get_temperature", func(c *fiber.Ctx) error {
		temp := rand.Int()
		return c.SendString(fmt.Sprintf("%d", temp))
	})

	app.Listen(":6969")
}
