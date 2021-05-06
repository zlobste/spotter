package data

import "time"

type Voting struct {
	Id          uint64    `db:"id" json:"id"`
	Victim      uint64    `db:"victim" json:"victim"`
	Type        uint64    `db:"type" json:"type"`
	Title       string    `db:"title" json:"title"`
	Description string    `db:"description" json:"description"`
	EndTime     time.Time `db:"end_time" json:"end_time"`
}

func (v Voting) ToMap() map[string]interface{} {
	result := map[string]interface{}{
		"victim":      v.Victim,
		"type":        v.Type,
		"title":       v.Title,
		"description": v.Description,
		"end_time":    v.EndTime,
	}

	return result
}

func (v Voting) ToReturn() map[string]interface{} {
	result := map[string]interface{}{
		"id":          v.Id,
		"victim":      v.Victim,
		"type":        v.Type,
		"title":       v.Title,
		"description": v.Description,
		"end_time":    v.EndTime,
	}

	return result
}
