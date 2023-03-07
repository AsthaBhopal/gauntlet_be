package handler

import (
	"gauntlet-be/m/v2/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetRequestLogs(c *gin.Context) {

	response, errResp := services.GetRequestLogs()
	if errResp != nil {
		c.JSON(errResp.HttpErrorCode, gin.H{
			"msg": errResp.Msg,
			"err": errResp.Err.Error(),
		})
		return
	}
	// type TempResp struct {
	// 	Data interface{}
	// }

	// s, _ := strconv.Unquote(datalist)

	// fmt.Println(s)

	// fmt.Println(datalist)

	// var tempResp map[string]interface{}
	// json.Unmarshal([]byte(s), &tempResp)

	c.JSON(http.StatusOK, response)
}
