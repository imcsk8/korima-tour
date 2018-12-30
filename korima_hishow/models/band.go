package models

import (
	"time"

	"github.com/gobuffalo/buffalo/binding"
	"github.com/gobuffalo/pop"
	uuid "github.com/gobuffalo/uuid"
	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"
)

type Band struct {
	ID          uuid.UUID    `json:"id" db:"id"`
	CreatedAt   time.Time    `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time    `json:"updated_at" db:"updated_at"`
	Name        string       `json:"title" db:"name"`
	Description string       `json:"content" db:"description"`
	Photo       string       `json:"photo" db:"photo"`
	PhotoFile   binding.File `json:"photo_file" db:"-" form:"photo_file"`
	OwnerID     uuid.UUID    `json:"owner_id" db:"owner_id"`
}

type Bands []Band

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
func (p *Band) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.StringIsPresent{Field: p.Name, Name: "Name"},
		&validators.StringIsPresent{Field: p.Description, Name: "Description"},
	), nil
}

// BeforeSave is a callback that preprocess the struct before saving it to the database
func (b *Band) BeforeSave(tx *pop.Connection) error {
	return BeforeSaveFile(b)
}

// AfterSave is a callback that saves the file after the venue is saved to the database
func (b *Band) AfterSave(tx *pop.Connection) error {
	return AfterSaveFile(b)
}

// SetPhotoName sets the photo filename parameter in the Photo field of te database
func (b *Band) SetPhotoName() error {
	b.Photo = b.PhotoFile.Filename
	return nil
}

// ValidPhoto returns true if there's an uploaded photo from the form
func (b *Band) ValidPhoto() bool {
	return b.PhotoFile.Valid() && b.PhotoFile.Filename != ""
}

// GetPhotoFilename returns the filename of the uploaded photo
func (b *Band) GetPhotoFileName() string {
	return b.PhotoFile.Filename
}

func (b *Band) GetPhotoStream() binding.File {
	return b.PhotoFile
}
