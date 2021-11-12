package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"

	"github.com/google/uuid"
)

// Book struct to describe book object.
type User struct {
	ID           uuid.UUID `db:"id" json:"id" validate:"required,uuid"`
	CreatedAt    time.Time `db:"created_at" json:"created_at"`
	UpdatedAt    time.Time `db:"updated_at" json:"updated_at"`
	FirstName    string    `db:"title" json:"firstname" validate:"required,lte=255"`
	LastName     string    `db:"author" json:"lastname" validate:"required,lte=255"`
	Status       int       `db:"status" json:"status" validate:"required,len=1"`
	Email        string    `db:"email" json:"email" validate:"required,email,lte=255"`
	PasswordHash string    `db:"password" json:"password,omitempty" validate:"required,lte=255"`
}

// Value make the User struct implement the driver.Valuer interface.
// This method simply returns the JSON-encoded representation of the struct.
func (user User) Value() (driver.Value, error) {
	return json.Marshal(user)
}

// Scan make the User struct implement the sql.Scanner interface.
// This method simply decodes a JSON-encoded value into the struct fields.
func (user *User) Scan(value interface{}) error {
	j, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(j, &user)
}
