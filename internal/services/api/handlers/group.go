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

func CreateGroupHandler(w http.ResponseWriter, r *http.Request) {
	log := context.Log(r)

	request, err := requests.NewCreateGroupRequest(r)
	if err != nil {
		if verr, ok := err.(ozzoval.Errors); ok {
			log.WithError(verr).Debug("failed to parse create user request")
			utils.Respond(w, http.StatusBadRequest, utils.Message(fmt.Sprintf("request was invalid in some way: %s", verr.Error())))
			return
		}
		log.WithError(err).Error("failed to parse create user request")
		utils.Respond(w, http.StatusInternalServerError, utils.Message("something bad happened parsing the request"))
		return
	}
	err = context.Groups(r).CreateGroup(request.Data)
	if err != nil {
		log.WithError(err).Error("failed to create group")
		utils.Respond(w, http.StatusInternalServerError, utils.Message("something bad happened creating the group"))
		return
	}

	utils.Respond(w, http.StatusOK, utils.Message(true))
}

func GetGroupHandler(w http.ResponseWriter, r *http.Request) {
	groupId := chi.URLParam(r, "id")
	if groupId == "" {
		utils.Respond(w, http.StatusForbidden, utils.Message("Group id is empty"))
		return
	}
	id, err := strconv.Atoi(groupId)
	if groupId == "" {
		utils.Respond(w, http.StatusForbidden, utils.Message("Invalid id"))
		return
	}

	group, err := context.Groups(r).GetGroupById(uint64(id))
	log := context.Log(r)
	if err != nil {
		log.WithError(err).Error("failed to find group")
		utils.Respond(w, http.StatusInternalServerError, utils.Message("something bad happened trying to find the group"))
		return
	}

	utils.Respond(w, http.StatusOK, utils.Message(group.ToReturn()))
}
