package utils

import (
	"fmt"
	"os"
	"server/models"
	"time"

	"github.com/gin-gonic/gin"
)

func Getstats(c *gin.Context, orgname string, start time.Time, end time.Time) error {
	var stats []models.Stats

	err := models.DB.Where("org_name = ? AND timestamp >= ? AND timestamp <= ?", orgname, start, end).Find(&stats).Error
	if err != nil {

		return err
	}
	if len(stats) == 0 {

		return err
	}

	csvData, err := ConvertToCSV(&stats)
	if err != nil {
		c.JSON(500, gin.H{"err": err.Error()})
		return err
	}

	nameByte := GenerateRandomBytes()
	time := time.Now()
	nameString := fmt.Sprintf("report(%x)_%s.csv", nameByte, time.Format("2006-01-02_15-04-05"))
	bucketName := os.Getenv("MINIO_BUCKET")

	err = UploadToMinio(csvData, bucketName, nameString)

	if err != nil {
		c.JSON(500, gin.H{"err": err.Error()})
		return err
	}
	var generatedStats models.GeneratedStats
	id, err := GetUserId(c)
	if err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return err
	}

	generatedStats.UserID = id
	generatedStats.Bucket = bucketName
	generatedStats.Name = nameString

	if err := models.DB.Create(&generatedStats).Error; err != nil {
		c.JSON(500, gin.H{"status": "error", "message": "error with creating bd }values"})
		return err
	}

	history := models.Histories{
		UserID:      id,
		Orgname:     orgname,
		CurrentTime: time,
		Start:       start,
		End:         end,
	}

	if err := models.DB.Create(&history).Error; err != nil {
		c.JSON(500, gin.H{"status": "error", "message": "error with creating bd }values"})
		return err
	}
	return nil
}
