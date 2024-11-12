package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

var (
	templates *Templates
	store     *session.Store
)

func main() {
	app := fiber.New()
	templates = newTemplates()
	store = session.New()

	app.Get("/index.js", func(c *fiber.Ctx) error {
		return c.SendFile("./index.js")
	})

	app.Get("/", func(c *fiber.Ctx) error {
		sess, err := store.Get(c)
		if err != nil {
			return nil
		}
		defer sess.Save()

		counter := sess.Get("counter")
		if counter == nil {
			counter = 0
		}

		err = templates.Render("index", map[string]string{"Counter": fmt.Sprintf("%d", counter.(int))}, c)
		return err
	})

	app.Post("/increment", handleIncrement)
	app.Post("/decrement", handleDecrement)

	app.Post("/name_input", handleName)

	app.Get("/get_temperature", handleTemperature)

	app.Listen(":6969")
}
