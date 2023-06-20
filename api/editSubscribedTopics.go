package api

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"selfletter-backend/db"
	"strings"
)

type editSubscribedTopicsResponse struct {
	Error string `json:"error"`
}

func EditSubscribedTopics(c *gin.Context) {
	dbHandle := db.GetDatabaseHandle()
	token := c.Query("token")
	topics := c.Query("topics")
	topicsSlice := strings.Split(topics, ",")

	if topics == "" {
		c.JSON(http.StatusBadRequest, editSubscribedTopicsResponse{
			Error: "no topics chosen",
		})
		return
	}

	var user db.User
	if dbHandle.First(&user, "token = ?", token).RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, editSubscribedTopicsResponse{
			Error: "there is no such user",
		})
		return
	}
	for _, topic := range topicsSlice {
		if dbHandle.First(&db.Topic{}, "name = ?", topic).RowsAffected == 0 {
			c.JSON(http.StatusBadRequest, editSubscribedTopicsResponse{
				Error: "there is no such topic: " + topic,
			})
			return
		}
	}

	err := dbHandle.Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("email = ?", user.Email).Delete(&db.UsersTopicsRel{}).Error; err != nil {
			return err
		}

		for _, topic := range topicsSlice {
			if err := tx.Create(&db.UsersTopicsRel{
				Email: user.Email,
				Topic: topic,
			}).Error; err != nil {
				return err
			}
		}

		return nil
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, editSubscribedTopicsResponse{Error: "database error"})
		return
	}

	c.JSON(http.StatusOK, editSubscribedTopicsResponse{Error: ""})
}
