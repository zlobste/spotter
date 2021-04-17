package data

import "time"

type Confirmation struct {
	UserId    uint64    `db:"user_id"`
	TimerId   uint64    `db:"timer_id"`
	Date      time.Time `db:"date"`
	Confirmed bool      `db:"confirmed"`
}

func (c Confirmation) ToMap() map[string]interface{} {
	result := map[string]interface{}{
		"user_id":   c.UserId,
		"timer_id":  c.TimerId,
		"date":      c.Date,
		"confirmed": c.Confirmed,
	}

	return result
}