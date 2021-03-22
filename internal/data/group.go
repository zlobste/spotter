package data

type Group struct {
	Id          uint64 `db:"id"`
	Title       string `db:"title"`
	Description string `db:"description"`
	Level       uint64 `db:"level"`
}
