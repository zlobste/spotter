package data

import "time"

type Confirmation struct {
	UserId    uint64    `db:"user_id"`
	TimerId   uint64    `db:"timer_id"`
	Date      time.Time `db:"date"`
	Confirmed bool      `db:"confirmed"`
}
