package data

const (
	RoleTypeAdmin = iota
	RoleTypeManager
	RoleTypeUser
)

type User struct {
	Id       uint64  `db:"id"`
	Name     string  `db:"name"`
	Surname  string  `db:"surname"`
	Email    string  `db:"email"`
	Password string  `db:"password"`
	Balance  float64 `db:"balance"`
	Salary   float64 `db:"salary"`
	Role     uint64  `db:"role"`
}

func (u User) ToMap() map[string]interface{} {
	result := map[string]interface{}{
		"name":     u.Name,
		"surname":  u.Surname,
		"email":    u.Email,
		"password": u.Password,
		"balance":  u.Balance,
		"salary":   u.Salary,
		"role":     u.Role,
	}

	return result
}

func (u User) ToReturn() map[string]interface{} {
	result := map[string]interface{}{
		"user_id": u.Id,
		"name":    u.Name,
		"surname": u.Surname,
		"email":   u.Email,
		"balance": u.Balance,
		"salary":  u.Salary,
		"role":    u.Role,
	}

	return result
}
