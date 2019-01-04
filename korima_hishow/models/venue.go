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

type Venue struct {
	ID           uuid.UUID    `json:"id" db:"id"`
	CreatedAt    time.Time    `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time    `json:"updated_at" db:"updated_at"`
	Name         string       `json:"name" db:"name"`
	Description  string       `json:"description" db:"description"`
	Photo        string       `json:"photo" db:"photo"`
	LocationText nulls.String `json:"location_text" db:"location_text"`
	Country      nulls.Int    `json:"country" db:"country"`
	State        nulls.Int    `json:"state" db:"state"`
	City         nulls.Int    `json:"city" db:"city"`
	PhotoFile    binding.File `json:"photo_file" db:"-" form:"photo_file"`
	OwnerID      uuid.UUID    `json:"owner_id" db:"owner_id"`
}

type Venues []Venue

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
func (p *Venue) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.StringIsPresent{Field: p.Name, Name: "Name"},
		&validators.StringIsPresent{Field: p.Description, Name: "Description"},
	), nil
}

// BeforeSave is a callback that sets fields to be added to the database
func (v *Venue) BeforeSave(tx *pop.Connection) error {
	return BeforeSaveFile(v)
}

// AfterSave is a callback that saves the file after the venue is saved to the database
func (v *Venue) AfterSave(tx *pop.Connection) error {
	return AfterSaveFile(v)
}

// SetPhotoName sets the photo filename parameter in the Photo field of te database
func (v *Venue) SetPhotoName() error {
	v.Photo = v.PhotoFile.Filename
	return nil
}

// ValidPhoto returns true if there's an uploaded photo from the form
func (v *Venue) ValidPhoto() bool {
	return v.PhotoFile.Valid() && v.PhotoFile.Filename != ""
}

// GetPhotoFilename returns the filename of the uploaded photo
func (v *Venue) GetPhotoFileName() string {
	return v.PhotoFile.Filename
}

func (v *Venue) GetPhotoStream() binding.File {
	return v.PhotoFile
}

// GetID returns a string representation of the ID
func (v *Venue) GetID() string {
	return v.ID.String()
}
