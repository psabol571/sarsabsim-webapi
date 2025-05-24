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

type BedsAPI interface {

	// CreateBed Post /api/beds
	// Creates a new bed
	CreateBed(c *gin.Context)

	// GetBed Get /api/beds/:bedId
	// Gets details about a specific bed
	GetBed(c *gin.Context)

	// GetBeds Get /api/beds
	// Gets list of all beds
	GetBeds(c *gin.Context)

	// GetBedsByDepartment Get /api/departments/:departmentId/beds
	// Gets list of beds for a specific department
	GetBedsByDepartment(c *gin.Context)

	// UpdateBed Put /api/beds/:bedId
	// Updates specific bed
	UpdateBed(c *gin.Context)

	// DeleteBed Delete /api/beds/:bedId
	// Deletes specific bed
	DeleteBed(c *gin.Context)
} 