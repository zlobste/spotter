package requests

import (
	"encoding/json"
	"github.com/pkg/errors"
	"github.com/zlobste/spotter/internal/data"
	"net/http"
)

type CreateTimerRequest struct {
	Data data.Timer `json:"data"`
}

func NewCreateTimerRequest(r *http.Request) (*CreateTimerRequest, error) {
	req := CreateTimerRequest{}

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, errors.Wrap(err, "failed to decode request body")
	}

	return &req, nil
}
