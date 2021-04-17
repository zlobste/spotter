package data

import "time"

type Voting struct {
	Id          uint64    `db:"id"`
	CreatorId   uint64    `db:"creator_id"`
	Title       string    `db:"title"`
	Description string    `db:"description"`
	EndTime     time.Time `db:"end_time"`
}

func (v Voting) ToMap() map[string]interface{} {
	result := map[string]interface{}{
		"creator_id":  v.CreatorId,
		"title":       v.Title,
		"description": v.Description,
		"end_time":    v.EndTime,
	}

	return result
}

func (v Voting) ToReturn() map[string]interface{} {
	result := map[string]interface{}{
		"id":          v.Id,
		"creator_id":  v.CreatorId,
		"title":       v.Title,
		"description": v.Description,
		"end_time":    v.EndTime,
	}

	return result
}
