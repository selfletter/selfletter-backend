package api

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"selfletter-backend/db"
)

type adminRemoveTopicRequest struct {
	Key   string `json:"key" binding:"required"`
	Topic string `json:"topic" binding:"required"`
}

type adminRemoveTopicResponse struct {
	Error string `json:"error"`
	Topic string `json:"topic"`
}

func AdminRemoveTopic(c *gin.Context) {
	var request adminRemoveTopicRequest
	dbHandle := db.GetDatabaseHandle()

	err := c.BindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, adminRemoveTopicResponse{
			Error: "bad json",
			Topic: "",
		})
		return
	}

	if dbHandle.First(&db.AdminKey{}, "key = ?", request.Key).RowsAffected == 0 {
		c.JSON(http.StatusForbidden, adminRemoveTopicResponse{
			Error: "invalid admin key",
			Topic: "",
		})
		return
	}

	if dbHandle.First(&db.Topic{}, "name = ?", request.Topic).RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, adminRemoveTopicResponse{
			Error: "topic doesn't exist",
			Topic: "",
		})
		return
	}

	err = dbHandle.Transaction(func(tx *gorm.DB) error {
		err = dbHandle.Where("name = ?", request.Topic).Delete(&db.Topic{}).Error
		if err != nil {
			return err
		}

		err = dbHandle.Where("topic = ?", request.Topic).Delete(&db.UsersTopicsRel{}).Error
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, adminRemoveTopicResponse{
			Error: "database error",
			Topic: "",
		})
		return
	}

	c.JSON(http.StatusOK, adminRemoveTopicResponse{
		Error: "",
		Topic: request.Topic,
	})
}
