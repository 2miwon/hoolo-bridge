package api

import (
	"context"

	"github.com/2miwon/hoolo-bridge/db"
	"github.com/gofiber/fiber/v2"
)

// @Summary 스케줄 조회
// @Description Get schedule by user id
// @Tags schedule
// @Accept json
// @Produce json
// @Param user_id query string true "User ID"
// @Success 200 {object} db.GetScheduleByUserIDRow
// @Failure 400
// @Failure 500
// @Router /schedule/:user_id [get]
func GetSchedule(c *fiber.Ctx, q *db.Queries) error {
	ctx := context.WithValue(context.Background(), "fiberCtx", c)

	userID := c.Params("user_id")

	res, err := q.GetScheduleByUserID(ctx, userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to get schedule",
		})
	}

	return c.JSON(res)
}

// @Summary 스케줄 생성
// @Description Create schedule
// @Tags schedule
// @Accept json
// @Produce json
// @Param db.CreateScheduleParams body db.CreateScheduleParams true "Create Schedule Request"
// @Success 200 {object} db.CreateScheduleRow
// @Failure 400
// @Failure 500
// @Router /schedule/create [post]
func CreateSchedule(c *fiber.Ctx, q *db.Queries) error {
	ctx := context.WithValue(context.Background(), "fiberCtx", c)

	var req db.CreateScheduleParams
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Failed to parse request",
		})
	}

	res, err := q.CreateSchedule(ctx, req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create schedule",
		})
	}

	return c.JSON(res)
}