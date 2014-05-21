package handler

import (
	"net/http"

	"github.com/zachlatta/hackwatch/backend/database"
)

func GetApprovedHackathons(w http.ResponseWriter, r *http.Request) *AppError {
	hackathons, err := database.GetApprovedHackathons()
	if err != nil {
		return &AppError{err, "error getting hackathons from database",
			http.StatusInternalServerError}
	}

	return renderJSON(w, hackathons, http.StatusOK)
}
