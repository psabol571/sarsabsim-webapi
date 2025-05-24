package hospital_mgmt

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/psabol571/sarsabsim-webapi/internal/db_service"
)

type implPatientsAPI struct {
}

func NewPatientsAPI() PatientsAPI {
	return &implPatientsAPI{}
}

func (o *implPatientsAPI) CreatePatient(c *gin.Context) {
	value, exists := c.Get("db_service")
	if !exists {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"status":  "Internal Server Error",
				"message": "db_service not found",
				"error":   "db_service not found",
			})
		return
	}

	db, ok := value.(db_service.DbService[Patient])
	if !ok {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"status":  "Internal Server Error",
				"message": "db_service context is not of type db_service.DbService",
				"error":   "cannot cast db_service context to db_service.DbService",
			})
		return
	}

	patient := Patient{}
	err := c.BindJSON(&patient)
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"status":  "Bad Request",
				"message": "Invalid request body",
				"error":   err.Error(),
			})
		return
	}

	if patient.Id == "" {
		patient.Id = uuid.New().String()
	}

	now := time.Now()
	patient.CreatedAt = now
	patient.UpdatedAt = now

	err = db.CreateDocument(c, patient.Id, &patient)

	switch err {
	case nil:
		c.JSON(
			http.StatusCreated,
			patient,
		)
	case db_service.ErrConflict:
		c.JSON(
			http.StatusConflict,
			gin.H{
				"status":  "Conflict",
				"message": "Patient already exists",
				"error":   err.Error(),
			},
		)
	default:
		c.JSON(
			http.StatusBadGateway,
			gin.H{
				"status":  "Bad Gateway",
				"message": "Failed to create patient in database",
				"error":   err.Error(),
			},
		)
	}
}

func (o *implPatientsAPI) GetPatient(c *gin.Context) {
	value, exists := c.Get("db_service")
	if !exists {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"status":  "Internal Server Error",
				"message": "db_service not found",
				"error":   "db_service not found",
			})
		return
	}

	db, ok := value.(db_service.DbService[Patient])
	if !ok {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"status":  "Internal Server Error",
				"message": "db_service context is not of type db_service.DbService",
				"error":   "cannot cast db_service context to db_service.DbService",
			})
		return
	}

	patientId := c.Param("patientId")
	patient, err := db.FindDocument(c, patientId)

	switch err {
	case nil:
		c.JSON(
			http.StatusOK,
			patient,
		)
	case db_service.ErrNotFound:
		c.JSON(
			http.StatusNotFound,
			gin.H{
				"status":  "Not Found",
				"message": "Patient not found",
				"error":   err.Error(),
			},
		)
	default:
		c.JSON(
			http.StatusBadGateway,
			gin.H{
				"status":  "Bad Gateway",
				"message": "Failed to find patient in database",
				"error":   err.Error(),
			})
	}
}

func (o *implPatientsAPI) GetPatients(c *gin.Context) {
	value, exists := c.Get("db_service")
	if !exists {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"status":  "Internal Server Error",
				"message": "db_service not found",
				"error":   "db_service not found",
			})
		return
	}

	db, ok := value.(db_service.DbService[Patient])
	if !ok {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"status":  "Internal Server Error",
				"message": "db_service context is not of type db_service.DbService",
				"error":   "cannot cast db_service context to db_service.DbService",
			})
		return
	}

	patients, err := db.FindAllDocuments(c)
	switch err {
	case nil:
		c.JSON(
			http.StatusOK,
			patients,
		)
	default:
		c.JSON(
			http.StatusBadGateway,
			gin.H{
				"status":  "Bad Gateway",
				"message": "Failed to retrieve patients from database",
				"error":   err.Error(),
			})
	}
}

