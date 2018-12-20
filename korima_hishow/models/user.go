package models

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/uuid"
	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID              uuid.UUID `json:"id" db:"id"`
	CreatedAt       time.Time `json:"created_at" db:"created_at"`
	UpdatedAt       time.Time `json:"updated_at" db:"updated_at"`
	Firstname       string    `json:"firstname" db:"firstname"`
	Username        string    `json:"username" db:"username" form:"username"`
	Middlename      string    `json:"middlename" db:"middlename"`
	Lastname        string    `json:"lastname" db:"lastname"`
	Mlastname       string    `json:"mlastname" db:"mlastname"`
	Email           string    `json:"email" db:"email"`
	Phone           string    `json:"phone" db:"phone"`
	Password        string    `json:"-" db:"-"`
	PasswordHash    string    `json:"-" db:"password_hash"`
	PasswordConfirm string    `json:"-" db:"-"`
	Admin           bool      `json:"admin" db:"admin"`
}

// Check for usernames

type UsernameNotTaken struct {
	Name  string
	Field string
	tx    *pop.Connection
}

type EmailNotTaken struct {
	Name  string
	Field string
	tx    *pop.Connection
}

// String is not required by pop and may be deleted
func (u User) String() string {
	ju, _ := json.Marshal(u)
	return string(ju)
}

// Users is not required by pop and may be deleted
type Users []User

// String is not required by pop and may be deleted
func (u Users) String() string {
	ju, _ := json.Marshal(u)
	return string(ju)
}

// Create and validates user
func (u *User) Create(tx *pop.Connection) (*validate.Errors, error) {
	u.Email = strings.ToLower(u.Email)
	u.Admin = false
	pwdHash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return validate.NewErrors(), err
	}
	u.PasswordHash = string(pwdHash)
	return tx.ValidateAndCreate(u)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (u *User) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.StringIsPresent{Field: u.Username, Name: "Username"},
		&UsernameNotTaken{Name: "Username", Field: u.Username, tx: tx},
		/*&validators.StringIsPresent{Field: u.Firstname, Name: "Firstname"},
		&validators.StringIsPresent{Field: u.Middlename, Name: "Middlename"},
		&validators.StringIsPresent{Field: u.Lastname, Name: "Lastname"},
		&validators.StringIsPresent{Field: u.Mlastname, Name: "Mlastname"},
		&validators.StringIsPresent{Field: u.Email, Name: "Email"},
		&validators.StringIsPresent{Field: u.Phone, Name: "Phone"},
		*/
		&validators.EmailIsPresent{Name: "Email", Field: u.Email},
		&EmailNotTaken{Name: "Email", Field: u.Email, tx: tx},
	), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (u *User) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (u *User) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// Check if username is already taken
func (v *UsernameNotTaken) IsValid(errors *validate.Errors) {
	query := v.tx.Where("username = ?", v.Field)
	qu := User{}
	err := query.First(&qu)
	if err == nil {
		// username already exists
		errors.Add(validators.GenerateKey(v.Name), fmt.Sprintf("The username %s is not available", v.Field))
	}
}

// Check if email is already taken
// TODO create a more generic interface for this validator
func (v *EmailNotTaken) IsValid(errors *validate.Errors) {
	query := v.tx.Where("email = ?", v.Field)
	qu := User{}
	err := query.First(&qu)
	if err == nil {
		// username already exists
		errors.Add(validators.GenerateKey(v.Name), fmt.Sprintf("The email %s is not available", v.Field))
	}
}
