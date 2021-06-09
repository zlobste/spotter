package requests

import (
	"encoding/json"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/pkg/errors"
	"github.com/zlobste/spotter/internal/data"
	"net/http"
)

type CreateUserRequest struct {
	Data data.User `json:"data"`
}

func (r CreateUserRequest) Validate() error {
	return validation.ValidateStruct(&r.Data,
		validation.Field(&r.Data.Name, validation.Required, validation.Length(4, 50)),
		validation.Field(&r.Data.Surname, validation.Required, validation.Length(4, 50)),
		validation.Field(&r.Data.Email, validation.Required, is.Email),
		validation.Field(&r.Data.Password, validation.Required, validation.Length(4, 50)),
	)
}

func NewCreateUserRequest(r *http.Request) (*CreateUserRequest, error) {
	req := CreateUserRequest{}
	req.Data.Role = data.RoleTypeDriver

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, errors.Wrap(err, "failed to decode request body")
	}

	return &req, req.Validate()
}
