package models

import (
	"io"
	"os"
	"path/filepath"
	"time"

	"github.com/gobuffalo/buffalo/binding"
	//"github.com/gobuffalo/logger"
	"github.com/gobuffalo/pop"
	uuid "github.com/gobuffalo/uuid"
	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type Venue struct {
	ID          uuid.UUID    `json:"id" db:"id"`
	CreatedAt   time.Time    `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time    `json:"updated_at" db:"updated_at"`
	Name        string       `json:"name" db:"name"`
	Description string       `json:"description" db:"description"`
	Photo       string       `json:"photo" db:"photo"`
	PhotoFile   binding.File `json:"photo_file" db:"-" form:"photo_file"`
	OwnerID     uuid.UUID    `json:"owner_id" db:"owner_id"`
}

type Venues []Venue

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
func (p *Venue) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.StringIsPresent{Field: p.Name, Name: "Name"},
		&validators.StringIsPresent{Field: p.Description, Name: "Description"},
	), nil
}

// BeforeCreate is a callback that preprocess the struct before saving it to the database
func (v *Venue) BeforeCreate(tx *pop.Connection) error {
	v.Photo = v.PhotoFile.Filename
	return nil
}

// AfterCreate is a callback that saves the file after the venue is saved to the database
func (v *Venue) AfterCreate(tx *pop.Connection) error {
	if !v.PhotoFile.Valid() {
		logrus.Infof("Invalid photo file: %v", v.PhotoFile)
		return nil
	}
	dir := filepath.Join(".", "uploads/photos")
	if err := os.MkdirAll(dir, 0755); err != nil {
		return errors.WithStack(err)
	}
	f, err := os.Create(filepath.Join(dir, v.PhotoFile.Filename))
	if err != nil {
		return errors.WithStack(err)
	}
	defer f.Close()
	_, err = io.Copy(f, v.PhotoFile)
	return err
}
