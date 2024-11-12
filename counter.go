package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func handleIncrement(c *fiber.Ctx) error {
	sess, err := store.Get(c)
	if err != nil {
		return err
	}
	defer sess.Save()

	counter := sess.Get("counter")
	if counter == nil {
		counter = 0
	}

	sess.Set("counter", counter.(int)+1)

	return c.SendString(fmt.Sprintf("%d", sess.Get("counter").(int)))
}

func handleDecrement(c *fiber.Ctx) error {
	sess, err := store.Get(c)
	if err != nil {
		return err
	}
	defer sess.Save()

	counter := sess.Get("counter")
	if counter == nil {
		counter = 0
	}

	sess.Set("counter", counter.(int)-1)

	return c.SendString(fmt.Sprintf("%d", sess.Get("counter").(int)))
}
