package model

import (
	"encoding/json"
	"errors"
	"io"
	"regexp"
	"time"
)

var (
	ErrInvalidName     = errors.New("invalid name")
	ErrInvalidWebsite  = errors.New("invalid website")
	ErrInvalidTwitter  = errors.New("invalid twitter")
	ErrInvalidFacebook = errors.New("invalid facebook")
	ErrInvalidLocation = errors.New("invalid location")
)

var (
	regexpURL     = regexp.MustCompile(`|^http(s)?://[a-z0-9-]+(.[a-z0-9-]+)*(:[0-9]+)?(/.*)?$|i`)
	regexpTwitter = regexp.MustCompile(`/@([A-Za-z0-9_]{1,15})/`)
)

type Hackathon struct {
	ID        int64     `db:"id"        json:"id"`
	Created   time.Time `db:"created"   json:"created"`
	Updated   time.Time `db:"updated"   json:"updated"`
	Name      string    `db:"name"      json:"name"`
	Website   string    `db:"website"   json:"website"`
	Twitter   string    `db:"twitter"   json:"twitter,omitempty"`
	Facebook  string    `db:"facebook"  json:"facebook,omitempty"`
	Date      time.Time `db:"date"      json:"date"`
	Location  string    `db:"location"  json:"location"`
	Latitude  float64   `db:"latitude"  json:"latitude"`
	Longitude float64   `db:"longitude" json:"longitude"`
	Approved  bool      `db:"approved"  json:"-"`
}

func NewHackathon(jsonReader io.Reader) (*Hackathon, error) {
	var hackathon Hackathon
	if err := json.NewDecoder(jsonReader).Decode(&hackathon); err != nil {
		return nil, err
	}

	hackathon.Latitude = 0
	hackathon.Longitude = 0

	if err := hackathon.validate(); err != nil {
		return nil, err
	}

	return &hackathon, nil
}

func (h *Hackathon) validate() error {
	switch {
	case len(h.Name) == 0 || len(h.Name) > 255:
		return ErrInvalidName
	case regexpURL.MatchString(h.Website) == false:
		return ErrInvalidWebsite
	case len(h.Location) == 0 || len(h.Location) > 255:
		return ErrInvalidLocation
	}

	if len(h.Twitter) > 0 {
		if !regexpTwitter.MatchString(h.Twitter) {
			return ErrInvalidTwitter
		}
	}

	if len(h.Facebook) > 0 {
		// TODO: Check to see if it's actually a Facebook URL, and not just any
		// URL.
		if !regexpURL.MatchString(h.Facebook) {
			return ErrInvalidFacebook
		}
	}

	return nil
}
