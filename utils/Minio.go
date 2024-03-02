package utils

import (
	"bytes"
	"context"
	"net/url"
	"server/models"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/minio/minio-go/v7"
)

func UploadToMinio(data []byte, bucketName, objectName string) error {
	reader := bytes.NewReader(data)

	// загрузка файла в minio
	_, err := models.Minio.PutObject(context.TODO(), bucketName, objectName, reader, int64(reader.Len()), minio.PutObjectOptions{})
	if err != nil {
		return err
	}

	return nil
}
func ListObjects(c *gin.Context, minioClient *minio.Client, bucketName string) {
	objectCh := minioClient.ListObjects(context.Background(), bucketName, minio.ListObjectsOptions{
		Recursive: true,
	})

	var objectKeys []string

	for object := range objectCh {
		if object.Err != nil {
			c.JSON(500, gin.H{"error": "Failed to list objects"})
			return
		}
		objectKeys = append(objectKeys, object.Key)
	}

	c.JSON(200, gin.H{"objects": objectKeys})
}
func DownloadMinioFile(c *gin.Context) {
	reqParams := make(url.Values)
	bucketName := c.Query("bucketname")
	fileName := c.Query("filename")
	presignedURL, err := models.Minio.PresignedGetObject(c.Request.Context(), bucketName, fileName, 5*time.Minute, reqParams)
	if err != nil {
		c.JSON(500, "Failed to generate presigned URL")
		return
	}
	c.JSON(200, gin.H{"url": presignedURL})
}
