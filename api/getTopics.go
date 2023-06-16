package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"selfletter-backend/db"
)

type GetTopicsResponse struct {
	Error  string     `json:"error"`
	Topics []db.Topic `json:"topics"`
}

func GetTopics(c *gin.Context) {
	dbHandle := db.GetDatabaseHandle()
	var topics []db.Topic
	if err := dbHandle.Find(&topics).Error; err != nil {
		c.JSON(http.StatusInternalServerError, GetTopicsResponse{
			Error:  "database error",
			Topics: nil,
		})
		return
	}

	c.JSON(http.StatusOK, GetTopicsResponse{
		Error:  "",
		Topics: topics,
	})
}
