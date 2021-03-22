package data

const (
	RoleTypeAdmin = iota
	RoleTypeManager
	RoleTypeUser
)

type Payment struct {
	Id       uint64 `db:"id"`
	Name     string `db:"name"`
	Surname  string `db:"surname"`
	Email    string `db:"email"`
	Password string `db:"password"`
	Balance  int64  `db:"balance"`
	Role     uint64 `db:"role"`
}
