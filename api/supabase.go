package api

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/gofiber/fiber/v2"
)

var s3Client *s3.Client

func InitS3Client() {
    s3Client = s3.New(s3.Options{
        EndpointResolver: s3.EndpointResolverFromURL("https://project_ref.supabase.co/storage/v1/s3"),
        Region:           "project_region",
        Credentials:      aws.NewCredentialsCache(credentials.NewStaticCredentialsProvider("your_access_key_id", "your_secret_access_key", "")),
        UsePathStyle:     true,
    })
}

type UploadResponse struct {
    Message string `json:"message"`
    URL     string `json:"url"`
}

// @Summary 파일 업로드
// @Description Upload file to Supabase
// @Tags file
// @Accept multipart/form-data
// @Produce json
// @Param file formData file true "File to upload"
// @Success 200 {object} UploadResponse
// @Failure 400
// @Failure 500
// @Router /upload [post]
func UploadBucket(c *fiber.Ctx) error {
    file, err := c.FormFile("file")
    if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Failed to get file from request",
        })
    }

    fileContent, err := file.Open()
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": "Failed to open file",
        })
    }
    defer fileContent.Close()

    bucketName := "your-bucket-name"
    filePath := fmt.Sprintf("uploads/%s", file.Filename)

    _, err = s3Client.PutObject(context.TODO(), &s3.PutObjectInput{
        Bucket: aws.String(bucketName),
        Key:    aws.String(filePath),
        Body:   fileContent,
    })
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": "Failed to upload file to Supabase",
        })
    }

    fileURL := fmt.Sprintf("https://project_ref.supabase.co/storage/v1/object/public/%s/%s", bucketName, filePath)

    return c.JSON(UploadResponse{
        Message: "File uploaded successfully",
        URL:     fileURL,
    })
}