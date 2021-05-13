package requests

import (
	"encoding/json"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/pkg/errors"
	"github.com/zlobste/spotter/internal/data"
	"net/http"
)

type CreateGroupRequest struct {
	Data data.Group `json:"data"`
}

func (r CreateGroupRequest) Validate() error {
	return validation.ValidateStruct(&r.Data,
		validation.Field(&r.Data.Title, validation.Required, validation.Length(4, 50)),
		validation.Field(&r.Data.Description, validation.Required, validation.Length(4, 500)),
	)
}

func NewCreateGroupRequest(r *http.Request) (*CreateGroupRequest, error) {
	req := CreateGroupRequest{}

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, errors.Wrap(err, "failed to decode request body")
	}

	return &req, req.Validate()
}

type AddUserToGroupRequest struct {
	Data data.UserGroup `json:"data"`
}

func NewAddUserToGroupRequest(r *http.Request) (*AddUserToGroupRequest, error) {
	req := AddUserToGroupRequest{}

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, errors.Wrap(err, "failed to decode request body")
	}

	return &req, nil
}
