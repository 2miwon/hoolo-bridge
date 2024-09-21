package api

import (
	"context"
	"database/sql"
	"log"

	"github.com/2miwon/hoolo-bridge/db"
	"github.com/2miwon/hoolo-bridge/utils"
	"github.com/gofiber/fiber/v2"
)

type MyInfoRequest struct {
    ID    string `json:"id"`
}

// @Summary 내정보 가져오기 (By ID)
// @Description Get user information by ID
// @Tags users
// @Accept  json
// @Produce  json
// @Param   MyInfoRequest  body    MyInfoRequest  true  "MyInfo Request"
// @Success 200 {object} db.GetUserByIDRow
// @Failure 404
// @Failure 400
// @Router /user/myinfo [post]
func FetchMyInfo(c *fiber.Ctx, q *db.Queries) error {
    ctx := context.WithValue(context.Background(), "fiberCtx", c)

    var req MyInfoRequest

    err := utils.ParseRequestBody(c, &req)
    if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Invalid request",
        })
    }

    user, err := q.GetUserByID(ctx, req.ID)
    if err != nil {
        log.Printf("Error fetching user: %v", err)
        return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
            "error": "User not exist",
        })
    }

    return c.JSON(user)
}

type SignUpRequest struct {
    ID    string `json:"id"`
    Password string `json:"password"`
    Username string `json:"username"`
}

// @Summary 회원가입
// @Description Register a new user with email, username and password
// @Tags users
// @Accept  json
// @Produce  json
// @Param   SignUpRequest  body    SignUpRequest  true  "SignUp Request"
// @Success 200 {object} db.CreateUserRow
// @Failure 400 
// @Failure 500
// @Router /register [post]
func SignUp(c *fiber.Ctx, q *db.Queries) error {
    ctx := context.WithValue(context.Background(), "fiberCtx", c)

    var req SignUpRequest

    err := utils.ParseRequestBody(c, &req)
    if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Invalid request",
        })
    }

    existingUser, err := q.GetUserByID(ctx, req.ID)
    if err == nil && existingUser.ID != "" {
        return c.Status(fiber.StatusConflict).JSON(fiber.Map{
            "error": "User ID already exists",
        })
    } else if err != nil && err != sql.ErrNoRows {
        log.Printf("Error checking user ID: %v", err)
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": "Failed to check user ID",
        })
    }

    user, err := q.CreateUser(ctx, db.CreateUserParams{
        ID:		req.ID,
        Password:	req.Password,
        Username:	req.Username,
    })
    if err != nil {
        log.Printf("Error creating user: %v", err)
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": "Failed to create user",
        })
    }

    return c.JSON(user)
}

type LoginRequest struct {
    ID    string `json:"id"`
    Password string `json:"password"`
}

// @Summary 로그인
// @Description Login with email and password
// @Tags users
// @Accept  json
// @Produce  json
// @Param   LoginRequest  body    LoginRequest  true  "Login Request"   
// @Success 200 {object} db.GetUserByEmailAndPasswordRow
// @Failure 400
// @Failure 500
// @Router /login [post]
func Login(c *fiber.Ctx, q *db.Queries) error {
    ctx := context.WithValue(context.Background(), "fiberCtx", c)

	var req LoginRequest

	err := utils.ParseRequestBody(c, &req)
	if err != nil {
        log.Printf("Error parsing request body: %v", err)

        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Invalid request",
        })
    }

	user, err := q.GetUserByEmailAndPassword(ctx, db.GetUserByEmailAndPasswordParams{
		ID:		req.ID,
		Password:	req.Password,
	})
	if err != nil {
        return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
            "error": "Invalid ID or Password",
        })
    }

	return c.JSON(user)
}