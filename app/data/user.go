package data

const (
	RoleTypeAdmin = iota
	RoleTypeManager
	RoleTypeUser
)

type User struct {
	Id       uint64 `db:"id"`
	Name     string `db:"name"`
	Surname  string `db:"surname"`
	Email    string `db:"email"`
	Password string `db:"password"`
	Balance  int64  `db:"balance"`
	Role     uint64 `db:"role"`
}

func (u User) ToMap() map[string]interface{} {
	result := map[string]interface{}{
		"name":     u.Name,
		"surname":  u.Surname,
		"email":    u.Email,
		"password": u.Password,
		"balance":  u.Balance,
		"role":     u.Role,
	}

	return result
}
