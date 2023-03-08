package services

import (
	"fmt"
	"gauntlet-be/m/v2/daos"
	"gauntlet-be/m/v2/globals"
	"net/http"
)

func GetRequestLogs() ([]daos.FormatedRequestLogs, *globals.HttpError) {
	resp, err := daos.GetAllRequestLogs()
	fmt.Println("GetRequestLogs : ", resp)
	respJson := []daos.FormatedRequestLogs{}

	for _, log := range resp {
		respJson = append(respJson, log.Format())
	}

	if err != nil {
		return nil, globals.NewHttpError(http.StatusInternalServerError, "error getting request logs in GetRequestLogs", err)
	}
	return respJson, nil
}
