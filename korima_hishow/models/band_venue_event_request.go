package models

import (
	"time"

	//"github.com/gobuffalo/pop"
	//"github.com/gobuffalo/pop/nulls"
	uuid "github.com/gobuffalo/uuid"
	//"github.com/gobuffalo/validate"
	//"github.com/gobuffalo/validate/validators"
	"github.com/gobuffalo/pop"
	"github.com/pkg/errors"
)

type BandVenueEventRequest struct {
	ID          uuid.UUID `json:"id" db:"id"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
	BandID      uuid.UUID `json:"band_id" db:"band_id"`
	VenueID     uuid.UUID `json:"venue_id" db:"venue_id"`
	RequestDate time.Time `json:"request_date" db:"request_date"`
	OwnerID     uuid.UUID `json:"owner_id" db:"owner_id"`
	Status      int       `json:"status" db:"status"`
	BandName    string    `json:"name" db:"name"`
}

type BandVenueEventRequests []BandVenueEventRequest

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
/*ADD DATE VALIDATION func (p *BandVenueEventRequest) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.StringIsPresent{Field: p.Name, Name: "Name"},
		&validators.StringIsPresent{Field: p.Description, Name: "Description"},
	), nil
}*/

// GetID returns a string representation of the ID
func (b *BandVenueEventRequest) GetID() string {
	return b.ID.String()
}

func (b *BandVenueEventRequest) GetRequests(tx *pop.Connection) (BandVenueEventRequests, error) {
	br := BandVenueEventRequests{}
	if err := tx.All(br); err != nil {
		return nil, errors.WithStack(err)
	}
	return br, nil
}
