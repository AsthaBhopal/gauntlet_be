package models

type AllocationLogs struct {
	// allocation_logs
	Id                int64  `json:"id"`
	VoucherId         int64  `json:"voucher_id"`
	AllocationAmount  int64  `json:"allocation_amount"`
	IncreasedAmount   int64  `json:"increased_amount"`
	Type              string `json:"type"`
	AllocationStatus  string `json:"allocation_status"`
	ClientCode        string `json:"client_code"`
	AllocationRequest string `json:"allocation_request"`
	Segment           string `json:"segment"`
	CreatedAt         string `json:"created_at"`
	UpdatedAt         string `json:"updated_at"`
	DeletedAt         string `json:"deleted_at"`
}
