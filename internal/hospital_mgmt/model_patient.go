/*
 * Hospital Management Api
 *
 * Hospital Management system for Web-In-Cloud
 *
 * API version: 1.0.0
 * Contact: xsabol@stuba.sk
 */

package hospital_mgmt

import "time"

type HospitalizationRecord struct {
	// Unique identifier of the hospitalization record
	Id string `json:"id"`

	// Description of the hospitalization
	Description string `json:"description"`
}

type Patient struct {
	// Unique identifier of the patient
	Id string `json:"id"`

	// First name of the patient
	FirstName string `json:"first_name"`

	// Last name of the patient
	LastName string `json:"last_name"`

	// Birth date of the patient
	BirthDate string `json:"birth_date"`

	// Gender of the patient (M/F/Other)
	Gender string `json:"gender"`

	// Phone number of the patient
	Phone string `json:"phone,omitempty"`

	// Email address of the patient
	Email string `json:"email,omitempty"`

	// List of hospitalization records
	HospitalizationRecords []HospitalizationRecord `json:"hospitalization_records,omitempty"`

	// Creation timestamp
	CreatedAt time.Time `json:"created_at"`

	// Last update timestamp
	UpdatedAt time.Time `json:"updated_at"`
} 