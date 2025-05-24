package main

import (
	// "log"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/psabol571/sarsabsim-webapi/api"
	"github.com/psabol571/sarsabsim-webapi/internal/ambulance_wl"
	"github.com/psabol571/sarsabsim-webapi/internal/hospital_mgmt"

	"context"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/psabol571/sarsabsim-webapi/internal/db_service"

	"github.com/rs/zerolog"
  	"github.com/rs/zerolog/log"


	  "go.opentelemetry.io/contrib/exporters/autoexport"
	  "go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
	  "go.opentelemetry.io/otel"
	  "go.opentelemetry.io/otel/propagation"
	  tracesdk "go.opentelemetry.io/otel/sdk/trace"
)

func main() {
	output := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: zerolog.TimeFormatUnix}
	log.Logger = zerolog.New(output).With().
	  Str("service", "ambulance-wl-list").
	  Timestamp().
	  Caller().
	  Logger()
  
	logLevelStr := os.Getenv("LOG_LEVEL")
	defaultLevel := zerolog.InfoLevel
	level, err := zerolog.ParseLevel(strings.ToLower(logLevelStr))
	if err != nil {
	  log.Warn().Str("LOG_LEVEL", logLevelStr).Msgf("Invalid log level, using default: %s", defaultLevel)
	  level = defaultLevel
	}
	// Set the global log level
	zerolog.SetGlobalLevel(level)

	  // initialize trace exporter
	  ctx, cancel := context.WithCancel(context.Background())
	  defer cancel()
	  traceExporter, err := autoexport.NewSpanExporter(ctx)
	  if err != nil {
		log.Fatal().Err(err).Msg("Failed to initialize trace exporter")
	  }
	
	  traceProvider := tracesdk.NewTracerProvider(tracesdk.WithBatcher(traceExporter))
	  otel.SetTracerProvider(traceProvider)
	  otel.SetTextMapPropagator(propagation.TraceContext{})
	  defer  traceProvider.Shutdown(ctx)
  
	log.Info().Msg("Server started")

	// log.Printf("Server started")
	port := os.Getenv("AMBULANCE_API_PORT")
	if port == "" {
		port = "8080"
	}
	environment := os.Getenv("AMBULANCE_API_ENVIRONMENT")
	if !strings.EqualFold(environment, "production") { // case insensitive comparison
		gin.SetMode(gin.DebugMode)
	}
	engine := gin.New()
	engine.Use(gin.Recovery())
	engine.Use(otelgin.Middleware("ambulance-webapi"))

	corsMiddleware := cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "PUT", "POST", "DELETE", "PATCH"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"},
		ExposeHeaders:    []string{""},
		AllowCredentials: false,
		MaxAge:           12 * time.Hour,
	})
	engine.Use(corsMiddleware)

	// setup context update middleware
	ambulanceDbService := db_service.NewMongoService[ambulance_wl.Ambulance](db_service.MongoServiceConfig{})
	defer ambulanceDbService.Disconnect(context.Background())
	
	// Hospital management db services
	departmentDbService := db_service.NewMongoService[hospital_mgmt.Department](db_service.MongoServiceConfig{})
	defer departmentDbService.Disconnect(context.Background())
	
	bedDbService := db_service.NewMongoService[hospital_mgmt.Bed](db_service.MongoServiceConfig{})
	defer bedDbService.Disconnect(context.Background())
	
	patientDbService := db_service.NewMongoService[hospital_mgmt.Patient](db_service.MongoServiceConfig{})
	defer patientDbService.Disconnect(context.Background())
	
	engine.Use(func(ctx *gin.Context) {
		// Set appropriate db service based on the request path
		path := ctx.Request.URL.Path
		if strings.HasPrefix(path, "/api/departments") {
			ctx.Set("db_service", departmentDbService)
		} else if strings.HasPrefix(path, "/api/beds") {
			ctx.Set("db_service", bedDbService)
		} else if strings.HasPrefix(path, "/api/patients") {
			ctx.Set("db_service", patientDbService)
		} else {
			// Default to ambulance db service for existing ambulance endpoints
			ctx.Set("db_service", ambulanceDbService)
		}
		ctx.Next()
	})

	// request routings
	handleFunctions := &ambulance_wl.ApiHandleFunctions{
		AmbulanceConditionsAPI:  ambulance_wl.NewAmbulanceConditionsApi(),
		AmbulanceWaitingListAPI: ambulance_wl.NewAmbulanceWaitingListApi(),
		AmbulancesAPI:           ambulance_wl.NewAmbulancesApi(),
	}
	ambulance_wl.NewRouterWithGinEngine(engine, *handleFunctions)

	// hospital management routings
	hospitalHandleFunctions := &hospital_mgmt.ApiHandleFunctions{
		DepartmentsAPI: hospital_mgmt.NewDepartmentsAPI(),
		BedsAPI:        hospital_mgmt.NewBedsAPI(),
		PatientsAPI:    hospital_mgmt.NewPatientsAPI(),
	}
	hospital_mgmt.NewRouterWithGinEngine(engine, *hospitalHandleFunctions)

	engine.GET("/openapi", api.HandleOpenApi)
	engine.Run(":" + port)
}
