package hospital_mgmt

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/psabol571/sarsabsim-webapi/internal/db_service"
)

type implBedsAPI struct {
}

func NewBedsAPI() BedsAPI {
	return &implBedsAPI{}
}

func (o *implBedsAPI) CreateBed(c *gin.Context) {
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

	db, ok := value.(db_service.DbService[Bed])
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

	bed := Bed{}
	err := c.BindJSON(&bed)
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

	if bed.Id == "" {
		bed.Id = uuid.New().String()
	}

	now := time.Now()
	bed.CreatedAt = now
	bed.UpdatedAt = now

	err = db.CreateDocument(c, bed.Id, &bed)

	switch err {
	case nil:
		c.JSON(
			http.StatusCreated,
			bed,
		)
	case db_service.ErrConflict:
		c.JSON(
			http.StatusConflict,
			gin.H{
				"status":  "Conflict",
				"message": "Bed already exists",
				"error":   err.Error(),
			},
		)
	default:
		c.JSON(
			http.StatusBadGateway,
			gin.H{
				"status":  "Bad Gateway",
				"message": "Failed to create bed in database",
				"error":   err.Error(),
			},
		)
	}
}

func (o *implBedsAPI) GetBed(c *gin.Context) {
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

	db, ok := value.(db_service.DbService[Bed])
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

	bedId := c.Param("bedId")
	bed, err := db.FindDocument(c, bedId)

	switch err {
	case nil:
		c.JSON(
			http.StatusOK,
			bed,
		)
	case db_service.ErrNotFound:
		c.JSON(
			http.StatusNotFound,
			gin.H{
				"status":  "Not Found",
				"message": "Bed not found",
				"error":   err.Error(),
			},
		)
	default:
		c.JSON(
			http.StatusBadGateway,
			gin.H{
				"status":  "Bad Gateway",
				"message": "Failed to find bed in database",
				"error":   err.Error(),
			})
	}
}

func (o *implBedsAPI) GetBeds(c *gin.Context) {
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

	db, ok := value.(db_service.DbService[Bed])
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

	beds, err := db.FindAllDocuments(c)
	switch err {
	case nil:
		c.JSON(
			http.StatusOK,
			beds,
		)
	default:
		c.JSON(
			http.StatusBadGateway,
			gin.H{
				"status":  "Bad Gateway",
				"message": "Failed to retrieve beds from database",
				"error":   err.Error(),
			})
	}
}

func (o *implBedsAPI) GetBedsByDepartment(c *gin.Context) {
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

	db, ok := value.(db_service.DbService[Bed])
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
	filter := map[string]interface{}{
		"departmentid": departmentId,
	}

	beds, err := db.FindDocumentsByFilter(c, filter)
	switch err {
	case nil:
		if beds == nil {
			beds = []*Bed{}
		}
		c.JSON(
			http.StatusOK,
			beds,
		)
	default:
		c.JSON(
			http.StatusBadGateway,
			gin.H{
				"status":  "Bad Gateway",
				"message": "Failed to retrieve beds by department from database",
				"error":   err.Error(),
			})
	}
}

func (o *implBedsAPI) UpdateBed(c *gin.Context) {
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

	db, ok := value.(db_service.DbService[Bed])
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

	bedId := c.Param("bedId")

	// First check if bed exists
	existingBed, err := db.FindDocument(c, bedId)
	if err != nil {
		switch err {
		case db_service.ErrNotFound:
			c.JSON(
				http.StatusNotFound,
				gin.H{
					"status":  "Not Found",
					"message": "Bed not found",
					"error":   err.Error(),
				},
			)
		default:
			c.JSON(
				http.StatusBadGateway,
				gin.H{
					"status":  "Bad Gateway",
					"message": "Failed to find bed in database",
					"error":   err.Error(),
				})
		}
		return
	}

	updatedBed := Bed{}
	err = c.BindJSON(&updatedBed)
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
	updatedBed.Id = bedId
	updatedBed.CreatedAt = existingBed.CreatedAt
	updatedBed.UpdatedAt = time.Now()

	err = db.UpdateDocument(c, bedId, &updatedBed)

	switch err {
	case nil:
		c.JSON(
			http.StatusOK,
			updatedBed,
		)
	case db_service.ErrNotFound:
		c.JSON(
			http.StatusNotFound,
			gin.H{
				"status":  "Not Found",
				"message": "Bed not found",
				"error":   err.Error(),
			},
		)
	default:
		c.JSON(
			http.StatusBadGateway,
			gin.H{
				"status":  "Bad Gateway",
				"message": "Failed to update bed in database",
				"error":   err.Error(),
			})
	}
}

func (o *implBedsAPI) DeleteBed(c *gin.Context) {
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

	db, ok := value.(db_service.DbService[Bed])
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

	bedId := c.Param("bedId")
	err := db.DeleteDocument(c, bedId)

	switch err {
	case nil:
		c.AbortWithStatus(http.StatusNoContent)
	case db_service.ErrNotFound:
		c.JSON(
			http.StatusNotFound,
			gin.H{
				"status":  "Not Found",
				"message": "Bed not found",
				"error":   err.Error(),
			},
		)
	default:
		c.JSON(
			http.StatusBadGateway,
			gin.H{
				"status":  "Bad Gateway",
				"message": "Failed to delete bed from database",
				"error":   err.Error(),
			})
	}
} 