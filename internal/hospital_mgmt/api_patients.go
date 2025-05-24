/*
 * Hospital Management Api
 *
 * Hospital Management system for Web-In-Cloud
 *
 * API version: 1.0.0
 * Contact: xsabol@stuba.sk
 */

package hospital_mgmt

import (
	"github.com/gin-gonic/gin"
)

type PatientsAPI interface {

	// CreatePatient Post /api/patients
	// Creates a new patient
	CreatePatient(c *gin.Context)

	// GetPatient Get /api/patients/:patientId
	// Gets details about a specific patient
	GetPatient(c *gin.Context)

	// GetPatients Get /api/patients
	// Gets list of all patients
	GetPatients(c *gin.Context)

	// UpdatePatient Put /api/patients/:patientId
	// Updates specific patient
	UpdatePatient(c *gin.Context)

	// DeletePatient Delete /api/patients/:patientId
	// Deletes specific patient
	DeletePatient(c *gin.Context)

	// AddHospitalizationRecord Post /api/patients/:patientId/hospitalizations
	// Adds a new hospitalization record to a patient
	AddHospitalizationRecord(c *gin.Context)

	// UpdateHospitalizationRecord Put /api/patients/:patientId/hospitalizations/:recordId
	// Updates a specific hospitalization record
	UpdateHospitalizationRecord(c *gin.Context)

	// DeleteHospitalizationRecord Delete /api/patients/:patientId/hospitalizations/:recordId
	// Deletes a specific hospitalization record
	DeleteHospitalizationRecord(c *gin.Context)
} 