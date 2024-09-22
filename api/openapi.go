package api

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"net/url"
	"os"
	"time"

	"github.com/2miwon/hoolo-bridge/db"
	"github.com/gofiber/fiber/v2"
)

type PlaceListResponse struct {
	Addr1 string `json:"addr1"`
	Addr2 string `json:"addr2"`
	Areacode string `json:"areacode"`
	Booktour string `json:"booktour"`
	Cat1 string `json:"cat1"`
	Cat2 string `json:"cat2"`
	Cat3 string `json:"cat3"`
	Contentid string `json:"contentid"`
	Contenttypeid string `json:"contenttypeid"`
	Createdtime string `json:"createdtime"`
	Firstimage string `json:"firstimage"`
	Firstimage2 string `json:"firstimage2"`
	CpyrhtDivCd string `json:"cpyrhtDivCd"`
	Mapx string `json:"mapx"`
	Mapy string `json:"mapy"`
	Mlevel string `json:"mlevel"`
	Modifiedtime string `json:"modifiedtime"`
	Sigungucode string `json:"sigungucode"`
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
func FetchRandomPlaceList(c *fiber.Ctx, q *db.Queries, n int) error {
	ctx := context.WithValue(context.Background(), "fiberCtx", c)
	base_url := os.Getenv("OPENAPI_LOCATION")
	
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	// result := make([]map[string]interface{}, 0, n)

	// for i := 0; i < n; i++ {
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
		// result = append(result, parsed)
	// }

	return c.JSON(parsed)
}

type PlaceDetailResponse struct {
	ContentId string `json:"contentid"`
	ContentTypeId string `json:"contenttypeid"`
	FirstImage string `json:"firstimage"`
	FirstImage2 string `json:"firstimage2"`
	CpyrhtDivCd string `json:"cpyrhtDivCd"`
	Addr1 string `json:"addr1"`
	Addr2 string `json:"addr2"`
	Zipcode string `json:"zipcode"`
	Mapx string `json:"mapx"`
	Mapy string `json:"mapy"`
	Mlevel string `json:"mlevel"`
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

// @Summary 장소 키워드로 검색하기 (장소명)
// @Description Search place
// @Tags place
// @Accept  json
// @Produce  json
// @Param keyword query string true "Keyword"
// @Success 200 {object} []PlaceListResponse
// @Failure 404
// @Failure 400
// @Router /place/search/{keyword} [get]
func SearchPlace(c *fiber.Ctx, q *db.Queries) error {
	ctx := context.WithValue(context.Background(), "fiberCtx", c)
	base_url := os.Getenv("OPENAPI_SEARCH")

	keyword := c.Query("keyword")
	encodedKeyword := url.QueryEscape(keyword)
	url := base_url + "&keyword=" + encodedKeyword

	resp, err := GetRequest(c, ctx, url)
	if err != nil {
		log.Printf("Error fetching place list: %v", err)
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Place list not exist",
		})
	}

	parsed := OpenApiParser(c, resp)

	return c.JSON(parsed)
}