package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"selfletter-backend/db"
	"strings"
)

type adminAddTopicRequest struct {
	Key  string `json:"key" binding:"required"`
	Name string `json:"name" binding:"required"`
}

type adminAddTopicResponse struct {
	Error string `json:"error"`
	Topic string `json:"topic"`
}

func AdminAddTopic(c *gin.Context) {
	var request adminAddTopicRequest
	dbHandle := db.GetDatabaseHandle()

	err := c.BindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, adminAddTopicResponse{
			Error: "bad json",
			Topic: "",
		})
		return
	}

	if dbHandle.First(&db.AdminKey{}, "key = ?", request.Key).RowsAffected == 0 {
		c.JSON(http.StatusForbidden, adminAddTopicResponse{
			Error: "invalid admin key",
			Topic: "",
		})
		return
	}

	if strings.Contains(request.Name, ",") {
		c.JSON(http.StatusBadRequest, adminAddTopicResponse{
			Error: "name can't contain \",\"",
			Topic: "",
		})
		return
	}

	if dbHandle.First(&db.Topic{}, "name = ?", request.Name).RowsAffected != 0 {
		c.JSON(http.StatusBadRequest, adminAddTopicResponse{
			Error: "topic already exists",
			Topic: request.Name,
		})
		return
	}

	if err := dbHandle.Create(db.Topic{Name: request.Name}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, adminAddTopicResponse{
			Error: "database error",
			Topic: "",
		})
		return
	}

	c.JSON(http.StatusOK, adminAddTopicResponse{
		Error: "",
		Topic: request.Name,
	})
}
