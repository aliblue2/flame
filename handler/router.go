package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/siongui/instago"
)

func RouterHandler(server *gin.Engine) {
	server.POST("/search", SearchUserInfoHandler)
}

func SearchUserInfoHandler(context *gin.Context) {
	var jsonData map[string]interface{}

	err := context.ShouldBindJSON(&jsonData)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	username := jsonData["username"].(string)

	profile, err := instago.GetUserInfoNoLogin(username)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"profile": profile})

}
