package api

import (
	"context"
	"log"

	"github.com/2miwon/hoolo-bridge/db"
	"github.com/2miwon/hoolo-bridge/utils"
	"github.com/gofiber/fiber/v2"
)

// @Summary 북마크 설정
// @Description Set bookmark by place id
// @Tags bookmark
// @Accept json
// @Produce json
// @Param db.SetBookmarkByHologIdParams body db.SetBookmarkByHologIdParams true "Set Bookmark Request"
// @Success 200 {object} db.SetBookmarkByHologIdRow
// @Failure 400
// @Failure 500
// @Router /bookmark/set [post]
func SetBookmark(c *fiber.Ctx, q *db.Queries) error {
	ctx := context.WithValue(context.Background(), "fiberCtx", c)

	var req db.SetBookmarkByHologIdParams

	err := utils.ParseRequestBody(c, &req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request",
		})
	}

	_, err = q.SetBookmarkByHologId(ctx, db.SetBookmarkByHologIdParams{
		HologID: req.HologID,
		UserID: req.UserID,
	})
	if err != nil {
		log.Printf("Error setting bookmark: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to set bookmark",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Bookmark set successfully",
	})
}

// @Summary 북마크 해제
// @Description Unset bookmark by place id
// @Tags bookmark
// @Accept json
// @Produce json
// @Param db.DeleteBookmarkByHologIdParams body db.DeleteBookmarkByHologIdParams true "Unset Bookmark Request"
// @Success 200 {object} db.DeleteBookmarkByHologIdRow
// @Failure 400
// @Failure 500
// @Router /bookmark/unset [post]
func UnsetBookmark(c *fiber.Ctx, q *db.Queries) error {
	ctx := context.WithValue(context.Background(), "fiberCtx", c)

	var req db.DeleteBookmarkByHologIdParams

	err := utils.ParseRequestBody(c, &req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request",
		})
	}

	_, err = q.DeleteBookmarkByHologId(ctx, db.DeleteBookmarkByHologIdParams{
		HologID: req.HologID,
		UserID: req.UserID,
	})
	if err != nil {
		log.Printf("Error unsetting bookmark: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to unset bookmark",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Bookmark unset successfully",
	})
}

// @Summary 북마크 리스트 가져오기
// @Description Get bookmark list by user id / place id
// @Tags bookmark
// @Accept json
// @Produce json
// @Param GetBookmarkByUserIDAndPlaceIDParams body db.GetBookmarkByUserIDAndPlaceIDParams true "Get Bookmark Request"
// @Success 200 {object} []db.GetBookmarkByUserIDAndPlaceIDRow
// @Failure 404
// @Failure 400
// @Router /bookmark/list [post]
func ListBookmark(c *fiber.Ctx, q *db.Queries) error {
	ctx := context.WithValue(context.Background(), "fiberCtx", c)

	var req db.GetBookmarkByUserIDAndPlaceIDParams

	err := utils.ParseRequestBody(c, &req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request",
		})
	}

	bookmarks, err := q.GetBookmarkByUserIDAndPlaceID(ctx, db.GetBookmarkByUserIDAndPlaceIDParams{
		PlaceID: req.PlaceID,
		UserID: req.UserID,
	})
	if err != nil {
		log.Printf("Error fetching bookmark list: %v", err)
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Bookmark list not exist",
		})
	}

	return c.JSON(bookmarks)
}

type CheckBookmarkRequest struct {
	Valid	bool	`json:"valid"`
}

// @Summary 북마크 전 일정 등록여부 확인
// @Description Check bookmark by user id / place id
// @Tags bookmark
// @Accept json
// @Produce json
// @Param GetBookmarkByUserIDAndHologIDParams body db.GetBookmarkByUserIDAndHologIDParams true "Check Bookmark Request"
// @Success 200 {object} CheckBookmarkRequest
// @Failure 404
// @Failure 400
// @Router /bookmark/check [post]
func CheckValidBookmark(c *fiber.Ctx, q *db.Queries) error {
	ctx := context.WithValue(context.Background(), "fiberCtx", c)

	var req db.GetBookmarkByUserIDAndHologIDParams

	err := utils.ParseRequestBody(c, &req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request",
		})
	}

	valid, err := q.GetBookmarkByUserIDAndHologID(ctx, db.GetBookmarkByUserIDAndHologIDParams{
		PlaceID: req.PlaceID,
		UserID: req.UserID,
	})
	if err != nil {
		log.Printf("Error setting bookmark: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to set bookmark",
		})
	}

	isValid := valid != 0

	// to return valid
	return c.JSON(fiber.Map{
		"valid": isValid,
	})
}
