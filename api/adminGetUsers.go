package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"selfletter-backend/db"
)

type adminGetUsersRequest struct {
	Key   string `json:"key" binding:"required"`
	Query string `json:"query"`
}

type adminGetUsersResponse struct {
	Error string    `json:"error"`
	Users []db.User `json:"users"`
}

func AdminGetUsers(c *gin.Context) {
	var request adminGetUsersRequest
	dbHandle := db.GetDatabaseHandle()

	err := c.BindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, adminGetUsersResponse{
			Error: "bad json",
			Users: nil,
		})
		return
	}

	if dbHandle.First(&db.AdminKey{}, "key = ?", request.Key).RowsAffected == 0 {
		c.JSON(http.StatusForbidden, adminGetUsersResponse{
			Error: "invalid admin key",
			Users: nil,
		})
		return
	}

	if request.Query == "" {
		var users []db.User
		dbHandle.Find(&users)
		c.JSON(http.StatusOK, adminGetUsersResponse{
			Error: "",
			Users: users,
		})
		return
	}

	var usersByEmail []db.User
	if err := dbHandle.Where("email LIKE ?", "%"+request.Query+"%").Find(&usersByEmail).Error; err != nil {
		c.JSON(http.StatusInternalServerError, adminGetUsersResponse{
			Error: "database error",
			Users: nil,
		})
		return
	}
	var usersByToken []db.User
	if err := dbHandle.Where("token LIKE ?", "%"+request.Query+"%").Find(&usersByToken).Error; err != nil {
		c.JSON(http.StatusInternalServerError, adminGetUsersResponse{
			Error: "database error",
			Users: nil,
		})
		return
	}

	users := append(usersByEmail, usersByToken...)

	if len(users) == 0 {
		c.JSON(http.StatusOK, adminGetUsersResponse{
			Error: "",
			Users: nil,
		})
		return
	}

	c.JSON(http.StatusOK, adminGetUsersResponse{
		Error: "",
		Users: users,
	})
}
