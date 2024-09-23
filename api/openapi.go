package api

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
)

type PlaceListResponse struct {
	Addr1 string `json:"addr1"`
	Contentid string `json:"contentid"`
	Firstimage string `json:"firstimage"`
	Firstimage2 string `json:"firstimage2"`
	Mapx string `json:"mapx"`
	Mapy string `json:"mapy"`
	Tel string `json:"tel"`
	Title string `json:"title"`	
}

// @Summary 장소 리스트 가져오기(랜덤)
// @Description Get place list randomly
// @Tags place
// @Accept  json
// @Produce  json
// @Success 200 {object} []PlaceListResponse
// @Failure 404
// @Failure 400
// @Router /place/list [get]
func FetchRandomPlaceList(c *fiber.Ctx, n int) error {
	ctx := context.WithValue(context.Background(), "fiberCtx", c)
	base_url := os.Getenv("OPENAPI_LOCATION")
	
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	result := make([]map[string]interface{}, 0, n)

	for len(result) < n {
		pageNo := r.Intn(246) + 1
		url := base_url + "&pageNo=" + fmt.Sprintf("%d", pageNo)

		resp, err := GetRequest(c, ctx, url)
		if err != nil {
			log.Printf("Error fetching place list: %v", err)
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Place list not exist",
			})
		}

		parsed := OpenApiParser(c, resp)
		if parsed == nil {
			log.Printf("Error parsing data: %v", parsed)
			return c.JSON([]PlaceRecentResponse{})
		}
		
		for _, item := range parsed {
    	    // data를 map[string]interface{} 타입으로 변환
    	    data, ok := item.(map[string]interface{})
    	    if !ok {
    	        log.Printf("Error parsing data: %v", item)
    	        continue
    	    }

    	    if photo, ok := data["firstimage"]; ok && photo != "" {
    	        result = append(result, data)
    	        if len(result) >= n {
    	            break
    	        }
    	    }
    	}
	}

	return c.JSON(result)
}

type PlaceDetailResponse struct {
	ContentId string `json:"contentid"`
	FirstImage string `json:"firstimage"`
	FirstImage2 string `json:"firstimage2"`
	Addr1 string `json:"addr1"`
	Zipcode string `json:"zipcode"`
	Mapx string `json:"mapx"`
	Mapy string `json:"mapy"`
	Overview string `json:"overview"`
}

// @Summary 장소 상세정보 가져오기
// @Description Get place detail
// @Tags place
// @Accept  json
// @Produce  json
// @Param id path string true "Place ID (contentId)"
// @Success 200 {object} PlaceDetailResponse
// @Failure 404
// @Failure 400
// @Router /place/detail/ [get]
func FetchPlaceDetail(c *fiber.Ctx) error {
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
	if result == nil {
		log.Printf("Error parsing data: %v", result)
		return c.JSON([]PlaceRecentResponse{})
	}

	log.Printf("place detail: %v", result)
	
	return c.JSON(result[0])
}

// @Summary 장소 키워드로 검색하기 (장소명)
// @Description Search place
// @Tags place
// @Accept  json
// @Produce  json
// @Param keyword query string true "Keyword"
// @Success 200 {object} []PlaceListResponse
// @Failure 404
// @Failure 400
// @Router /place/search/ [get]
func SearchPlace(c *fiber.Ctx) error {
	ctx := context.WithValue(context.Background(), "fiberCtx", c)
	base_url := os.Getenv("OPENAPI_SEARCH")

	keyword := c.Params("keyword")
	// encodedKeyword := url.QueryEscape(keyword)
	url := base_url + "&numOfRows=16&pageNo=1&keyword=" + keyword

	resp, err := GetRequest(c, ctx, url)
	if err != nil {
		log.Printf("Error fetching place list: %v", err)
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Place list not exist",
		})
	}

	parsed := OpenApiParser(c, resp)
	if parsed == nil {
		log.Printf("Error parsing data: %v", parsed)
		return c.JSON([]PlaceRecentResponse{})
	}

	return c.JSON(parsed)
}