package models

import "github.com/gofrs/uuid"

type Album struct {
	ID       int       `gorm:"primary_key" json:"id"`
	UserUUID uuid.UUID `gorm:"type:uuid;primary_key" json:"user_uuid"`
}
