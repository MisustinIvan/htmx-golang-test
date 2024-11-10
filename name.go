package main

import "github.com/gofiber/fiber/v2"

func handleName(c *fiber.Ctx) error {
	name := c.FormValue("name", "unknown_name")
	err := templates.Render("name_sent", map[string]string{"Name": name}, c)
	if err != nil {
		return err
	}
	return c.SendStatus(fiber.StatusOK)
}
