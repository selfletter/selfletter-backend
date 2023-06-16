package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"selfletter-backend/db"
	"selfletter-backend/helpers"
)

type AdminAddAdminKeyRequest struct {
	Key string `json:"key" binding:"required"`
}

type AdminAddAdminKeyResponse struct {
	Error string `json:"error"`
	Key   string `json:"key"`
}

func AdminAddAdminKey(c *gin.Context) {
	var request AdminAddAdminKeyRequest
	dbHandle := db.GetDatabaseHandle()

	err := c.BindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, AdminAddAdminKeyResponse{
			Error: "bad json",
			Key:   "",
		})
		return
	}

	if dbHandle.First(&db.AdminKey{}, "key = ?", request.Key).RowsAffected == 0 {
		c.JSON(http.StatusForbidden, AdminAddAdminKeyResponse{
			Error: "invalid admin key",
			Key:   "",
		})
		return
	}

	key := helpers.GenerateSecureToken()
	for i := 0; i < 10; i++ {
		if dbHandle.First(&db.User{}, "key = ?", key).RowsAffected != 0 {
			key = helpers.GenerateSecureToken()
		} else {
			break
		}
		if i == 9 {
			c.JSON(http.StatusInternalServerError, AdminAddAdminKeyResponse{
				Error: "too many collisions",
				Key:   "",
			})
			return
		}
	}

	if err := dbHandle.Create(&db.AdminKey{Key: key}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, AdminAddAdminKeyResponse{
			Error: "database error",
			Key:   "",
		})
		return
	}

	file, err := os.OpenFile("admin_keys.txt", os.O_RDWR, os.ModeAppend)
	if err != nil {
		c.JSON(http.StatusMultiStatus, AdminAddAdminKeyResponse{
			Error: "warning: admin key added to database, but not saved in file on server",
			Key:   key,
		})
		return
	}
	defer file.Close()

	_, _ = file.Seek(0, 2)
	_, err = file.WriteString("\n" + key)
	if err != nil {
		c.JSON(http.StatusMultiStatus, AdminAddAdminKeyResponse{
			Error: "warning: admin key added to database, but not saved in file on server",
			Key:   key,
		})
		return
	}

	_ = file.Sync()
	if err != nil {
		c.JSON(http.StatusMultiStatus, AdminAddAdminKeyResponse{
			Error: "warning: admin key added to database, but not saved in file on server",
			Key:   key,
		})
		return
	}

	c.JSON(http.StatusOK, AdminAddAdminKeyResponse{
		Error: "",
		Key:   key,
	})
}
