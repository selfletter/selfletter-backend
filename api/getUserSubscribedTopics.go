package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"selfletter-backend/db"
)

type getUserSubscribedTopicsResponse struct {
	Error  string     `json:"error"`
	Topics []db.Topic `json:"topics"`
}

func GetUserSubscribedTopics(c *gin.Context) {
	dbHandle := db.GetDatabaseHandle()
	token := c.Query("token")

	if token == "" {
		c.JSON(http.StatusBadRequest, getUserSubscribedTopicsResponse{
			Error:  "empty token",
			Topics: nil,
		})
		return
	}
	if dbHandle.First(&db.User{}, "token = ?", token).RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, getUserSubscribedTopicsResponse{
			Error:  "invalid token",
			Topics: nil,
		})
		return
	}

	var user db.User
	err := dbHandle.Where("token = ?", token).Find(&user).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, getUserSubscribedTopicsResponse{
			Error:  "database error",
			Topics: nil,
		})
		return
	}

	var userTopicsRelations []db.UsersTopicsRel
	var topics []db.Topic

	err = dbHandle.Where("email = ?", user.Email).Find(&userTopicsRelations).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, getUserSubscribedTopicsResponse{
			Error:  "database error",
			Topics: nil,
		})
		return
	}

	for _, userTopicRel := range userTopicsRelations {
		topics = append(topics, db.Topic{Name: userTopicRel.Topic})
	}

	c.JSON(http.StatusOK, getUserSubscribedTopicsResponse{
		Error:  "",
		Topics: topics,
	})
}
