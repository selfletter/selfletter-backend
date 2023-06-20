package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"selfletter-backend/config"
	"selfletter-backend/db"
)

type adminWriteServerConfigRequest struct {
	Key    string        `json:"key" binding:"required"`
	Config config.Config `json:"config" binding:"required"`
}

type adminWriteServerConfigResponse struct {
	Error string `json:"error"`
}

func AdminWriteServerConfig(c *gin.Context) {
	var request adminWriteServerConfigRequest
	dbHandle := db.GetDatabaseHandle()

	err := c.BindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, adminWriteServerConfigResponse{Error: "bad json"})
		return
	}

	if dbHandle.First(&db.AdminKey{}, "key = ?", request.Key).RowsAffected == 0 {
		c.JSON(http.StatusForbidden, adminWriteServerConfigResponse{Error: "invalid admin key"})
		return
	}

	err = config.WriteConfig(request.Config)
	if err != nil {
		c.JSON(http.StatusInternalServerError, adminWriteServerConfigResponse{Error: "failed to write config on server"})
		return
	}

	c.JSON(http.StatusOK, adminWriteServerConfigResponse{Error: ""})
}
