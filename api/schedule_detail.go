package api

import (
	"context"
	"log"

	"github.com/2miwon/hoolo-bridge/db"
	"github.com/2miwon/hoolo-bridge/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// @Summary 스케줄 상세장소 조회
// @Description Get schedule by schedule id
// @Tags schedule
// @Accept json
// @Produce json
// @Param user_id query string true "Schedule ID"
// @Success 200 {object} []db.ScheduleDetail
// @Failure 400
// @Failure 500
// @Router /schedule/detail/ [get]
func GetScheduleDetail(c *fiber.Ctx, q *db.Queries) error {
	ctx := context.WithValue(context.Background(), "fiberCtx", c)

	scheduleIdStr := c.Params("schedule_id")

	scheduleId, err := uuid.Parse(scheduleIdStr)
    if err != nil {
		log.Printf("Invalid schedule ID format: %v", err)
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Invalid schedule ID format",
        })
    }	

	res, err := q.GetMyScheduleDetailsByScheduleId(ctx, scheduleId)
	if err != nil {
		log.Printf("Error fetching schedule: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to get schedule",
		})
	}

	return c.JSON(res)
}

// @Summary 스케줄 장소별 상세장소 조회
// @Description Get schedule by schedule id
// @Tags schedule
// @Accept json
// @Produce json
// @Param user_id query string true "Schedule ID"
// @Success 200 {object} []db.ScheduleDetail
// @Failure 400
// @Failure 500
// @Router /schedule/detail/place [post]
func GetScheduleDetailByPlaceID(c *fiber.Ctx, q *db.Queries) error {
	ctx := context.WithValue(context.Background(), "fiberCtx", c)

	var req	db.GetScheduleDetailByScheduleIdAndPlaceIdParams

	err := utils.ParseRequestBody(c, &req)
	if err != nil {
		log.Printf("Error parsing request body: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Failed to parse request",
		})
	}

	res, err := q.GetScheduleDetailByScheduleIdAndPlaceId(ctx, db.GetScheduleDetailByScheduleIdAndPlaceIdParams{
		ScheduleID: req.ScheduleID,
		PlaceID: req.PlaceID,
	})
	if err != nil {
		log.Printf("Error fetching schedule: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to get schedule",
		})
	}

	return c.JSON(res)
}

type CreateScheduleDetailRequest struct {
	ScheduleID string `json:"schedule_id"`
	PlaceID string `json:"place_id"`
	Title string `json:"title"`
}

// @Summary 스케줄 상세장소 생성
// @Description Create schedule detail
// @Tags schedule
// @Accept json
// @Produce json
// @Param CreateScheduleDetailRequest body CreateScheduleDetailRequest true "Create Schedule Detail Request"
// @Success 200 {object} db.ScheduleDetail
// @Failure 400
// @Failure 500
// @Router /schedule/detail/create [post]
func CreateScheduleDetail(c *fiber.Ctx, q *db.Queries) error {
	ctx := context.WithValue(context.Background(), "fiberCtx", c)

	var req CreateScheduleDetailRequest

	err := utils.ParseRequestBody(c, &req)
	if err != nil {
		log.Printf("Error parsing request body: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Failed to parse request",
		})
	}

	uuid, err := uuid.Parse(req.ScheduleID)
	if err != nil {
		log.Printf("Invalid schedule ID format: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid schedule ID format",
		})
	}

	var rst = db.CreateScheduleDetailParams{
		ScheduleID: uuid,
		PlaceID: req.PlaceID,
		Title: req.Title,
	}

	res, err := q.CreateScheduleDetail(ctx, rst)
	if err != nil {
		log.Printf("Error creating schedule detail: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create schedule detail",
		})
	}

	return c.JSON(res)
}

type DeleteScheduleDetailRequest struct {
	ScheduleID string `json:"schedule_id"`
	PlaceID string `json:"place_id"`
}

// @Summary 스케줄 상세장소 삭제
// @Description Delete schedule detail
// @Tags schedule
// @Accept json
// @Produce json
// @Param DeleteScheduleDetailRequest body DeleteScheduleDetailRequest true "Delete Schedule Detail Request"
// @Success 200 {object} db.ScheduleDetail
// @Failure 400
// @Failure 500
// @Router /schedule/detail/delete [post]
func DeleteScheduleDetail(c *fiber.Ctx, q *db.Queries) error {
	ctx := context.WithValue(context.Background(), "fiberCtx", c)

	var req DeleteScheduleDetailRequest

	err := utils.ParseRequestBody(c, &req)
	if err != nil {
		log.Printf("Error parsing request body: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Failed to parse request",
		})
	}

	uuid, err := uuid.Parse(req.ScheduleID)
	if err != nil {
		log.Printf("Invalid schedule ID format: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid schedule ID format",
		})
	}

	var rst = db.DeleteScheduleDetailParams{
		ScheduleID: uuid,
		PlaceID: req.PlaceID,
	}

	res, err := q.DeleteScheduleDetail(ctx, rst)
	if err != nil {
		log.Printf("Error deleting schedule detail: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to delete schedule detail",
		})
	}

	return c.JSON(res)
}
