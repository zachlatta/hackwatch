package handler

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
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

func ApproveHackathon(w http.ResponseWriter, r *http.Request) *AppError {
	// TODO: Add authorization
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		return &AppError{err, "bad id", http.StatusBadRequest}
	}

	hackathon, err := database.GetHackathon(id)
	if err != nil {
		if err == sql.ErrNoRows {
			return &AppError{err, "not found", http.StatusNotFound}
		}

		return &AppError{err, "could not get hackathon from database",
			http.StatusInternalServerError}
	}

	hackathon.Approved = true

	err = database.SaveHackathon(hackathon)
	if err != nil {
		return &AppError{err, "could not save approved hackathon",
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
