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
	"net/http"

	"github.com/gin-gonic/gin"
)

// Route is the information for every URI.
type Route struct {
	// Name is the name of this Route.
	Name        string
	// Method is the string for the HTTP method. ex) GET, POST etc..
	Method      string
	// Pattern is the pattern of the URI.
	Pattern     string
	// HandlerFunc is the handler function of this route.
	HandlerFunc gin.HandlerFunc
}

// NewRouter returns a new router.
func NewRouter(handleFunctions ApiHandleFunctions) *gin.Engine {
	return NewRouterWithGinEngine(gin.Default(), handleFunctions)
}

// NewRouterWithGinEngine add routes to existing gin engine.
func NewRouterWithGinEngine(router *gin.Engine, handleFunctions ApiHandleFunctions) *gin.Engine {
	for _, route := range getRoutes(handleFunctions) {
		if route.HandlerFunc == nil {
			route.HandlerFunc = DefaultHandleFunc
		}
		switch route.Method {
		case http.MethodGet:
			router.GET(route.Pattern, route.HandlerFunc)
		case http.MethodPost:
			router.POST(route.Pattern, route.HandlerFunc)
		case http.MethodPut:
			router.PUT(route.Pattern, route.HandlerFunc)
		case http.MethodPatch:
			router.PATCH(route.Pattern, route.HandlerFunc)
		case http.MethodDelete:
			router.DELETE(route.Pattern, route.HandlerFunc)
		}
	}

	return router
}

// Default handler for not yet implemented routes
func DefaultHandleFunc(c *gin.Context) {
	c.String(http.StatusNotImplemented, "501 not implemented")
}

type ApiHandleFunctions struct {
	// Routes for the DepartmentsAPI part of the API
	DepartmentsAPI DepartmentsAPI
	// Routes for the BedsAPI part of the API
	BedsAPI BedsAPI
	// Routes for the PatientsAPI part of the API
	PatientsAPI PatientsAPI
}

func getRoutes(handleFunctions ApiHandleFunctions) []Route {
	return []Route{
		// Department routes
		{
			"CreateDepartment",
			http.MethodPost,
			"/api/departments",
			handleFunctions.DepartmentsAPI.CreateDepartment,
		},
		{
			"GetDepartment",
			http.MethodGet,
			"/api/departments/:departmentId",
			handleFunctions.DepartmentsAPI.GetDepartment,
		},
		{
			"GetDepartments",
			http.MethodGet,
			"/api/departments",
			handleFunctions.DepartmentsAPI.GetDepartments,
		},
		{
			"UpdateDepartment",
			http.MethodPut,
			"/api/departments/:departmentId",
			handleFunctions.DepartmentsAPI.UpdateDepartment,
		},
		{
			"DeleteDepartment",
			http.MethodDelete,
			"/api/departments/:departmentId",
			handleFunctions.DepartmentsAPI.DeleteDepartment,
		},
		// Bed routes
		{
			"CreateBed",
			http.MethodPost,
			"/api/beds",
			handleFunctions.BedsAPI.CreateBed,
		},
		{
			"GetBed",
			http.MethodGet,
			"/api/beds/:bedId",
			handleFunctions.BedsAPI.GetBed,
		},
		{
			"GetBeds",
			http.MethodGet,
			"/api/beds",
			handleFunctions.BedsAPI.GetBeds,
		},
		{
			"GetBedsByDepartment",
			http.MethodGet,
			"/api/departments/:departmentId/beds",
			handleFunctions.BedsAPI.GetBedsByDepartment,
		},
		{
			"UpdateBed",
			http.MethodPut,
			"/api/beds/:bedId",
			handleFunctions.BedsAPI.UpdateBed,
		},
		{
			"DeleteBed",
			http.MethodDelete,
			"/api/beds/:bedId",
			handleFunctions.BedsAPI.DeleteBed,
		},
		// Patient routes
		{
			"CreatePatient",
			http.MethodPost,
			"/api/patients",
			handleFunctions.PatientsAPI.CreatePatient,
		},
		{
			"GetPatient",
			http.MethodGet,
			"/api/patients/:patientId",
			handleFunctions.PatientsAPI.GetPatient,
		},
		{
			"GetPatients",
			http.MethodGet,
			"/api/patients",
			handleFunctions.PatientsAPI.GetPatients,
		},
		{
			"UpdatePatient",
			http.MethodPut,
			"/api/patients/:patientId",
			handleFunctions.PatientsAPI.UpdatePatient,
		},
		{
			"DeletePatient",
			http.MethodDelete,
			"/api/patients/:patientId",
			handleFunctions.PatientsAPI.DeletePatient,
		},
		// Hospitalization record routes
		{
			"AddHospitalizationRecord",
			http.MethodPost,
			"/api/patients/:patientId/hospitalizations",
			handleFunctions.PatientsAPI.AddHospitalizationRecord,
		},
		{
			"UpdateHospitalizationRecord",
			http.MethodPut,
			"/api/patients/:patientId/hospitalizations/:recordId",
			handleFunctions.PatientsAPI.UpdateHospitalizationRecord,
		},
		{
			"DeleteHospitalizationRecord",
			http.MethodDelete,
			"/api/patients/:patientId/hospitalizations/:recordId",
			handleFunctions.PatientsAPI.DeleteHospitalizationRecord,
		},
	}
} 