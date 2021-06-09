package requests

import (
	"encoding/json"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/pkg/errors"
	"github.com/zlobste/spotter/internal/data"
	"net/http"
)

type RegisterRequest struct {
	Data data.User `json:"data"`
}

func (r RegisterRequest) Validate() error {
	return validation.ValidateStruct(
		&r.Data,
		validation.Field(&r.Data.Name, validation.Required, validation.Length(4, 50)),
		validation.Field(&r.Data.Surname, validation.Required, validation.Length(4, 50)),
		validation.Field(&r.Data.Email, validation.Required, is.Email),
		validation.Field(&r.Data.Password, validation.Required, validation.Length(4, 50)),
	)
}

func NewRegisterRequest(r *http.Request) (*RegisterRequest, error) {
	req := RegisterRequest{}
	req.Data.Role = data.RoleTypeDriver

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, errors.Wrap(err, "failed to decode request body")
	}

	return &req, req.Validate()
}


type LoginRequest struct {
	Data LoginData `json:"data"`
}

type LoginData struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     int64  `json:"role"`
}

type JWT struct {
	Token string `json:"token"`
}

func NewLoginRequest(r *http.Request) (*LoginRequest, error) {
	req := LoginRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, errors.Wrap(err, "failed to decode request body")
	}
	return &req, req.Validate()
}

func (r LoginRequest) Validate() error {
	return validation.ValidateStruct(
		&r.Data,
		validation.Field(&r.Data.Email, validation.Required, is.Email),
		validation.Field(&r.Data.Password, validation.Required, validation.Length(4, 50)),
	)
}
