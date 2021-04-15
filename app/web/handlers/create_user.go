package handlers

import (
	"fmt"
	ozzoval "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/zlobste/spotter/app/context"
	"github.com/zlobste/spotter/app/utils/render"
	"github.com/zlobste/spotter/app/utils/validation"
	"github.com/zlobste/spotter/app/web/requests"
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	log := context.Log(r)

	request, err := requests.NewCreateUserRequest(r)
	if err != nil {
		if verr, ok := err.(ozzoval.Errors); ok {
			log.WithError(verr).Debug("failed to parse create user request")
			render.Respond(w, http.StatusBadRequest, render.Message(fmt.Sprintf("request was invalid in some way: %s", verr.Error())))
			return
		}
		log.WithError(err).Error("failed to parse create user request")
		render.Respond(w, http.StatusInternalServerError, render.Message("something bad happened parsing the request"))
		return
	}

	// encrypt users password
	request.Data.Password, err = validation.HashAndSalt(request.Data.Password)
	if err != nil {
		log.WithError(err).Error("failed to hash user password")
		render.Respond(w, http.StatusInternalServerError, render.Message("something bad happened hashing user password"))
		return
	}

	// check if we have such user already
	user, err := context.Users(r).GetUser(request.Data.Email)
	if err != nil {
		log.WithError(err).Error("failed to get user")
		render.Respond(w, http.StatusInternalServerError, render.Message("something bad happened"))
		return
	}
	if user != nil {
		log.WithError(err).Debug("specified user exist alreay")
		render.Respond(w, http.StatusNotFound, render.Message("specified user exist already"))
		return
	}

	err = context.Users(r).CreateUser(request.Data)
	if err != nil {
		log.WithError(err).Error("failed to create user")
		render.Respond(w, http.StatusInternalServerError, render.Message("something bad happened creating the user"))
		return
	}

	// check that user is created successfully
	user, err = context.Users(r).GetUser(request.Data.Email)
	if err != nil {
		log.WithError(err).Error("failed to find user")
		render.Respond(w, http.StatusInternalServerError, render.Message("something bad happened trying to find the user"))
		return
	}

	render.Respond(w, http.StatusOK, render.Message(user.ToReturn()))
}
