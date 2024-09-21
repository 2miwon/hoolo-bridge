package api

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/2miwon/hoolo-bridge/db"
	"github.com/gofiber/fiber/v2"
)

// @Summary 장소 리스트 가져오기(랜덤)
// @Description Get place list randomly
// @Tags openapi
// @Accept  json
// @Produce  json
// @Success 200 {object} map[string]interface{}
// @Failure 404
// @Failure 400
// @Router /place/list [get]
func FetchRandomPlaceList(c *fiber.Ctx, q *db.Queries, n int) error {
	ctx := context.WithValue(context.Background(), "fiberCtx", c)
	base_url := os.Getenv("OPENAPI_LOCATION")
	
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	result := make([]map[string]interface{}, 0, n)

	for i := 0; i < n; i++ {
		pageNo := r.Intn(2477) + 1

		url := base_url + "&numOfRows=1&pageNo=" + fmt.Sprintf("%d", pageNo)

		resp, err := GetRequest(c, ctx, url)
		if err != nil {
			log.Printf("Error fetching place list: %v", err)
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Place list not exist",
			})
		}

		parsed := OpenApiParser(c, resp)
		result = append(result, parsed)
	}

	return c.JSON(result)
}

// @Summary 장소 상세정보 가져오기
// @Description Get place detail
// @Tags openapi
// @Accept  json
// @Produce  json
// @Param id path string true "Place ID"
// @Success 200 {object} map[string]interface{}
// @Failure 404
// @Failure 400
// @Router /place/detail/{id} [get]
func FetchPlaceDetail(c *fiber.Ctx, q *db.Queries) error {
	ctx := context.WithValue(context.Background(), "fiberCtx", c)
	base_url := os.Getenv("OPENAPI_COMMON")

	id := c.Params("id")
	url := base_url + "&contentId=" + id

	resp, err := GetRequest(c, ctx, url)
	if err != nil {
		log.Printf("Error fetching place detail: %v", err)
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Place detail not exist",
		})
	}

	result := OpenApiParser(c, resp)

	log.Printf("place detail: %v", result)
	
	return c.JSON(result)
}