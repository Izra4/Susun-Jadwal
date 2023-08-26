package models

import "time"

type KrsAddReq struct {
	Totals int32 `json:"totals"`
	Userid int32 `json:"userid"`
}

type KrsUpdateReq struct {
	Totals int32 `json:"totals"`
	Userid int32 `json:"userid"`
	ID     int32 `json:"id"`
}

type KrsResult struct {
	ID        int32     `json:"id"`
	Createdat time.Time `json:"created_at"`
	Updatedat time.Time `json:"updated_at"`
	Deletedat time.Time `json:"deleted_at"`
	Totals    int32     `json:"totals"`
	Userid    int32     `json:"userid"`
}
