package services

import (
	"fmt"
	"gauntlet-be/m/v2/daos"
	"gauntlet-be/m/v2/globals"
	"net/http"
)

func GetAllocationLogs() ([]daos.AllocationLogs, *globals.HttpError) {
	resp, err := daos.GetAllAllocationLogs()
	fmt.Println("AllocationLogs : ", resp)
	if err != nil {
		return nil, globals.NewHttpError(http.StatusInternalServerError, "error getting allocation logs in AllocationLogs", err)
	}
	return resp, nil
}
