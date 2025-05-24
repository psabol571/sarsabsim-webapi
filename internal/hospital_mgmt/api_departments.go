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

type DepartmentsAPI interface {

	// CreateDepartment Post /api/departments
	// Creates a new department
	CreateDepartment(c *gin.Context)

	// GetDepartment Get /api/departments/:departmentId
	// Gets details about a specific department
	GetDepartment(c *gin.Context)

	// GetDepartments Get /api/departments
	// Gets list of all departments
	GetDepartments(c *gin.Context)

	// UpdateDepartment Put /api/departments/:departmentId
	// Updates specific department
	UpdateDepartment(c *gin.Context)

	// DeleteDepartment Delete /api/departments/:departmentId
	// Deletes specific department
	DeleteDepartment(c *gin.Context)
} 