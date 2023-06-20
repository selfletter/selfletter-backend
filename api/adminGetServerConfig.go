package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"selfletter-backend/config"
	"selfletter-backend/db"
)

type adminGetServerConfigRequest struct {
	Key string `json:"key" binding:"required"`
}

type adminGetServerConfigResponse struct {
	Error  string        `json:"error"`
	Config config.Config `json:"config"`
}

func AdminGetServerConfig(c *gin.Context) {
	var request adminGetServerConfigRequest
	dbHandle := db.GetDatabaseHandle()

	err := c.BindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, adminGetServerConfigResponse{
			Error:  "bad json",
			Config: config.Config{},
		})
		return
	}

	if dbHandle.First(&db.AdminKey{}, "key = ?", request.Key).RowsAffected == 0 {
		c.JSON(http.StatusForbidden, adminGetServerConfigResponse{
			Error:  "invalid admin key",
			Config: config.Config{},
		})
		return
	}

	c.JSON(http.StatusOK, adminGetServerConfigResponse{
		Error:  "",
		Config: config.GetConfig(),
	})
}
