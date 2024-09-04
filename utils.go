package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
		// log.Println(err)
	}
}

func jsonParser(c *fiber.Ctx) map[string]interface{} {
	var body map[string]interface{}
	err := c.BodyParser(&body)
	if err != nil {
		return nil
	}
	return body
}

func contextChecker(c *fiber.Ctx) bool {
	return c.Context() != nil
}
