package models

import (
	"encoding/json"
	"time"
)

/*
	{
	  "status": "success",
	  "messages": {
	    "code": "01010000"
	  },
	  "data": {
	    "inquiryResponse": [
	      {
	        "curDate": "02-Feb-2023",
	        "segment": "FO",
	        "cmCode": "12227",
	        "tmCode": "12227",
	        "cpCode": "",
	        "cliCode": "600S1871",
	        "accType": "C",
	        "amt": 18554.16,
	        "filler1": "",
	        "filler2": "",
	        "filler3": "",
	        "filler4": "",
	        "filler5": "",
	        "filler6": "",
	        "action": "U",
	        "errCd": "01050102"
	      }
	    ]
	  }
	}
*/

type RequestLogs struct {
	Id            int64     `json:"id"`
	Response      string    `json:"response"`
	MessageId     string    `json:"message_id"`
	StatusCode    int64     `json:"status_code"`
	ProcessStatus int64     `json:"process_status"`
	Processed     string    `json:"processed"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	DeletedAt     time.Time `json:"deleted_at"`
}

func (rlogs *RequestLogs) Format() FormatedRequestLogs {
	var _resp Response
	json.Unmarshal([]byte(rlogs.Response), &_resp)

	var resp = FormatedRequestLogs{
		Id:            rlogs.Id,
		MessageId:     rlogs.MessageId,
		StatusCode:    rlogs.StatusCode,
		ProcessStatus: rlogs.ProcessStatus,
		Processed:     rlogs.Processed,
		CreatedAt:     rlogs.CreatedAt,
		UpdatedAt:     rlogs.UpdatedAt,
		DeletedAt:     rlogs.DeletedAt,
		Response:      _resp,
	}
	return resp
}

type FormatedRequestLogs struct {
	/*
		This is a wrapper struct to convert Response from escaped JSON
	*/
	Id            int64       `json:"id"`
	Response      interface{} `json:"response"`
	MessageId     string      `json:"message_id"`
	StatusCode    int64       `json:"status_code"`
	ProcessStatus int64       `json:"process_status"`
	Processed     string      `json:"processed"`
	CreatedAt     time.Time   `json:"created_at"`
	UpdatedAt     time.Time   `json:"updated_at"`
	DeletedAt     time.Time   `json:"deleted_at"`
}

type Response struct {
	Status   string `json:"status"`
	Messages struct {
		Code string `json:"code"`
	} `json:"messages"`
	Data struct {
		InquiryResponse []struct {
			CurDate string  `json:"curDate"`
			Segment string  `json:"segment"`
			CmCode  string  `json:"cmCode"`
			TmCode  string  `json:"tmCode"`
			CpCode  string  `json:"cpCode"`
			CliCode string  `json:"cliCode"`
			AccType string  `json:"accType"`
			Amt     float64 `json:"amt"`
			Filler1 string  `json:"filler1"`
			Filler2 string  `json:"filler2"`
			Filler3 string  `json:"filler3"`
			Filler4 string  `json:"filler4"`
			Filler5 string  `json:"filler5"`
			Filler6 string  `json:"filler6"`
			Action  string  `json:"action"`
			ErrCd   string  `json:"errCd"`
		} `json:"inquiryResponse"`
	} `json:"data"`
}
