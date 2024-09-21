package utils

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func CheckErr(err error) {
	if err != nil {
		log.Fatal(err)
		// log.Println(err)
	}
}

func JsonParser(c *fiber.Ctx) map[string]interface{} {
	var body map[string]interface{}
	err := c.BodyParser(&body)
	if err != nil {
		return nil
	}
	return body
}

func ContextChecker(c *fiber.Ctx) bool {
	log.Printf("context: %v", c.Context())
	return c.Context() != nil
}

func ParseRequestBody(c *fiber.Ctx, req interface{}) error {
	log.Printf("request body: %v", req)
    if err := c.BodyParser(req); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Invalid request body",
        })
    }
    return nil
}

