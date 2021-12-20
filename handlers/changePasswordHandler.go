package handlers

import (
	"net/http"
	"rbarrero/visago/gobex"

	"github.com/gin-gonic/gin"
)

type UserData struct {
	User    string `form:"username" json:"username" binding:"required"`
	Pass    string `form:"password" json:"password" binding:"required"`
	NewPass string `form:"newpassword" json:"newpassword" binding:"required"`
}

// ChangePasswordHandler endpoint for change user's pass
func ChangePasswordHandler(c *gin.Context) {
	var userData UserData
	gobex, ok := c.MustGet("gobex").(gobex.Gobex)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Can't connect to active directory"})
	}
	if err := c.ShouldBindJSON(&userData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	if gobex.ChangeUserPassword(userData.User, userData.Pass, userData.NewPass) {
		c.JSON(http.StatusOK, gin.H{"message": "Password updated successfully"})
	}
}
