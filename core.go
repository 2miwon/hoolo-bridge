package main

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v4/pgxpool"
)

// @Summary
// @Description Register a new user with email, username and password
// @Tags users
// @Accept  json
// @Produce  json
// @Param   email     body    string     true        "Email"
// @Param   username  body    string     true        "Username"
// @Param   password  body    string     true        "Password"
// @Success 200 {object} User
// @Failure 400 {object} string "User already exists"
// @Failure 500 {object} string "Internal server error"
// @Router /register [post]
func getMyInfo(c *fiber.Ctx, db *pgxpool.Pool) error {
	rows, err := db.Query(context.Background(), "SELECT id, name FROM users")
        if err != nil {
            return err
        }
        defer rows.Close()

        var users []map[string]interface{}
        for rows.Next() {
            var id int
            var name string
            err := rows.Scan(&id, &name)
            if err != nil {
                return err
            }
            users = append(users, map[string]interface{}{
                "id":   id,
                "name": name,
            })
        }

        return c.JSON(users)
}