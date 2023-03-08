package services

import (
	"fmt"
	"gauntlet-be/m/v2/daos"
	"gauntlet-be/m/v2/globals"
	"net/http"
)

func GetNSELedhersLogs() ([]daos.NSELedgersLogs, *globals.HttpError) {
	resp, err := daos.GetAllNSELedgersLogs()
	fmt.Println("NSE ledgers logs : ", resp)
	if err != nil {
		return nil, globals.NewHttpError(http.StatusInternalServerError, "error getting allocation logs in AllocationLogs", err)
	}
	return resp, nil
}
