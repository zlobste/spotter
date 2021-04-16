package data

type Vote struct {
	UserId   uint64 `db:"user_id"`
	VotingId uint64 `db:"voting_id"`
	Decided  bool   `db:"decided"`
}

type VotesStorage interface {}