package data

import "time"

type Timer struct {
	Id        uint64        `db:"id"`
	GroupId   uint64        `db:"group_id"`
	StartTime time.Time     `db:"start_time"`
	Duration  time.Duration `db:"duration"`
}

func (t Timer) ToMap() map[string]interface{} {
	result := map[string]interface{}{
		"group_id":   t.GroupId,
		"start_time": t.StartTime,
		"duration":   t.Duration,
	}

	return result
}

func (t Timer) ToReturn() map[string]interface{} {
	result := map[string]interface{}{
		"id":         t.Id,
		"group_id":   t.GroupId,
		"start_time": t.StartTime,
		"duration":   t.Duration,
	}

	return result
}
