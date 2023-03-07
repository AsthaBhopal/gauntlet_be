package daos

import (
	"fmt"
	"gauntlet-be/m/v2/daos/models"
	"gauntlet-be/m/v2/database"
)

func GetAllNSELedgersLogs() (logs []models.NSELedgersLogs, err error) {

	database.Db.Table("nse_ledgers").Find(&logs)
	if err != nil {
		fmt.Println("Error getting nse ledgers logs in GetAllNSELedgersLogs", err)
		return
	}

	return logs, nil
}
