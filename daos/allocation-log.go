package daos

import (
	"fmt"
	"gauntlet-be/m/v2/database"
	"time"

	"github.com/google/uuid"
)

type AllocationLogs struct {
	Id                int64     `json:"id"`
	VoucherId         int64     `json:"voucher_id"`
	AllocationAmount  int64     `json:"allocation_amount"`
	IncreasedAmount   int64     `json:"increased_amount"`
	Type              string    `json:"type"`
	AllocationStatus  string    `json:"allocation_status"`
	ClientCode        string    `json:"client_code"`
	AllocationRequest string    `json:"allocation_request"`
	Segment           string    `json:"segment"`
	Identifier        uuid.UUID `json:"identifier"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
	DeletedAt         time.Time `json:"deleted_at"`
}

func GetAllAllocationLogs() (logs []AllocationLogs, err error) {
	err = database.Db.Table("nse_allocations").Find(&logs).Error
	if err != nil {
		fmt.Println("Error getting allocation logs in GetAllAllocationLogs", err)
		return
	}

	return logs, nil
}
