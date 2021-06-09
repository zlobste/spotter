package handlers

import (
	"github.com/go-chi/chi"
	"github.com/zlobste/spotter/internal/context"
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
