package models

import (
	"gorm.io/gorm"
	"time"
)

type Contact struct {
	ID             uint           `gorm:"primaryKey" json:"id"`
	PhoneNumber    *string        `json:"phoneNumber,omitempty"`
	Email          *string        `json:"email,omitempty"`
	LinkedID       *uint          `json:"linkedId,omitempty"`
	LinkPrecedence string         `json:"linkPrecedence"`
	CreatedAt      time.Time      `json:"createdAt"`
	UpdatedAt      time.Time      `json:"updatedAt"`
	DeletedAt      gorm.DeletedAt `gorm:"index" json:"deletedAt,omitempty"`
}

type IdentifyRequest struct {
	Email       *string `json:"email,omitempty"`
	PhoneNumber *string `json:"phoneNumber,omitempty"`
}

type IdentifyResponse struct {
	Contact ConsolidatedContact `json:"contact"`
}

type ConsolidatedContact struct {
	PrimaryContactID    uint     `json:"primaryContactId"`
	Emails              []string `json:"emails"`
	PhoneNumbers        []string `json:"phoneNumbers"`
	SecondaryContactIDs []uint   `json:"secondaryContactIds"`
}
