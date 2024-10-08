package api

import (
	"context"
	"log"
	"time"

	"github.com/2miwon/hoolo-bridge/db"
	"github.com/2miwon/hoolo-bridge/utils"
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
// @Router /schedule/ [get]
func GetSchedule(c *fiber.Ctx, q *db.Queries) error {
	ctx := context.WithValue(context.Background(), "fiberCtx", c)

	userID := c.Params("user_id")

	res, err := q.GetScheduleByUserID(ctx, userID)
	if err != nil {
		log.Printf("Failed to get schedule: %v", err)
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
	if 	err := utils.ParseRequestBody(c, &req); err != nil {
		log.Printf("Failed to parse request: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Failed to parse request",
		})
	}

	loc, err := time.LoadLocation("Asia/Seoul")
    if err != nil {
        log.Printf("Failed to load location: %v", err)
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": "Failed to load location",
        })
    }

	res, err := q.CreateSchedule(ctx, db.CreateScheduleParams{
		UserID:    req.UserID,
		StartDate: req.StartDate.In(loc),
		EndDate:   req.EndDate.In(loc),
	})
	if err != nil {
		log.Printf("Failed to create schedule: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create schedule",
		})
	}

	return c.JSON(res)
}

// @Summary 스케줄 수정
// @Description Update schedule
// @Tags schedule
// @Accept json
// @Produce json
// @Param db.UpdateScheduleParams body db.UpdateScheduleParams true "Update Schedule Request"
// @Success 200 {object} db.UpdateScheduleRow
// @Failure 400
// @Failure 500
// @Router /schedule/update [post]
func UpdateSchedule(c *fiber.Ctx, q *db.Queries) error {
	ctx := context.WithValue(context.Background(), "fiberCtx", c)

	var req db.UpdateScheduleParams
	if 	err := utils.ParseRequestBody(c, &req); err != nil {
		log.Printf("Failed to parse request: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Failed to parse request",
		})
	}

	res, err := q.UpdateSchedule(ctx, db.UpdateScheduleParams{
		ID:		 	req.ID,
		StartDate:  req.StartDate,
		EndDate:    req.EndDate,
	})
	if err != nil {
		log.Printf("Failed to update schedule: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to update schedule",
		})
	}

	return c.JSON(res)
}