func (o *implPatientsAPI) UpdatePatient(c *gin.Context) {
	value, exists := c.Get("db_service")
	if !exists {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"status":  "Internal Server Error",
				"message": "db_service not found",
				"error":   "db_service not found",
			})
		return
	}

	db, ok := value.(db_service.DbService[Patient])
	if !ok {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"status":  "Internal Server Error",
				"message": "db_service context is not of type db_service.DbService",
				"error":   "cannot cast db_service context to db_service.DbService",
			})
		return
	}

	patientId := c.Param("patientId")

	// First check if patient exists
	existingPatient, err := db.FindDocument(c, patientId)
	if err != nil {
		switch err {
		case db_service.ErrNotFound:
			c.JSON(
				http.StatusNotFound,
				gin.H{
					"status":  "Not Found",
					"message": "Patient not found",
					"error":   err.Error(),
				},
			)
		default:
			c.JSON(
				http.StatusBadGateway,
				gin.H{
					"status":  "Bad Gateway",
					"message": "Failed to find patient in database",
					"error":   err.Error(),
				})
		}
		return
	}

	updatedPatient := Patient{}
	err = c.BindJSON(&updatedPatient)
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"status":  "Bad Request",
				"message": "Invalid request body",
				"error":   err.Error(),
			})
		return
	}

	// Preserve certain fields
	updatedPatient.Id = patientId
	updatedPatient.CreatedAt = existingPatient.CreatedAt
	updatedPatient.UpdatedAt = time.Now()

	err = db.UpdateDocument(c, patientId, &updatedPatient)

	switch err {
	case nil:
		c.JSON(
			http.StatusOK,
			updatedPatient,
		)
	case db_service.ErrNotFound:
		c.JSON(
			http.StatusNotFound,
			gin.H{
				"status":  "Not Found",
				"message": "Patient not found",
				"error":   err.Error(),
			},
		)
	default:
		c.JSON(
			http.StatusBadGateway,
			gin.H{
				"status":  "Bad Gateway",
				"message": "Failed to update patient in database",
				"error":   err.Error(),
			})
	}
}

func (o *implPatientsAPI) DeletePatient(c *gin.Context) {
	value, exists := c.Get("db_service")
	if !exists {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"status":  "Internal Server Error",
				"message": "db_service not found",
				"error":   "db_service not found",
			})
		return
	}

	db, ok := value.(db_service.DbService[Patient])
	if !ok {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"status":  "Internal Server Error",
				"message": "db_service context is not of type db_service.DbService",
				"error":   "cannot cast db_service context to db_service.DbService",
			})
		return
	}

	patientId := c.Param("patientId")
	err := db.DeleteDocument(c, patientId)

	switch err {
	case nil:
		c.AbortWithStatus(http.StatusNoContent)
	case db_service.ErrNotFound:
		c.JSON(
			http.StatusNotFound,
			gin.H{
				"status":  "Not Found",
				"message": "Patient not found",
				"error":   err.Error(),
			},
		)
	default:
		c.JSON(
			http.StatusBadGateway,
			gin.H{
				"status":  "Bad Gateway",
				"message": "Failed to delete patient from database",
				"error":   err.Error(),
			})
	}
}

func (o *implPatientsAPI) AddHospitalizationRecord(c *gin.Context) {
	value, exists := c.Get("db_service")
	if !exists {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"status":  "Internal Server Error",
				"message": "db_service not found",
				"error":   "db_service not found",
			})
		return
	}

	db, ok := value.(db_service.DbService[Patient])
	if !ok {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"status":  "Internal Server Error",
				"message": "db_service context is not of type db_service.DbService",
				"error":   "cannot cast db_service context to db_service.DbService",
			})
		return
	}

	patientId := c.Param("patientId")

	// First find the existing patient
	patient, err := db.FindDocument(c, patientId)
	if err != nil {
		switch err {
		case db_service.ErrNotFound:
			c.JSON(
				http.StatusNotFound,
				gin.H{
					"status":  "Not Found",
					"message": "Patient not found",
					"error":   err.Error(),
				},
			)
		default:
			c.JSON(
				http.StatusBadGateway,
				gin.H{
					"status":  "Bad Gateway",
					"message": "Failed to find patient in database",
					"error":   err.Error(),
				})
		}
		return
	}

	newRecord := HospitalizationRecord{}
	err = c.BindJSON(&newRecord)
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"status":  "Bad Request",
				"message": "Invalid request body",
				"error":   err.Error(),
			})
		return
	}

	if newRecord.Id == "" {
		newRecord.Id = uuid.New().String()
	}

	// Add the new record to the patient's list
	patient.HospitalizationRecords = append(patient.HospitalizationRecords, newRecord)
	patient.UpdatedAt = time.Now()

	err = db.UpdateDocument(c, patientId, patient)

	switch err {
	case nil:
		c.JSON(
			http.StatusCreated,
			newRecord,
		)
	case db_service.ErrNotFound:
		c.JSON(
			http.StatusNotFound,
			gin.H{
				"status":  "Not Found",
				"message": "Patient not found",
				"error":   err.Error(),
			},
		)
	default:
		c.JSON(
			http.StatusBadGateway,
			gin.H{
				"status":  "Bad Gateway",
				"message": "Failed to add hospitalization record",
				"error":   err.Error(),
			})
	}
}

