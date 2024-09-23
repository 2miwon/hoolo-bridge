package api

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/gofiber/fiber/v2"
	storage_go "github.com/supabase-community/storage-go"
)

var s3Client *s3.Client

func InitS3Client() {
    s3Client = s3.New(s3.Options{
        EndpointResolver: s3.EndpointResolverFromURL(os.Getenv("BUCKET_ENDPOINT")),
        Region:           "ap-northeast-2",
        Credentials:      aws.NewCredentialsCache(credentials.NewStaticCredentialsProvider(os.Getenv("BUCKET_ACCESS_KEY"), os.Getenv("BUCKET_SECRET_ACCESS_KEY"), "")),
        UsePathStyle:     true,
    })
}

type UploadResponse struct {
    Message string `json:"message"`
    URL     string `json:"url"`
}

// @Summary 파일 업로드 with supabase
// @Description Upload file to Supabase
// @Tags file
// @Accept multipart/form-data
// @Produce json
// @Param file formData file true "File to upload"
// @Success 200 {object} UploadResponse
// @Failure 400
// @Failure 500
// @Router /upload [post]
func UploadBucketSupabase(c *fiber.Ctx) error {
    file, err := c.FormFile("file")
    if err != nil {
        log.Printf("Failed to get file from request: %v", err)
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Failed to get file from request",
        })
    }

    fileContent, err := file.Open()
    if err != nil {
        log.Printf("Failed to open file: %v", err)
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": "Failed to open file",
        })
    }
    defer fileContent.Close()

    storageClient := storage_go.NewClient(os.Getenv("BUCKET_ENDPOINT"), os.Getenv("BUCKET_SECRET_ACCESS_KEY"), nil)

    // Define the bucket and the path where you want to store the file
    bucketName := "hoolo_image"
    filePath := fmt.Sprintf("uploads/%s", file.Filename)

    // check 
    _, err = storageClient.CreateBucket("bucket-id", storage_go.BucketOptions{
        Public: true,
    })
    if err != nil {
        log.Printf("Failed to create bucket: %v", err)
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": "Failed to create bucket",
        })
    }

    // resp, err := storageClient.CreateSignedUploadUrl(bucketName, file.Filename)
    // if err != nil {
    //     log.Printf("Failed to get signed URL: %v", err)
    //     return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
    //         "error": "Failed to get signed URL",
    //     })
    // }

    // res, err := storageClient.UploadToSignedUrl(resp.Url, fileContent)
    // if err != nil {
    //     log.Printf("Failed to upload file to Supabase: %v", err)
    //     return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
    //         "error": "Failed to upload file to Supabase",
    //     })
    // }

    uploadResponse, err := storageClient.UploadFile(bucketName, bucketName + "/" + filePath, fileContent)
    if err != nil {
        log.Printf("Failed to upload file to Supabase: %v", err)
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": "Failed to upload file to Supabase",
        })
    }

    return c.JSON(uploadResponse)
}

// @Summary 파일 업로드 with S3
// @Description Upload file to Supabase
// @Tags file
// @Accept multipart/form-data
// @Produce json
// @Param file formData file true "File to upload"
// @Success 200 {object} UploadResponse
// @Failure 400
// @Failure 500
// @Router /upload/s3 [post]
func UploadBucket(c *fiber.Ctx) error {
    file, err := c.FormFile("file")
    if err != nil {
        log.Printf("Failed to get file from request: %v", err)
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Failed to get file from request",
        })
    }

    fileContent, err := file.Open()
    if err != nil {
        log.Printf("Failed to open file: %v", err)
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": "Failed to open file",
        })
    }
    defer fileContent.Close()

    bucketName := "hoolo_image"
    filePath := fmt.Sprintf("uploads/%s", file.Filename)

    _, err = s3Client.PutObject(context.TODO(), &s3.PutObjectInput{
        Bucket: aws.String(bucketName),
        Key:    aws.String(filePath),
        Body:   fileContent,
    })
    if err != nil {
        log.Printf("Failed to upload file to Supabase: %v", err)
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