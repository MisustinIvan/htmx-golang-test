package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

var counter int = 69

func handleIncrement(c *fiber.Ctx) error {
	counter += 1
	return c.SendString(fmt.Sprintf("%d", counter))
}

func handleDecrement(c *fiber.Ctx) error {
	counter -= 1
	return c.SendString(fmt.Sprintf("%d", counter))
}
