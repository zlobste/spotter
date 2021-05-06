package data

const (
	RoleTypeAdmin = iota
	RoleTypeManager
	RoleTypeUser
)

type User struct {
	Id       uint64  `db:"id" json:"id"`
	Name     string  `db:"name" json:"name"`
	Surname  string  `db:"surname" json:"surname"`
	Email    string  `db:"email" json:"email"`
	Password string  `db:"password" json:"password"`
	Balance  float64 `db:"balance" json:"balance"`
	Salary   float64 `db:"salary" json:"salary"`
	Role     uint64  `db:"role" json:"role"`
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

func (u *User) ToReturn() map[string]interface{} {
	result := map[string]interface{}{
		"id":      u.Id,
		"name":    u.Name,
		"surname": u.Surname,
		"email":   u.Email,
		"balance": u.Balance,
		"salary":  u.Salary,
		"role":    u.Role,
	}

	return result
}
