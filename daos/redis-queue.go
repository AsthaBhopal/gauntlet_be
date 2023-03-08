package daos

import (
	"fmt"
	"gauntlet-be/m/v2/database"
	"time"
)

type NSELedgersLogs struct {
	Id            int64     `json:"id"`
	ClientCode    string    `json:"client_code"`
	Segment       string    `json:"segment"`
	Selection     string    `json:"selection"`
	Amount        float64   `json:"amount"`
	PendingAmount float64   `json:"pending_amount"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	DeletedAt     time.Time `json:"deleted_at"`
}

func GetAllNSELedgersLogs() (logs []NSELedgersLogs, err error) {

	database.Db.Table("nse_ledgers").Find(&logs)
	if err != nil {
		fmt.Println("Error getting nse ledgers logs in GetAllNSELedgersLogs", err)
		return
	}

	return logs, nil
}
