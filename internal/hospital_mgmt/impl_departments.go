package hospital_mgmt

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/psabol571/sarsabsim-webapi/internal/db_service"
)

type implDepartmentsAPI struct {
}

func NewDepartmentsAPI() DepartmentsAPI {
	return &implDepartmentsAPI{}
}

func (o *implDepartmentsAPI) CreateDepartment(c *gin.Context) {
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

	db, ok := value.(db_service.DbService[Department])
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

	department := Department{}
	err := c.BindJSON(&department)
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

	if department.Id == "" {
		department.Id = uuid.New().String()
	}

	now := time.Now()
	department.CreatedAt = now
	department.UpdatedAt = now

	err = db.CreateDocument(c, department.Id, &department)

	switch err {
	case nil:
		c.JSON(
			http.StatusCreated,
			department,
		)
	case db_service.ErrConflict:
		c.JSON(
			http.StatusConflict,
			gin.H{
				"status":  "Conflict",
				"message": "Department already exists",
				"error":   err.Error(),
			},
		)
	default:
		c.JSON(
			http.StatusBadGateway,
			gin.H{
				"status":  "Bad Gateway",
				"message": "Failed to create department in database",
				"error":   err.Error(),
			},
		)
	}
}

func (o *implDepartmentsAPI) GetDepartment(c *gin.Context) {
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

	db, ok := value.(db_service.DbService[Department])
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

	departmentId := c.Param("departmentId")
	department, err := db.FindDocument(c, departmentId)

	switch err {
	case nil:
		c.JSON(
			http.StatusOK,
			department,
		)
	case db_service.ErrNotFound:
		c.JSON(
			http.StatusNotFound,
			gin.H{
				"status":  "Not Found",
				"message": "Department not found",
				"error":   err.Error(),
			},
		)
	default:
		c.JSON(
			http.StatusBadGateway,
			gin.H{
				"status":  "Bad Gateway",
				"message": "Failed to find department in database",
				"error":   err.Error(),
			})
	}
}

func (o *implDepartmentsAPI) GetDepartments(c *gin.Context) {
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

	db, ok := value.(db_service.DbService[Department])
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

	departments, err := db.FindAllDocuments(c)
	switch err {
	case nil:
		c.JSON(
			http.StatusOK,
			departments,
		)
	default:
		c.JSON(
			http.StatusBadGateway,
			gin.H{
				"status":  "Bad Gateway",
				"message": "Failed to retrieve departments from database",
				"error":   err.Error(),
			})
	}
}

func (o *implDepartmentsAPI) UpdateDepartment(c *gin.Context) {
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

	db, ok := value.(db_service.DbService[Department])
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

	departmentId := c.Param("departmentId")

	// First check if department exists
	existingDepartment, err := db.FindDocument(c, departmentId)
	if err != nil {
		switch err {
		case db_service.ErrNotFound:
			c.JSON(
				http.StatusNotFound,
				gin.H{
					"status":  "Not Found",
					"message": "Department not found",
					"error":   err.Error(),
				},
			)
		default:
			c.JSON(
				http.StatusBadGateway,
				gin.H{
					"status":  "Bad Gateway",
					"message": "Failed to find department in database",
					"error":   err.Error(),
				})
		}
		return
	}

	updatedDepartment := Department{}
	err = c.BindJSON(&updatedDepartment)
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
	updatedDepartment.Id = departmentId
	updatedDepartment.CreatedAt = existingDepartment.CreatedAt
	updatedDepartment.UpdatedAt = time.Now()

	err = db.UpdateDocument(c, departmentId, &updatedDepartment)

	switch err {
	case nil:
		c.JSON(
			http.StatusOK,
			updatedDepartment,
		)
	case db_service.ErrNotFound:
		c.JSON(
			http.StatusNotFound,
			gin.H{
				"status":  "Not Found",
				"message": "Department not found",
				"error":   err.Error(),
			},
		)
	default:
		c.JSON(
			http.StatusBadGateway,
			gin.H{
				"status":  "Bad Gateway",
				"message": "Failed to update department in database",
				"error":   err.Error(),
			})
	}
}

func (o *implDepartmentsAPI) DeleteDepartment(c *gin.Context) {
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

	db, ok := value.(db_service.DbService[Department])
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

	departmentId := c.Param("departmentId")
	err := db.DeleteDocument(c, departmentId)

	switch err {
	case nil:
		c.AbortWithStatus(http.StatusNoContent)
	case db_service.ErrNotFound:
		c.JSON(
			http.StatusNotFound,
			gin.H{
				"status":  "Not Found",
				"message": "Department not found",
				"error":   err.Error(),
			},
		)
	default:
		c.JSON(
			http.StatusBadGateway,
			gin.H{
				"status":  "Bad Gateway",
				"message": "Failed to delete department from database",
				"error":   err.Error(),
			})
	}
} 