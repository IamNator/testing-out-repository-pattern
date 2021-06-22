package models

import (
	"time"
)

//User ...
type User struct {
	ID        uint      `json:"id,omitempty" gorm:"primary_key"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Age       int       `json:"age"`
	CreatedAt time.Time `json:"created_at,omitempty" sql:"DEFAULT:current_timestamp"`
	UpdatedAt time.Time `json:"updated_at,omitempty" sql:"DEFAULT:current_timestamp"`
}
