package handler

import (
	"net/http"

	"github.com/zachlatta/hackwatch/backend/database"
	"github.com/zachlatta/hackwatch/backend/model"
)

func NewHackathon(w http.ResponseWriter, r *http.Request) *AppError {
	defer r.Body.Close()

	// TODO: Add authorization
	hackathon, err := model.NewHackathon(r.Body)
	if err != nil {
		return &AppError{err, err.Error(), http.StatusBadRequest}
	}

	err = database.SaveHackathon(hackathon)
	if err != nil {
		return &AppError{err, "error creating hackathon",
			http.StatusInternalServerError}
	}

	return renderJSON(w, hackathon, http.StatusOK)
}

func GetApprovedHackathons(w http.ResponseWriter, r *http.Request) *AppError {
	hackathons, err := database.GetApprovedHackathons()
	if err != nil {
		return &AppError{err, "error getting hackathons from database",
			http.StatusInternalServerError}
	}

	return renderJSON(w, hackathons, http.StatusOK)
}
