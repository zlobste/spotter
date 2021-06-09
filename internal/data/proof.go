package data

import "time"

type Proof struct {
	Id         uint64    `db:"id" json:"id"`
	TimerId    uint64    `db:"timer_id" json:"timer_id"`
	Time       time.Time `db:"time" json:"time"`
	Percentage float64   `db:"percentage" json:"percentage"`
	Confirmed  bool      `db:"confirmed" json:"confirmed"`
}

func (p Proof) ToMap() map[string]interface{} {
	result := map[string]interface{}{
		"timer_id":   p.TimerId,
		"time":       p.Time,
		"percentage": p.Percentage,
		"confirmed":  p.Confirmed,
	}

	return result
}
func (p Proof) ToReturn() map[string]interface{} {
	result := map[string]interface{}{
		"id":         p.Id,
		"timer_id":   p.TimerId,
		"time":       p.Time,
		"percentage": p.Percentage,
		"confirmed":  p.Confirmed,
	}
	return result
}
