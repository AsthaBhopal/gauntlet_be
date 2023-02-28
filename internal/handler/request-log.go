package handler

import (
	"fmt"
	"gauntlet-be/m/v2/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetRequestLogs(c *gin.Context) {

	response, errResp := services.GetRequestLogs()
	fmt.Println("GetRequestLogs ", response)

	if errResp != nil {
		c.JSON(errResp.HttpErrorCode, gin.H{
			"msg": errResp.Msg,
			"err": errResp.Err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response)
}
