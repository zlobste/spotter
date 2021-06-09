package handlers

import (
	"github.com/go-chi/chi"
	"github.com/zlobste/spotter/internal/context"
	"github.com/zlobste/spotter/internal/utils"
	"net/http"
	"strconv"
)

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
		utils.Respond(w, http.StatusInternalServerError, utils.Message(err.Error()))
		return
	}

	utils.Respond(w, http.StatusOK, utils.Message(timer.ToReturn()))
}

func GetTimersByDriverHandler(w http.ResponseWriter, r *http.Request) {
	userId := chi.URLParam(r, "user_id")
	if userId == "" {
		utils.Respond(w, http.StatusForbidden, utils.Message("id is empty"))
		return
	}
	id, err := strconv.Atoi(userId)
	if userId == "" {
		utils.Respond(w, http.StatusForbidden, utils.Message("invalid id"))
		return
	}
	timers, err := context.Timers(r).GetTimersByDriver(uint64(id))
	log := context.Log(r)
	if err != nil {
		log.WithError(err).Error("failed to find timers")
		utils.Respond(w, http.StatusInternalServerError, utils.Message(err.Error()))
		return
	}
	utils.Respond(w, http.StatusOK, utils.Message(timers))
}

func GetPendingTimerHandler(w http.ResponseWriter, r *http.Request) {
	userId := chi.URLParam(r, "user_id")
	if userId == "" {
		utils.Respond(w, http.StatusForbidden, utils.Message("id is empty"))
		return
	}
	id, err := strconv.Atoi(userId)
	if userId == "" {
		utils.Respond(w, http.StatusForbidden, utils.Message("invalid id"))
		return
	}
	timer, err := context.Timers(r).GetPendingTimer(uint64(id))
	log := context.Log(r)
	if err != nil {
		log.WithError(err).Error("failed to find pending timer")
		utils.Respond(w, http.StatusInternalServerError, utils.Message(err.Error()))
		return
	}
	utils.Respond(w, http.StatusOK, utils.Message(timer))
}

func StartTimerHandler(w http.ResponseWriter, r *http.Request) {
	userId := chi.URLParam(r, "user_id")
	if userId == "" {
		utils.Respond(w, http.StatusForbidden, utils.Message("User id is empty"))
		return
	}
	id, err := strconv.Atoi(userId)
	if err != nil {
		utils.Respond(w, http.StatusForbidden, utils.Message("Invalid id"))
		return
	}
	log := context.Log(r)
	if err := context.Timers(r).StartTimer(uint64(id)); err != nil {
		log.WithError(err).Error("failed to start timer")
		utils.Respond(w, http.StatusInternalServerError, utils.Message(err.Error()))
		return
	}

	utils.Respond(w, http.StatusOK, utils.Message(true))
}

func StopTimerHandler(w http.ResponseWriter, r *http.Request) {
	userId := chi.URLParam(r, "user_id")
	if userId == "" {
		utils.Respond(w, http.StatusForbidden, utils.Message("User id is empty"))
		return
	}
	id, err := strconv.Atoi(userId)
	if err != nil {
		utils.Respond(w, http.StatusForbidden, utils.Message("Invalid id"))
		return
	}
	log := context.Log(r)
	if err := context.Timers(r).StopTimer(uint64(id)); err != nil {
		log.WithError(err).Error("failed to stop timer")
		utils.Respond(w, http.StatusInternalServerError, utils.Message(err.Error()))
		return
	}

	utils.Respond(w, http.StatusOK, utils.Message(true))
}
