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

type BedStatus struct {
	// Patient ID currently occupying the bed (if any)
	PatientId string `json:"patient_id,omitempty"`

	// Description of the bed status
	Description string `json:"description,omitempty"`
}

type Bed struct {
	// Unique identifier of the bed
	Id string `json:"id" bson:"_id"`

	// Department ID where the bed is located
	DepartmentId string `json:"department_id"`

	// Type of the bed (standard, ICU, etc.)
	BedType string `json:"bed_type"`

	// Quality rating of the bed (0.0 - 1.0)
	BedQuality float64 `json:"bed_quality"`

	// Current status of the bed
	Status BedStatus `json:"status"`

	// Creation timestamp
	CreatedAt time.Time `json:"created_at"`

	// Last update timestamp
	UpdatedAt time.Time `json:"updated_at"`
} 