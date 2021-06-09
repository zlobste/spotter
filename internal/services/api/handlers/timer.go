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

func CreateTimerHandler(w http.ResponseWriter, r *http.Request) {
	log := context.Log(r)

	request, err := requests.NewCreateTimerRequest(r)
	if err != nil {
		if verr, ok := err.(ozzoval.Errors); ok {
			log.WithError(verr).Debug("failed to parse create timer request")
			utils.Respond(w, http.StatusBadRequest, utils.Message(fmt.Sprintf("request was invalid in some way: %s", verr.Error())))
			return
		}
		log.WithError(err).Error("failed to parse create timer request")
		utils.Respond(w, http.StatusInternalServerError, utils.Message("something bad happened parsing the request"))
		return
	}

	err = context.Timers(r).CreateTimer(request.Data)
	if err != nil {
		log.WithError(err).Error("failed to create timer")
		utils.Respond(w, http.StatusInternalServerError, utils.Message("something bad happened creating the timer"))
		return
	}

	utils.Respond(w, http.StatusOK, utils.Message(true))
}

func GetTimerHandler(w http.ResponseWriter, r *http.Request) {
	timerId := chi.URLParam(r, "timer_id")
	if timerId == "" {
		utils.Respond(w, http.StatusForbidden, utils.Message("Timer id is empty"))
		return
	}

	id, err := strconv.Atoi(timerId)
	if timerId == "" {
		utils.Respond(w, http.StatusForbidden, utils.Message("Invalid id"))
		return
	}

	timer, err := context.Timers(r).GetTimerById(uint64(id))
	log := context.Log(r)
	if err != nil {
		log.WithError(err).Error("failed to find timer")
		utils.Respond(w, http.StatusInternalServerError, utils.Message("something bad happened trying to find the timer"))
		return
	}

	utils.Respond(w, http.StatusOK, utils.Message(timer.ToReturn()))
}

func GetTimersByDriverHandler(w http.ResponseWriter, r *http.Request) {
	driverId := chi.URLParam(r, "driver_id")
	if driverId == "" {
		utils.Respond(w, http.StatusForbidden, utils.Message("id is empty"))
		return
	}
	id, err := strconv.Atoi(driverId)
	if driverId == "" {
		utils.Respond(w, http.StatusForbidden, utils.Message("invalid id"))
		return
	}
	timers, err := context.Timers(r).GetTimersByDriver(uint64(id))
	log := context.Log(r)
	if err != nil {
		log.WithError(err).Error("failed to find timers")
		utils.Respond(w, http.StatusInternalServerError, utils.Message("something bad happened trying to find the timers"))
		return
	}
	utils.Respond(w, http.StatusOK, utils.Message(timers))
}
