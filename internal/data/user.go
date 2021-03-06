package data

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
)

const (
	RoleTypeAdmin = iota
	RoleTypeManager
	RoleTypeDriver
)

type User struct {
	Id       uint64 `db:"id" json:"id"`
	Name     string `db:"name" json:"name"`
	Surname  string `db:"surname" json:"surname"`
	Email    string `db:"email" json:"email"`
	Password string `db:"password" json:"password"`
	Role     uint64 `db:"role" json:"role"`
	Blocked  bool   `db:"blocked" json:"blocked"`
}

func (u User) ToMap() map[string]interface{} {
	result := map[string]interface{}{
		"name":     u.Name,
		"surname":  u.Surname,
		"email":    u.Email,
		"password": u.Password,
		"role":     u.Role,
		"blocked":  u.Blocked,
	}
	return result
}

func (u *User) ToReturn() map[string]interface{} {
	result := map[string]interface{}{
		"id":      u.Id,
		"name":    u.Name,
		"surname": u.Surname,
		"email":   u.Email,
		"role":    u.Role,
		"blocked": u.Blocked,
	}
	return result
}

func (u *User) EncryptPassword() error {
	if len(u.Password) == 0 {
		return errors.New("Wrong password length!")
	}
	enc, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.MinCost)
	if err != nil {
		return err
	}
	u.Password = string(enc)
	return nil
}

func (u *User) ComparePassword(password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(password), []byte(u.Password)) != nil
}
