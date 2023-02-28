package daos

import (
	"fmt"
	"gauntlet-be/m/v2/daos/models"
	"gauntlet-be/m/v2/database"
)

func GetAllAllocationLogs() (logs []models.AllocationLogs, err error) {

	database.Db.Table("nse_allocations").Find(&logs)
	if err != nil {
		fmt.Println("Error getting allocation logs in GetAllAllocationLogs", err)
		return
	}

	return logs, nil
}
