package usecase

import (
	"server/utils"
	"time"

	"github.com/gin-gonic/gin"
)

func Getstats(c *gin.Context) {

	orgName := c.Query("org_name")
	startStr := c.Query("start")
	endStr := c.Query("end")
	start, err := time.Parse(time.RFC3339, startStr)
	if err != nil {
		c.JSON(400, gin.H{"error": "Ошибка при парсинге времени start", "details": err.Error()})
		return
	}
	end, err := time.Parse(time.RFC3339, endStr)
	if err != nil {
		c.JSON(400, gin.H{"error": "Ошибка при парсинге времени start", "details": err.Error()})
		return
	}

	errorr := utils.Getstats(c, orgName, start, end)
	if errorr != nil {
		c.JSON(500, gin.H{"err": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "report generated"})
}
