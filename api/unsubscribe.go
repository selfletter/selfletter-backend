package api

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"selfletter-backend/db"
)

type unsubscribeResponse struct {
	Error string `json:"error"`
}

func Unsubscribe(c *gin.Context) {
	dbHandle := db.GetDatabaseHandle()
	token := c.Query("token")
	if token == "" {
		c.JSON(http.StatusBadRequest, unsubscribeResponse{Error: "empty token"})
		return
	}
	if dbHandle.First(&db.User{}, "token = ?", token).RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, unsubscribeResponse{Error: "invalid token"})
		return
	}

	err := dbHandle.Transaction(func(tx *gorm.DB) error {
		var user db.User
		if err := tx.Where("token = ?", token).Find(&user).Delete(&db.User{}).Error; err != nil {
			return err
		}

		if err := tx.Where("email = ?", user.Email).Delete(&db.UsersTopicsRel{}).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, unsubscribeResponse{Error: "database error"})
		return
	}

	c.JSON(http.StatusOK, unsubscribeResponse{Error: ""})
}
