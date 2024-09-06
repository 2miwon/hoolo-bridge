package api

import (
	"context"

	"github.com/2miwon/hoolo-bridge/db"
	"github.com/2miwon/hoolo-bridge/utils"
	"github.com/gofiber/fiber/v2"
)

// @Summary
// @Description Get user info by user ID
// @Tags users
// @Accept  json
// @Produce  json
// @Param   id    body    string     true        "User ID"
// @Success 200 {object} db.User
// @Failure 400 {object} string "User not found"
// @Failure 500 {object} string "Internal server error"
// @Router /myinfo [post]
func GetMyInfo(c *fiber.Ctx, q *db.Queries) error {
    ctx := context.WithValue(context.Background(), "fiberCtx", c)

    type MyInfoRequest struct {
            ID    string `json:"id"`
        }

        var req MyInfoRequest

        err := utils.ParseRequestBody(c, &req)
        utils.CheckErr(err)

        user, err := q.GetUserByID(ctx, req.ID)
        utils.CheckErr(err)

        return c.JSON(user)
}

// @Summary
// @Description Register a new user with email, username and password
// @Tags users
// @Accept  json
// @Produce  json
// @Param   email     body    string     true        "Email"
// @Param   username  body    string     true        "Username"
// @Param   password  body    string     true        "Password"
// @Success 200 {object} db.User
// @Failure 400 {object} string "User already exists"
// @Failure 500 {object} string "Internal server error"
// @Router /register [post]
func SignUp(c *fiber.Ctx, q *db.Queries) error {
    ctx := context.WithValue(context.Background(), "fiberCtx", c)

    type SignUpRequest struct {
            ID    string `json:"id"`
            Password string `json:"password"`
            Username string `json:"username"`
        }

        var req SignUpRequest

        err := utils.ParseRequestBody(c, &req)
        utils.CheckErr(err)

        user, err := q.CreateUser(ctx, db.CreateUserParams{
            ID:		req.ID,
            Password:	req.Password,
            Username:	req.Username,
        })
        utils.CheckErr(err)

        return c.JSON(user)
}

// @Summary
// @Description Register a new user with email, username and password
// @Tags users
// @Accept  json
// @Produce  json
// @Param   email     body    string     true        "Email"
// @Param   password  body    string     true        "Password"
// @Success 200 {object} db.User
// @Failure 400 {object} string "User already exists"
// @Failure 500 {object} string "Internal server error"
// @Router /login [post]
func Login(c *fiber.Ctx, q *db.Queries) error {
    ctx := context.WithValue(context.Background(), "fiberCtx", c)

	type LoginRequest struct {
            ID    string `json:"id"`
            Password string `json:"password"`
        }

		var req LoginRequest

		err := utils.ParseRequestBody(c, &req)
		utils.CheckErr(err)

		user, err := q.GetUserByEmailAndPassword(ctx, db.GetUserByEmailAndPasswordParams{
			ID:		req.ID,
			Password:	req.Password,
		})
		utils.CheckErr(err)

		return c.JSON(user)
}