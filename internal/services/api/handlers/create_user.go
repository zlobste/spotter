package handlers

import (
	"fmt"
	ozzoval "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/zlobste/spotter/internal/context"
	"github.com/zlobste/spotter/internal/services/api/requests"
	"github.com/zlobste/spotter/internal/utils"
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	log := context.Log(r)

	request, err := requests.NewCreateUserRequest(r)
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

	// encrypt users password
	request.Data.Password, err = utils.HashAndSalt(request.Data.Password)
	if err != nil {
		log.WithError(err).Error("failed to hash user password")
		utils.Respond(w, http.StatusInternalServerError, utils.Message("something bad happened hashing user password"))
		return
	}

	// check if we have such user already
	user, err := context.Users(r).GetUser(request.Data.Email)
	if err != nil {
		log.WithError(err).Error("failed to get user")
		utils.Respond(w, http.StatusInternalServerError, utils.Message("something bad happened"))
		return
	}
	if user != nil {
		log.WithError(err).Debug("specified user exist alreay")
		utils.Respond(w, http.StatusNotFound, utils.Message("specified user exist already"))
		return
	}

	err = context.Users(r).CreateUser(request.Data)
	if err != nil {
		log.WithError(err).Error("failed to create user")
		utils.Respond(w, http.StatusInternalServerError, utils.Message("something bad happened creating the user"))
		return
	}

	user, err = context.Users(r).GetUser(request.Data.Email)
	if err != nil {
		log.WithError(err).Error("failed to find user")
		utils.Respond(w, http.StatusInternalServerError, utils.Message("something bad happened trying to find the user"))
		return
	}

	utils.Respond(w, http.StatusOK, utils.Message(user.ToReturn()))
}
