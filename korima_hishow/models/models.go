package models

import (
	"io"
	"log"
	"os"
	"path/filepath"

	"github.com/gobuffalo/buffalo/binding"
	"github.com/gobuffalo/envy"
	"github.com/gobuffalo/pop"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

// Field allows us to have generic functions for database fields
type field interface {
	SetPhotoName() error
	ValidPhoto() bool
	GetPhotoFileName() string
	GetPhotoStream() binding.File
}

// DB is a connection to your database to be used
// throughout your application.
var DB *pop.Connection

func init() {
	var err error
	env := envy.Get("GO_ENV", "development")
	DB, err = pop.Connect(env)
	if err != nil {
		log.Fatal(err)
	}
	pop.Debug = env == "development"
}

// BeforeSave is a callback that preprocess the struct before saving it to the database
func BeforeSaveFile(f field) error {
	f.SetPhotoName()
	return nil
}

// AfterSave is a callback that saves the file after the venue is saved to the database
func AfterSaveFile(f field) error {
	if !f.ValidPhoto() {
		logrus.Infof("Invalid photo file: %v", f)
		return nil
	}
	dir := filepath.Join(".", "public/uploads/photos")
	if err := os.MkdirAll(dir, 0755); err != nil {
		return errors.WithStack(err)
	}
	pf, err := os.Create(filepath.Join(dir, f.GetPhotoFileName()))
	if err != nil {
		return errors.WithStack(err)
	}
	defer pf.Close()
	_, err = io.Copy(pf, f.GetPhotoStream())
	return err
}
