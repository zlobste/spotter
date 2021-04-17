package data

import "time"

type Voting struct {
	Id          uint64    `db:"id"`
	Victim      uint64    `db:"victim"`
	Type        uint64    `db:"type"`
	Title       string    `db:"title"`
	Description string    `db:"description"`
	EndTime     time.Time `db:"end_time"`
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
