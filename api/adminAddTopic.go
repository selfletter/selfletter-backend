package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"selfletter-backend/db"
	"strings"
)

type AdminAddTopicRequest struct {
	Key  string `json:"key" binding:"required"`
	Name string `json:"name" binding:"required"`
}

type AdminAddTopicResponse struct {
	Error string `json:"error"`
	Topic string `json:"topic"`
}

func AdminAddTopic(c *gin.Context) {
	var request AdminAddTopicRequest
	dbHandle := db.GetDatabaseHandle()

	err := c.BindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, AdminAddTopicResponse{
			Error: "bad json",
			Topic: "",
		})
		return
	}

	if dbHandle.First(&db.AdminKey{}, "key = ?", request.Key).RowsAffected == 0 {
		c.JSON(http.StatusForbidden, AdminAddTopicResponse{
			Error: "invalid admin key",
			Topic: "",
		})
		return
	}

	if strings.Contains(request.Name, ",") {
		c.JSON(http.StatusBadRequest, AdminAddTopicResponse{
			Error: "name can't contain \",\"",
			Topic: "",
		})
		return
	}

	if dbHandle.First(&db.Topic{}, "name = ?", request.Name).RowsAffected != 0 {
		c.JSON(http.StatusBadRequest, AdminAddTopicResponse{
			Error: "topic already exists",
			Topic: request.Name,
		})
		return
	}

	if err := dbHandle.Create(db.Topic{Name: request.Name}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, AdminAddTopicResponse{
			Error: "database error",
			Topic: "",
		})
		return
	}

	c.JSON(http.StatusOK, AdminAddTopicResponse{
		Error: "",
		Topic: request.Name,
	})
}
