package data

import "time"

type Voting struct {
	Id          uint64    `db:"id"`
	CreatorId   uint64    `db:"creator_id"`
	Title       string    `db:"title"`
	Description string    `db:"description"`
	EndTime     time.Time `db:"end_time"`
}

type VotingsStorage interface {}