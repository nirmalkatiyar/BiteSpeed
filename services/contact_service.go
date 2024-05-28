package services

import (
	"github.com/nirmalkatiyar/bitespeed/models"
	"gorm.io/gorm"
	"time"
)

func IdentifyContact(db *gorm.DB, request models.IdentifyRequest) (response models.IdentifyResponse, err error) {
	var contacts []models.Contact
	if request.Email != nil && request.PhoneNumber != nil {
		db.Where("email=? OR phone_number=?", request.Email, request.PhoneNumber).Find(&contacts)
	} else if request.Email != nil {
		db.Where("email=?", request.Email).Find(&contacts)
	} else if request.PhoneNumber != nil {
		db.Where("phone_number=?", request.PhoneNumber).Find(&contacts)
	}

	// first time entry
	if len(contacts) == 0 {
		newContact := models.Contact{
			Email:          request.Email,
			PhoneNumber:    request.PhoneNumber,
			LinkPrecedence: "primary",
			CreatedAt:      time.Now(),
			UpdatedAt:      time.Now(),
		}
		db.Create(&newContact)
		return models.IdentifyResponse{
			Contact: models.ConsolidatedContact{
				PrimaryContactID: newContact.ID,
				Emails: []string{
					*request.Email,
				},
				PhoneNumbers: []string{
					*request.PhoneNumber,
				},
				SecondaryContactIDs: []uint{},
			},
		}, nil
	}

	primaryContact := findPrimaryContact(contacts)
	emails, phoneNumbers, secondaryContactIDs := consolidateContacts(contacts)

	if request.Email != nil && !contains(emails, *request.Email) {
		secondaryContact := models.Contact{
			Email:          request.Email,
			PhoneNumber:    primaryContact.PhoneNumber,
			LinkedID:       &primaryContact.ID,
			LinkPrecedence: "secondary",
			CreatedAt:      time.Now(),
			UpdatedAt:      time.Now(),
		}
		db.Create(&secondaryContact)
		secondaryContactIDs = append(secondaryContactIDs, secondaryContact.ID)
		emails = append(emails, *request.Email)
	}
	if request.PhoneNumber != nil && !contains(phoneNumbers, *request.PhoneNumber) {
		secondaryContact := models.Contact{
			Email:          primaryContact.Email,
			PhoneNumber:    request.PhoneNumber,
			LinkedID:       &primaryContact.ID,
			LinkPrecedence: "secondary",
			CreatedAt:      time.Now(),
			UpdatedAt:      time.Now(),
		}
		db.Create(&secondaryContact)
		secondaryContactIDs = append(secondaryContactIDs, secondaryContact.ID)
		phoneNumbers = append(phoneNumbers, *request.PhoneNumber)
	}

	finalResponse := models.IdentifyResponse{
		Contact: models.ConsolidatedContact{
			PrimaryContactID:    primaryContact.ID,
			Emails:              emails,
			PhoneNumbers:        phoneNumbers,
			SecondaryContactIDs: secondaryContactIDs,
		},
	}
	return finalResponse, nil
}

func findPrimaryContact(contacts []models.Contact) models.Contact {
	for _, contact := range contacts {
		if contact.LinkPrecedence == "primary" {
			return contact
		}
	}
	return contacts[0]
}
func consolidateContacts(contacts []models.Contact) ([]string, []string, []uint) {
	emails := make([]string, 0)
	phoneNumbers := make([]string, 0)
	secondaryContactIDs := make([]uint, 0)

	for _, contact := range contacts {
		if contact.Email != nil && !contains(emails, *contact.Email) {
			emails = append(emails, *contact.Email)
		}
		if contact.PhoneNumber != nil && !contains(phoneNumbers, *contact.PhoneNumber) {
			phoneNumbers = append(phoneNumbers, *contact.PhoneNumber)
		}
		if contact.LinkPrecedence == "secondary" {
			secondaryContactIDs = append(secondaryContactIDs, contact.ID)
		}
	}
	return emails, phoneNumbers, secondaryContactIDs
}

func contains(slice []string, item string) bool {
	for _, v := range slice {
		if v == item {
			return true
		}
	}
	return false
}
