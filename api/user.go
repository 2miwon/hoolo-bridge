package api

import (
	"context"

	"github.com/2miwon/hoolo-bridge/db"
	"github.com/2miwon/hoolo-bridge/utils"
	"github.com/gofiber/fiber/v2"
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
func GetMyInfo(c *fiber.Ctx, q *db.Queries) error {
    ctx := context.WithValue(context.Background(), "fiberCtx", c)

    type MyInfoRequest struct {
            Email    string `json:"email"`
        }

        var req MyInfoRequest

        err := utils.ParseRequestBody(c, &req)
        utils.CheckErr(err)

        user, err := q.GetUserByID(ctx, req.Email)
        utils.CheckErr(err)

        return c.JSON(user)
}

func SignUp(c *fiber.Ctx, q *db.Queries) error {
    ctx := context.WithValue(context.Background(), "fiberCtx", c)

    type SignUpRequest struct {
            Id    string `json:"id"`
            Password string `json:"password"`
            Username string `json:"username"`
        }

        var req SignUpRequest

        err := utils.ParseRequestBody(c, &req)
        utils.CheckErr(err)

        user, err := q.CreateUser(ctx, db.CreateUserParams{
            Email:		req.Id,
            Password:	req.Password,
            Username:	req.Username,
        })
        utils.CheckErr(err)

        return c.JSON(user)
}

func Login(c *fiber.Ctx, q *db.Queries) error {
    ctx := context.WithValue(context.Background(), "fiberCtx", c)

	type LoginRequest struct {
            Id    string `json:"id"`
            Password string `json:"password"`
        }

		var req LoginRequest

		err := utils.ParseRequestBody(c, &req)
		utils.CheckErr(err)

		user, err := q.GetUserByEmailAndPassword(ctx, db.GetUserByEmailAndPasswordParams{
			Email:		req.Id,
			Password:	req.Password,
		})
		utils.CheckErr(err)

		return c.JSON(user)
}