package api

import (
	"context"
	"log"

	"github.com/2miwon/hoolo-bridge/db"
	"github.com/gofiber/fiber/v2"
)

// @Summary 공지사항 리스트 가져오기
// @Description Get announce list
// @Tags announce
// @Accept json
// @Produce json
// @Success 200 {object} []db.ListAnnouncesRow
// @Failure 404
// @Failure 400
// @Router /announce/list [get]
func ListAnnounces(c *fiber.Ctx, q *db.Queries) error {
	ctx := context.WithValue(context.Background(), "fiberCtx", c)

	list, err := q.ListAnnounces(ctx)
	if err != nil {
		log.Printf("Error fetching announce list: %v", err)
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Announce list not exist",
		})
	}

	return c.JSON(list)
}