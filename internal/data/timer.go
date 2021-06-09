package data

import "time"

type Timer struct {
	Id        uint64    `db:"id" json:"id"`
	UserId    uint64    `db:"user_id" json:"user_id"`
	StartTime time.Time `db:"start_time" json:"start_time"`
	EndTime   time.Time `db:"end_time" json:"end_time"`
	Pending   bool      `db:"pending" json:"pending"`
}

func (t Timer) ToMap() map[string]interface{} {
	result := map[string]interface{}{
		"user_id":    t.UserId,
		"start_time": t.StartTime,
		"end_time":   t.EndTime,
		"pending":    t.Pending,
	}

	return result
}

func (t Timer) ToReturn() map[string]interface{} {
	result := map[string]interface{}{
		"id":         t.Id,
		"user_id":    t.UserId,
		"start_time": t.StartTime,
		"end_time":   t.EndTime,
		"pending":    t.Pending,
	}

	return result
}
