package requests

import (
	"encoding/json"
	"github.com/pkg/errors"
	"net/http"
)

type MakeProofRequest struct {
	TimerId    uint64  `json:"timer_id"`
	Percentage float64 `json:"percentage"`
}

func NewMakeProofRequest(r *http.Request) (* MakeProofRequest, error) {
	req :=  MakeProofRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, errors.Wrap(err, "failed to decode request body")
	}

	return &req, nil
}
