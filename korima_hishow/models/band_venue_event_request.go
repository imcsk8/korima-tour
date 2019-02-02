package models

import (
	"time"

	"github.com/gobuffalo/buffalo/binding"
	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/pop/nulls"
	uuid "github.com/gobuffalo/uuid"
	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"
)

type BandVenueEventRequest struct {
	ID          uuid.UUID `json:"id" db:"id"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
	BandID      uuid.UUID `json:"band_id" db:"band_id"`
	VenueID     uuid.UUID `json:"venue_id" db:"venue_id"`
	RequestDate time.Date `json:"request_date" db:"request_date"`
}

type BandVenueEventRequests []BandVenueEventRequest

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
/*ADD DATE VALIDATION func (p *BandVenueEventRequest) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.StringIsPresent{Field: p.Name, Name: "Name"},
		&validators.StringIsPresent{Field: p.Description, Name: "Description"},
	), nil
}*/

// BeforeSave is a callback that sets fields to be added to the database
func (v *BandVenueEventRequest) BeforeSave(tx *pop.Connection) error {
	return BeforeSaveFile(v)
}

// AfterSave is a callback that saves the file after the request is saved to the database
func (v *BandVenueEventRequest) AfterSave(tx *pop.Connection) error {
	return AfterSaveFile(v)
}

// GetID returns a string representation of the ID
func (v *BandVenueEventRequest) GetID() string {
	return v.ID.String()
}
