package database

import (
	"time"

	"github.com/zachlatta/hackwatch/backend/model"
)

func GetHackathon(id int64) (*model.Hackathon, error) {
	var hackathon model.Hackathon
	err := db.Get(&hackathon, "SELECT * FROM hackathons WHERE id=$1", id)
	if err != nil {
		return nil, err
	}
	return &hackathon, err
}

func GetHackathons() ([]*model.Hackathon, error) {
	hackathons := []*model.Hackathon{}
	err := db.Select(&hackathons, "SELECT * FROM hackathons ORDER BY id")
	if err != nil {
		return nil, err
	}
	return hackathons, nil
}

func SaveHackathon(h *model.Hackathon) error {
	if h.ID == 0 {
		h.Created = time.Now()
	}
	h.Updated = time.Now()

	err := db.QueryRowx("INSERT INTO hackathons (created, updated, name, website, twitter, facebook, date, location, latitude, longitude, approved) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11) RETURNING id", h.Created, h.Updated, h.Name, h.Website, h.Twitter, h.Facebook, h.Date, h.Location, h.Latitude, h.Longitude, h.Approved).Scan(&h.ID)
	if err != nil {
		return err
	}

	return nil
}

func DeleteHackathon(id int64) error {
	_, err := db.Exec("DELETE FROM hackathons WHERE id=$1", id)
	if err != nil {
		return err
	}
	return nil
}
