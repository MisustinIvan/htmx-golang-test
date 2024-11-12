package main

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/storage/sqlite3/v2"
)

var (
	templates *Templates
	store     *session.Store
)

func main() {
	app := fiber.New()
	templates = newTemplates()
	storage := sqlite3.New(sqlite3.Config{
		Database:        "./db.db",
		Table:           "fiber_storage",
		Reset:           false,
		GCInterval:      10 * time.Second,
		MaxIdleConns:    100,
		MaxOpenConns:    100,
		ConnMaxLifetime: 1 * time.Second,
	})
	store = session.New(session.Config{
		Storage: storage,
	})

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
