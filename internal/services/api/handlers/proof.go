package handlers

import (
	"fmt"
	"github.com/go-chi/chi"
	ozzoval "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/zlobste/spotter/internal/context"
	"github.com/zlobste/spotter/internal/services/api/requests"
	"github.com/zlobste/spotter/internal/utils"
	"net/http"
	"strconv"
)

func GetProofsByTimerHandler(w http.ResponseWriter, r *http.Request) {
	timerId := chi.URLParam(r, "timer_id")
	if timerId == "" {
		utils.Respond(w, http.StatusForbidden, utils.Message("id is empty"))
		return
	}
	id, err := strconv.Atoi(timerId)
	if timerId == "" {
		utils.Respond(w, http.StatusForbidden, utils.Message("invalid id"))
		return
	}
	proofs, err := context.Proofs(r).GetProofsByTimer(uint64(id))
	log := context.Log(r)
	if err != nil {
		log.WithError(err).Error("failed to find proofs")
		utils.Respond(w, http.StatusInternalServerError, utils.Message(err.Error()))
		return
	}
	utils.Respond(w, http.StatusOK, utils.Message(proofs))
}

func MakeProofHandler(w http.ResponseWriter, r *http.Request) {
	log := context.Log(r)
	request, err := requests.NewMakeProofRequest(r)
	if err != nil {
		if verr, ok := err.(ozzoval.Errors); ok {
			log.WithError(verr).Debug("failed to parse create proof request")
			utils.Respond(w, http.StatusBadRequest, utils.Message(fmt.Sprintf("request was invalid in some way: %s", verr.Error())))
			return
		}
		log.WithError(err).Error("failed to parse create user request")
		utils.Respond(w, http.StatusInternalServerError, utils.Message("something bad happened parsing the request"))
		return
	}
	timer, err := context.Timers(r).GetTimerById(request.TimerId)
	if err != nil {
		log.WithError(err).Error("failed to find pending timer")
		utils.Respond(w, http.StatusInternalServerError, utils.Message(err.Error()))
		return
	}
	if !timer.Pending {
		log.WithError(err).Error("timer is not pending")
		utils.Respond(w, http.StatusInternalServerError, utils.Message("Timer is not pending"))
		return
	}
	if err := context.Proofs(r).MakeProof(request.TimerId, request.Percentage); err != nil {
		log.WithError(err).Error("failed to stop timer")
		utils.Respond(w, http.StatusInternalServerError, utils.Message(err.Error()))
		return
	}

	utils.Respond(w, http.StatusOK, utils.Message(true))
}
