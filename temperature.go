package main

import (
	"fmt"
	"math/rand"

	"github.com/gofiber/fiber/v2"
)

func handleTemperature(c *fiber.Ctx) error {
	temp := rand.Int()
	return c.SendString(fmt.Sprintf("%d", temp))
}
