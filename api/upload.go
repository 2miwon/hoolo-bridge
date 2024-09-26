package api

import (
	"fmt"
	"log"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)
type UploadRequest struct {
	File string `json:"file"`
}

// @Summary 파일 업로드 with imgbb
// @Description Upload file to Supabase
// @Tags file
// @Accept multipart/form-data
// @Produce json
// @Param file formData file true "File to upload"
// @Success 200 {object} UploadResponse
// @Failure 400
// @Failure 500
// @Router /upload/imgbb [post]
// func ImgBBUpload(c *fiber.Ctx) error {
// 	apiKey := os.Getenv("IMGBB_API_KEY")
	
// 	file, err := c.FormFile("file")
//     if err != nil {
//         log.Printf("Failed to get file from request: %v", err)
//         return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
//             "error": "Failed to get file from request",
//         })
//     }
	
//     fileContent, err := file.Open()
//     if err != nil {
//         log.Printf("Failed to open file: %v", err)
//         return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
//             "error": "Failed to open file",
//         })
//     }
//     defer fileContent.Close()

// 	fileBytes, err := ioutil.ReadAll(fileContent)
//     if err != nil {
//         log.Printf("Failed to read file content: %v", err)
//         return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
//             "error": "Failed to read file content",
//         })
//     }

//     // Create a new buffer and multipart writer
//     requestBody := &bytes.Buffer{}
//     writer := multipart.NewWriter(requestBody)

// 	err = writer.WriteField("key", apiKey)
//     if err != nil {
//         log.Printf("Failed to add API key to form: %v", err)
//         return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
//             "error": "Failed to add API key to form",
//         })
//     }

//     // Create a form file field
//     part, err := writer.CreateFormFile("image", file.Filename)
//     if err != nil {
// 		log.Printf("Failed to create form file: %v", err)
//         return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
// 			"error": "Failed to create form file",
// 		})
//     }

// 	// Copy the file content to the form file field
//     _, err = io.Copy(part, bytes.NewReader(fileBytes))
//     if err != nil {
//         log.Printf("Failed to copy file content: %v", err)
//         return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
//             "error": "Failed to copy file content",
//         })
//     }
//     // Close the multipart writer to set the terminating boundary
//     err = writer.Close()
//     if err != nil {
// 		log.Printf("Failed to close writer: %v", err)
//         return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
// 			"error": "Failed to close writer",
// 		})
//     }

// 	log.Printf("Request body: %v", requestBody)

//     // Create a new POST request
//     url := "https://api.imgbb.com/1/upload"
//     req, err := http.NewRequest("POST", url, requestBody)
//     if err != nil {
// 		log.Printf("Failed to create request: %v", err)
//         return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
// 			"error": "Failed to create request",
// 		})
//     }

//     // Set the Content-Type header
//     req.Header.Set("Content-Type", writer.FormDataContentType())

//     // Send the request
//     client := &http.Client{}
//     resp, err := client.Do(req)
//     if err != nil {
// 		log.Printf("Failed to send request: %v", err)
//         return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
// 			"error": "Failed to send request",
// 		})
//     }
//     defer resp.Body.Close()

//     // Read the response body
//     respBody, err := io.ReadAll(resp.Body)
//     if err != nil {
// 		log.Printf("Failed to read response body: %v", err)
//         return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
// 			"error": "Failed to read response body",
// 		})
//     }

// 	var jsonResponse map[string]interface{}
// 	err = json.Unmarshal(respBody, &jsonResponse)
// 	if err != nil {
// 	    log.Printf("Failed to parse response body as JSON: %v", err)
// 	    return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
// 	        "error": "Failed to parse response body as JSON",
// 	    })
// 	}

//     // Return the response as a string
//     return c.JSON(jsonResponse)
// }

// @Summary presigned uri 받아오기 with S3
// @Description Get presigned uri from S3
// @Tags file
// @Accept json
// @Produce json
// @Success 200 {string} urlStr
// @Failure 400
// @Failure 500
// @Router /upload/s3 [get]
func UploadS3(c *fiber.Ctx) error {
	creds := credentials.NewSharedCredentials("", "heewon-test")

    sess, err := session.NewSession(&aws.Config{
        Region: aws.String("ap-northeast-2"),
		Credentials: creds,
	})

	if err != nil {
		log.Println("Failed to create session", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create session",
		})
	}

	svc := s3.New(sess)

	fileUUID := uuid.New().String()
    filePath := fmt.Sprintf("%s", fileUUID)

   	req, _ := svc.PutObjectRequest(&s3.PutObjectInput{
        Bucket: aws.String("hoolo"),
        Key:    aws.String(filePath),
    })

    urlStr, err := req.Presign(15 * time.Minute)
    if err != nil {
        log.Println("Failed to sign request", err)
    }

    return c.JSON(fiber.Map{
		"presigned_uri": urlStr,
	})
}

