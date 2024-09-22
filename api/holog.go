package api

import (
	"context"
	"encoding/json"
	"log"
	"os"

	"github.com/2miwon/hoolo-bridge/db"
	"github.com/gofiber/fiber/v2"
)

type PlaceRecentResponse struct {
    db.ListHologsMostByWeekRow
    PlaceDetail PlaceDetailResponse `json:"place_detail"`
}

// @Summary 최근 일주일간 가장 많이 홀로그가 생성된 장소 id 리스트 가져오기
// @Description Get most mentioned place list in the last week
// @Tags place
// @Accept json
// @Produce json
// @Success 200 {object} []PlaceRecentResponse
// @Failure 404
// @Failure 400
// @Router /place/recent [get]
func FetchMostPlaceList(c *fiber.Ctx, q *db.Queries) error {
	ctx := context.WithValue(context.Background(), "fiberCtx", c)
	
	list, err := q.ListHologsMostByWeek(ctx)
	if err != nil {
		log.Printf("Error fetching most place list: %v", err)
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Most place list not exist",
		})
	}

	rst := make([]PlaceRecentResponse, len(list))
	base_url := os.Getenv("OPENAPI_COMMON")

	for i := 0; i < len(list); i++ {
		content_id := list[i].PlaceID

		url := base_url + "&contentId=" + content_id

		resp, err := GetRequest(c, ctx, url)
		if err != nil {
			log.Printf("Error fetching place detail: %v", err)
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Place detail not exist",
			})
		}

    	placeDetail := OpenApiParser(c, resp)

		placeDetailJSON, err := json.Marshal(placeDetail)
    	if err != nil {
    	    log.Printf("Error marshalling place detail: %v", err)
    	    return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
    	        "error": "Failed to marshal place detail",
    	    })
    	}

    	var place PlaceDetailResponse
    	if err := json.Unmarshal(placeDetailJSON, &place); err != nil {
    	    log.Printf("Error unmarshalling place detail: %v", err)
    	    return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
    	        "error": "Failed to parse place detail",
    	    })
    	}

		rst[i] = PlaceRecentResponse{
			ListHologsMostByWeekRow: list[i],
			PlaceDetail: place,
		}
	}

	return c.JSON(rst)
}

// @Summary 특정 장소와 관련된 홀로그 리스트 가져오기 (최신순)
// @Description Get holog list related to specific place
// @Tags holog
// @Accept json
// @Produce json
// @Success 200 {object} []db.ListHologsByPlaceIdRow
// @Failure 404
// @Failure 400
// @Router /holog/relate/:id [get]
func FetchRelatePlaceList(c *fiber.Ctx, q *db.Queries) error {
	ctx := context.WithValue(context.Background(), "fiberCtx", c)

	id := c.Params("id")

	list, err := q.ListHologsByPlaceId(ctx, db.ListHologsByPlaceIdParams{
		PlaceID: id,
		Limit: 5,
	})

	if err != nil {
		log.Printf("Error fetching most place list: %v", err)
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Most place list not exist",
		})
	}

	return c.JSON(list)
}

// @Summary 홀로그 생성하기
// @Description Create a new holog
// @Tags holog
// @Accept json
// @Produce json
// @Param CreateHologParams body db.CreateHologParams true "CreateHolog Request"
// @Success 200 {object} db.CreateHologRow
// @Failure 400
// @Failure 500
// @Router /holog/create [post]
func CreateHolog(c *fiber.Ctx, q *db.Queries) error {
	ctx := context.WithValue(context.Background(), "fiberCtx", c)

	var req db.CreateHologParams

	err := c.BodyParser(&req)
	if err != nil {
		log.Printf("Error parsing request body: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request",
		})
	}

	holog, err := q.CreateHolog(ctx, req)
	if err != nil {
		log.Printf("Error creating holog: %v", err)
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Error creating holog",
		})
	}

	return c.JSON(holog)
}

// @Summary 특정 유저가 생성한 홀로그 리스트 가져오기
// @Description Get holog list
// @Tags holog
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} []db.ListHologsByUserIDRow
// @Failure 404
// @Failure 400
// @Router /holog/user/:id [get]
func ListHologsByUserID(c *fiber.Ctx, q *db.Queries) error {
	ctx := context.WithValue(context.Background(), "fiberCtx", c)

	user_id := c.Params("id")

	hologs, err := q.ListHologsByUserID(ctx, user_id)
	if err != nil {
		log.Printf("Error fetching hologs: %v", err)
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Error fetching hologs",
		})
	}

	return c.JSON(hologs)
}