func (o *implPatientsAPI) UpdateHospitalizationRecord(c *gin.Context) {
	value, exists := c.Get("db_service")
	if !exists {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"status":  "Internal Server Error",
				"message": "db_service not found",
				"error":   "db_service not found",
			})
		return
	}

	db, ok := value.(db_service.DbService[Patient])
	if !ok {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"status":  "Internal Server Error",
				"message": "db_service context is not of type db_service.DbService",
				"error":   "cannot cast db_service context to db_service.DbService",
			})
		return
	}

	patientId := c.Param("patientId")
	recordId := c.Param("recordId")

	// First find the existing patient
	patient, err := db.FindDocument(c, patientId)
	if err != nil {
		switch err {
		case db_service.ErrNotFound:
			c.JSON(
				http.StatusNotFound,
				gin.H{
					"status":  "Not Found",
					"message": "Patient not found",
					"error":   err.Error(),
				},
			)
		default:
			c.JSON(
				http.StatusBadGateway,
				gin.H{
					"status":  "Bad Gateway",
					"message": "Failed to find patient in database",
					"error":   err.Error(),
				})
		}
		return
	}

	updatedRecord := HospitalizationRecord{}
	err = c.BindJSON(&updatedRecord)
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"status":  "Bad Request",
				"message": "Invalid request body",
				"error":   err.Error(),
			})
		return
	}

	// Find and update the specific record
	recordFound := false
	for i, record := range patient.HospitalizationRecords {
		if record.Id == recordId {
			updatedRecord.Id = recordId
			patient.HospitalizationRecords[i] = updatedRecord
			recordFound = true
			break
		}
	}

	if !recordFound {
		c.JSON(
			http.StatusNotFound,
			gin.H{
				"status":  "Not Found",
				"message": "Hospitalization record not found",
				"error":   "record with specified ID not found",
			},
		)
		return
	}

	patient.UpdatedAt = time.Now()

	err = db.UpdateDocument(c, patientId, patient)

	switch err {
	case nil:
		c.JSON(
			http.StatusOK,
			updatedRecord,
		)
	case db_service.ErrNotFound:
		c.JSON(
			http.StatusNotFound,
			gin.H{
				"status":  "Not Found",
				"message": "Patient not found",
				"error":   err.Error(),
			},
		)
	default:
		c.JSON(
			http.StatusBadGateway,
			gin.H{
				"status":  "Bad Gateway",
				"message": "Failed to update hospitalization record",
				"error":   err.Error(),
			})
	}
}

func (o *implPatientsAPI) DeleteHospitalizationRecord(c *gin.Context) {
	value, exists := c.Get("db_service")
	if !exists {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"status":  "Internal Server Error",
				"message": "db_service not found",
				"error":   "db_service not found",
			})
		return
	}

	db, ok := value.(db_service.DbService[Patient])
	if !ok {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"status":  "Internal Server Error",
				"message": "db_service context is not of type db_service.DbService",
				"error":   "cannot cast db_service context to db_service.DbService",
			})
		return
	}

	patientId := c.Param("patientId")
	recordId := c.Param("recordId")

	// First find the existing patient
	patient, err := db.FindDocument(c, patientId)
	if err != nil {
		switch err {
		case db_service.ErrNotFound:
			c.JSON(
				http.StatusNotFound,
				gin.H{
					"status":  "Not Found",
					"message": "Patient not found",
					"error":   err.Error(),
				},
			)
		default:
			c.JSON(
				http.StatusBadGateway,
				gin.H{
					"status":  "Bad Gateway",
					"message": "Failed to find patient in database",
					"error":   err.Error(),
				})
		}
		return
	}

	// Find and remove the specific record
	recordFound := false
	for i, record := range patient.HospitalizationRecords {
		if record.Id == recordId {
			// Remove the record from the slice
			patient.HospitalizationRecords = append(
				patient.HospitalizationRecords[:i],
				patient.HospitalizationRecords[i+1:]...,
			)
			recordFound = true
			break
		}
	}

	if !recordFound {
		c.JSON(
			http.StatusNotFound,
			gin.H{
				"status":  "Not Found",
				"message": "Hospitalization record not found",
				"error":   "record with specified ID not found",
			},
		)
		return
	}

	patient.UpdatedAt = time.Now()

	err = db.UpdateDocument(c, patientId, patient)

	switch err {
	case nil:
		c.AbortWithStatus(http.StatusNoContent)
	case db_service.ErrNotFound:
		c.JSON(
			http.StatusNotFound,
			gin.H{
				"status":  "Not Found",
				"message": "Patient not found",
				"error":   err.Error(),
			},
		)
	default:
		c.JSON(
			http.StatusBadGateway,
			gin.H{
				"status":  "Bad Gateway",
				"message": "Failed to delete hospitalization record",
				"error":   err.Error(),
			})
	}
} 