package data

type Vote struct {
	UserId   uint64 `db:"user_id"`
	VotingId uint64 `db:"voting_id"`
	Decided  bool   `db:"decided"`
}

func (v Vote) ToMap() map[string]interface{} {
	result := map[string]interface{}{
		"user_id":   v.UserId,
		"voting_id": v.VotingId,
		"decided":   v.Decided,
	}

	return result
}
