package daos

import (
	"fmt"
	"gauntlet-be/m/v2/daos/models"
	"gauntlet-be/m/v2/database"
)

func GetAllRequestLogs() (logs []models.RequestLogs, err error) {

	// 	errors := db.First(&user).Limit(10).Find(&users).GetErrors()

	// fmt.Println(len(errors))

	// for _, err := range errors {
	//   fmt.Println(err)
	// }

	err = database.Db.Table("allocation_requests").Find(&logs).Error

	if err != nil {
		fmt.Println("Error getting Request logs in GetAllRequestLogs", err)
		return
	}
	fmt.Println("GetAllRequestLogs ", logs)

	return logs, nil
}
