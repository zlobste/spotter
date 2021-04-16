package data

import "time"

type Timer struct {
	Id        uint64        `db:"id"`
	GroupId   uint64        `db:"group_id"`
	StartTime time.Time     `db:"start_time"`
	Duration  time.Duration `db:"duration"`
}

type TimersStorage interface {}