package main

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/swagger"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/joho/godotenv"

	_ "github.com/2miwon/hoolo-bridge/docs"
)

// @title Hoolo API
// @version 0.1
// @description This is a swagger docs for Fiber
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email fiber@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host 3.36.212.250:3000 // TODO: change to your server's IP
// @BasePath /docs
func main() {
	err := godotenv.Load()
	checkErr(err)
	db_uri := os.Getenv("DB_URI")

	config, err := pgxpool.ParseConfig(db_uri)
    checkErr(err)

    db, err := pgxpool.ConnectConfig(context.Background(), config)
    checkErr(err)
    defer db.Close()

	fmt.Println("Pinged your deployment. You successfully connected to Postgre DB!")

	app := fiber.New()

	app.Static("/public", "./")

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendStatus(200)
	})

	app.Get("/ping", func(c *fiber.Ctx) error {
		var result int
        err := db.QueryRow(context.Background(), "SELECT 1").Scan(&result)
        if err != nil {
            return c.Status(fiber.StatusInternalServerError).SendString("Database connection failed")
        }

        return c.SendString("Database connection successful")
	})

	app.Get("/docs/*", swagger.HandlerDefault)

	// history, bookmark 정보들 다 있음
	app.Post("/user/my_info", func(c *fiber.Ctx) error {
		if !contextChecker(c) { return errors.New("CONTEXT IS NIL") }
		return getMyInfo(c, db)
	})

	app.Listen(":3000")
}