package api

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func GetRequest(c *fiber.Ctx, ctx context.Context, url string) (map[string]interface{}, error) {
	log.Printf("request url: %v", url)
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
        log.Printf("Error creating request: %v", err)
		return nil, c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create request",
		})
    }

	req.Header.Set("Accept", "*/*")
    
    tr := &http.Transport{
        TLSClientConfig: &tls.Config{
            MinVersion: tls.VersionTLS10,   // 낮은 버전부터 시도
            MaxVersion: tls.VersionTLS13,   // 필요한 경우 제한
            InsecureSkipVerify: true,       // 테스트 목적 (실제 사용 시 false로 설정)
        },
        ForceAttemptHTTP2: false,           // 필요 시 HTTP/2 강제 비활성화
    }
    client := &http.Client{Transport: tr}
    resp, err := client.Do(req)
	if err != nil {
        log.Printf("Error making request: %v", err)
        return nil, c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": "Failed to make request",
        })
    }
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
        log.Printf("Unexpected status code: %v", resp.StatusCode)
        return nil, c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": "Unexpected status code",
        })
    }

	var result map[string]interface{}
	log.Printf("response: %v", resp.Body)
    if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
        log.Printf("Error decoding response: %v", err)
        return nil, c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": "Failed to decode response",
        })
    }

	return result, nil
}

func OpenApiParser(c *fiber.Ctx, decode_resp map[string]interface{}) []interface{} {
	response, ok := decode_resp["response"].(map[string]interface{})
    if !ok {
        log.Printf("Type assertion failed for 'response'")
        return nil
    }

    body, ok := response["body"].(map[string]interface{})
    if !ok {
        log.Printf("Type assertion failed for 'body'")
        return nil
    }

    items, ok := body["items"].(map[string]interface{})
    if !ok {
        log.Printf("Type assertion failed for 'items'")
        return nil
    }

    item, ok := items["item"].([]interface{})
    if !ok {
        log.Printf("Type assertion failed for 'item'")
        return nil
    }

	return item
}