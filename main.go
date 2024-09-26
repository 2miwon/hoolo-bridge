package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/swagger"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"

	"github.com/2miwon/hoolo-bridge/api"
	"github.com/2miwon/hoolo-bridge/db"
	_ "github.com/2miwon/hoolo-bridge/docs"
	"github.com/2miwon/hoolo-bridge/utils"
)

func setupDatabase() (*pgxpool.Pool, error) {
	runSchema := flag.Bool("deploy", false, "Run schema setup")
    flag.Parse()

    err := godotenv.Load()
    utils.CheckErr(err)

    dbURI := os.Getenv("DB_URI")
    config, err := pgxpool.ParseConfig(dbURI)
    utils.CheckErr(err)

    config.ConnConfig.DefaultQueryExecMode = pgx.QueryExecModeSimpleProtocol

    pool, err := pgxpool.NewWithConfig(context.Background(), config)
    utils.CheckErr(err)

	if *runSchema {
		schema, err := os.ReadFile("sqlc/schema.sql")
		utils.CheckErr(err)

		_, err = pool.Exec(context.Background(), string(schema))
		utils.CheckErr(err)
	}

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
// @host localhost:3000
// @BasePath /
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
		AllowMethods: "GET, POST, PUT, DELETE",
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
		return api.FetchMyInfo(c, q)
	})

	app.Post("/user/resign", func(c *fiber.Ctx) error {
		if !utils.ContextChecker(c) { return errors.New("CONTEXT IS NIL") }
		return api.Resign(c, q)
	})

	app.Get("/place/list", func(c *fiber.Ctx) error {
		if !utils.ContextChecker(c) { return errors.New("CONTEXT IS NIL") }
		return api.FetchRandomPlaceList(c, 5)
	})

	app.Get("/place/detail/:id", func(c *fiber.Ctx) error {
		if !utils.ContextChecker(c) { return errors.New("CONTEXT IS NIL") }
		return api.FetchPlaceDetail(c)
	})

	app.Get("/place/recent", func(c *fiber.Ctx) error {
		return api.FetchMostPlaceList(c, q)
	})

	app.Get("/place/search/:keyword", func(c *fiber.Ctx) error {
		if !utils.ContextChecker(c) { return errors.New("CONTEXT IS NIL") }
		return api.SearchPlace(c)
	})

	app.Get("/holog/relate/:id", func(c *fiber.Ctx) error {
		if !utils.ContextChecker(c) { return errors.New("CONTEXT IS NIL") }
		return api.FetchRelatePlaceList(c, q)
	})

	app.Post("holog/create", func(c *fiber.Ctx) error {
		if !utils.ContextChecker(c) { return errors.New("CONTEXT IS NIL") }
		return api.CreateHolog(c, q)
	})

	app.Post("/holog/delete", func(c *fiber.Ctx) error {
		if !utils.ContextChecker(c) { return errors.New("CONTEXT IS NIL") }
		return api.DeleteHolog(c, q)
	})
	
	app.Post("/holog/hide", func(c *fiber.Ctx) error {
		if !utils.ContextChecker(c) { return errors.New("CONTEXT IS NIL") }
		return api.HideHolog(c, q)
	})

	app.Get("/holog/:user_id", func(c *fiber.Ctx) error {
		if !utils.ContextChecker(c) { return errors.New("CONTEXT IS NIL") }
		return api.ListHologsByUserID(c, q)
	})

	app.Get("/schedule/:user_id", func(c *fiber.Ctx) error {
		if !utils.ContextChecker(c) { return errors.New("CONTEXT IS NIL") }
		return api.GetSchedule(c, q)
	})

	app.Post("/schedule/create", func(c *fiber.Ctx) error {
		if !utils.ContextChecker(c) { return errors.New("CONTEXT IS NIL") }
		return api.CreateSchedule(c, q)
	})

	app.Post("/schedule/update", func(c *fiber.Ctx) error {
		if !utils.ContextChecker(c) { return errors.New("CONTEXT IS NIL") }
		return api.UpdateSchedule(c, q)
	})

	app.Get("/schedule/detail/:schedule_id", func(c *fiber.Ctx) error {
		if !utils.ContextChecker(c) { return errors.New("CONTEXT IS NIL") }
		return api.GetScheduleDetail(c, q)
	})

	app.Post("/schedule/detail/place", func(c *fiber.Ctx) error {
		if !utils.ContextChecker(c) { return errors.New("CONTEXT IS NIL") }
		return api.GetScheduleDetailByPlaceID(c, q)
	})

	app.Post("/schedule/detail/create", func(c *fiber.Ctx) error {
		if !utils.ContextChecker(c) { return errors.New("CONTEXT IS NIL") }
		return api.CreateScheduleDetail(c, q)
	})

	app.Post("/schedule/detail/delete", func(c *fiber.Ctx) error {
		if !utils.ContextChecker(c) { return errors.New("CONTEXT IS NIL") }
		return api.DeleteScheduleDetail(c, q)
	})

	app.Post("/bookmark/set", func(c *fiber.Ctx) error {
		if !utils.ContextChecker(c) { return errors.New("CONTEXT IS NIL") }
		return api.SetBookmark(c, q)
	})

	app.Post("/bookmark/unset", func(c *fiber.Ctx) error {
		if !utils.ContextChecker(c) { return errors.New("CONTEXT IS NIL") }
		return api.UnsetBookmark(c, q)
	})

	app.Post("/bookmark/list", func(c *fiber.Ctx) error {
		if !utils.ContextChecker(c) { return errors.New("CONTEXT IS NIL") }
		return api.ListBookmark(c, q)
	})

	// app.Get("/bookmark/list", func(c *fiber.Ctx) error {
	// 	if !utils.ContextChecker(c) { return errors.New("CONTEXT IS NIL") }
	// 	return api.FetchBookmarkList(c, q)
	// }

	app.Get("/announce/list", func(c *fiber.Ctx) error {
		if !utils.ContextChecker(c) { return errors.New("CONTEXT IS NIL") }
		return api.ListAnnounces(c, q)
	})
	
	app.Post("/upload", api.UploadBucketSupabase)
	app.Get("/upload/s3", api.UploadS3)
	// app.Post("/upload/imgbb", api.ImgBBUpload)

	app.Listen(":3000")
}