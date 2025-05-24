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

type DepartmentCapacity struct {
	// Maximum number of beds in the department
	MaximumBeds int `json:"maximum_beds"`

	// Actual number of beds available
	ActualBeds int `json:"actual_beds"`

	// Number of currently occupied beds
	OccupiedBeds int `json:"occupied_beds"`
}

type Department struct {
	// Unique identifier of the department
	Id string `json:"id"`

	// Name of the department
	Name string `json:"name"`

	// Description of the department
	Description string `json:"description"`

	// Floor number where the department is located
	Floor int `json:"floor"`

	// Capacity information for the department
	Capacity DepartmentCapacity `json:"capacity"`

	// Creation timestamp
	CreatedAt time.Time `json:"created_at"`

	// Last update timestamp
	UpdatedAt time.Time `json:"updated_at"`
} 