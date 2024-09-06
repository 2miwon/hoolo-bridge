package main

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/swagger"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"

	"github.com/2miwon/hoolo-bridge/api"
	"github.com/2miwon/hoolo-bridge/db"
	_ "github.com/2miwon/hoolo-bridge/docs"
	"github.com/2miwon/hoolo-bridge/utils"
)

func setupDatabase() (*pgxpool.Pool, error) {
    err := godotenv.Load()
    utils.CheckErr(err)

    dbURI := os.Getenv("DB_URI")
    config, err := pgxpool.ParseConfig(dbURI)
    utils.CheckErr(err)

    pool, err := pgxpool.NewWithConfig(context.Background(), config)
    utils.CheckErr(err)

	schema, err := os.ReadFile("sqlc/schema.sql")
	utils.CheckErr(err)

	_, err = pool.Exec(context.Background(), string(schema))
	utils.CheckErr(err)

	fmt.Println("Pinged your deployment. You successfully connected to Postgre DB!")

    return pool, nil
}

// @title Hoolo API
// @version 0.1
// @description This is a Hoolo swagger docs for Fiber
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email zhengsfsf@gmail.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host 3.36.212.250:3000 // TODO: change to your server's IP
// @BasePath /docs
func main() {
	pool, err := setupDatabase()
	utils.CheckErr(err)
	defer pool.Close()

	q := db.New(pool)

	app := fiber.New()

	app.Static("/public", "./")

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendStatus(200)
	})

	app.Get("/docs/*", swagger.HandlerDefault)

	app.Get("/ping", func(c *fiber.Ctx) error {
		err = pool.Ping(context.Background())

		if err != nil {
			return c.SendString("Database connection failed")
		}

        return c.SendString("Database connection successful")
	})

	app.Get("/test/account/admin", func(c *fiber.Ctx) error {
		return c.SendString("Hello, Admin!")
	})

	app.Get("/test/account/user", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Post("/register", func(c *fiber.Ctx) error {
		if !utils.ContextChecker(c) { return errors.New("CONTEXT IS NIL") }
		return api.SignUp(c, q)
	})

	app.Post("/login", func(c *fiber.Ctx) error {
		if !utils.ContextChecker(c) { return errors.New("CONTEXT IS NIL") }
		return api.Login(c, q)
	})

	app.Post("/user/myinfo", func(c *fiber.Ctx) error {
		if !utils.ContextChecker(c) { return errors.New("CONTEXT IS NIL") }
		return api.GetMyInfo(c, q)
	})

	app.Listen(":3000")
}