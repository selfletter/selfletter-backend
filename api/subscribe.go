package api

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"net/mail"
	"selfletter-backend/db"
	"selfletter-backend/helpers"
	"strings"
)

type SubscribeResponse struct {
	Error string `json:"error"`
	Token string `json:"token"`
}

func Subscribe(c *gin.Context) {
	dbHandle := db.GetDatabaseHandle()
	token := helpers.GenerateSecureToken()
	for i := 0; i < 10; i++ {
		if dbHandle.First(&db.User{}, "token = ?", token).RowsAffected != 0 {
			token = helpers.GenerateSecureToken()
		} else {
			break
		}
		if i == 9 {
			c.JSON(http.StatusInternalServerError, SubscribeResponse{
				Error: "too many collisions",
				Token: "",
			})
			return
		}

	}
	topics := c.Query("topics")
	topicsSlice := strings.Split(topics, ",")
	email := c.Query("email")

	if topics == "" {
		c.JSON(http.StatusBadRequest, SubscribeResponse{
			Error: "no topics chosen",
			Token: "",
		})
		return
	}
	if email == "" {
		c.JSON(http.StatusBadRequest, SubscribeResponse{
			Error: "email is empty",
			Token: "",
		})
		return
	}
	if dbHandle.First(&db.User{}, "email = ?", email).RowsAffected != 0 {
		c.JSON(http.StatusBadRequest, SubscribeResponse{
			Error: "user already exists",
			Token: "",
		})
		return
	}
	for _, topic := range topicsSlice {
		if dbHandle.First(&db.Topic{}, "name = ?", topic).RowsAffected == 0 {
			c.JSON(http.StatusBadRequest, SubscribeResponse{
				Error: "there is no such topic: " + topic,
				Token: "",
			})
			return
		}
	}
	if _, err := mail.ParseAddress(email); err != nil {
		c.JSON(http.StatusBadRequest, SubscribeResponse{
			Error: "bad email",
			Token: "",
		})
		return
	}

	err := dbHandle.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&db.User{Token: token, Email: email}).Error; err != nil {
			return err
		}

		for _, topic := range topicsSlice {
			if err := tx.Create(&db.UsersTopicsRel{Email: email, Topic: topic}).Error; err != nil {
				return err
			}
		}

		return nil
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, SubscribeResponse{
			Error: "database error",
			Token: "",
		})
		return
	}

	c.JSON(http.StatusOK, SubscribeResponse{
		Error: "",
		Token: token,
	})
}
