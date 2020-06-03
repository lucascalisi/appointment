// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"log"
	"net/http"

	errors "github.com/go-openapi/errors"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"

	"github.com/appointment/config"
	"github.com/appointment/http/restapi/operations"
	"github.com/appointment/http/restapi/operations/patients"
	"github.com/appointment/http/restapi/operations/professionals"
	"github.com/appointment/mysql"
)

//go:generate swagger generate server --target ../../http --name APPointmentBackend --spec ../swagger/swagger.yml

func configureFlags(api *operations.HealthyCalendarAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.HealthyCalendarAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	cfg, err := config.GetConfig()
	if err != nil {
		log.Fatalf("could not get configuration from environment: %v", err)
	}

	db, err := mysql.NewDB(cfg.DBCfg.DBUser, cfg.DBCfg.DBPassword, cfg.DBCfg.DBHost,
		cfg.DBCfg.DBPort, cfg.DBCfg.DBName, cfg.Debug)
	if err != nil {
		log.Fatalf("could not establish connection to the database: %v", err)
	}

	mysql.SetConnectionParameters(db)

	if api.PatientsCancelAppoinmentForPatientHandler == nil {
		api.PatientsCancelAppoinmentForPatientHandler = patients.CancelAppoinmentForPatientHandlerFunc(func(params patients.CancelAppoinmentForPatientParams) middleware.Responder {
			return middleware.NotImplemented("operation patients.CancelAppoinmentForPatient has not yet been implemented")
		})
	}
	if api.ProfessionalsCancelAppointmentProfessionalHandler == nil {
		api.ProfessionalsCancelAppointmentProfessionalHandler = professionals.CancelAppointmentProfessionalHandlerFunc(func(params professionals.CancelAppointmentProfessionalParams) middleware.Responder {
			return middleware.NotImplemented("operation professionals.CancelAppointmentProfessional has not yet been implemented")
		})
	}
	if api.PatientsConfirmAppointmentForPatientHandler == nil {
		api.PatientsConfirmAppointmentForPatientHandler = patients.ConfirmAppointmentForPatientHandlerFunc(func(params patients.ConfirmAppointmentForPatientParams) middleware.Responder {
			return middleware.NotImplemented("operation patients.ConfirmAppointmentForPatient has not yet been implemented")
		})
	}
	api.ProfessionalsGetAppointmentByProfessionalAppointmentIDHandler = getProfessionalAppointment(db)

	api.PatientsGetAppointmentsByPatientHandler = getPatientAppointments(db)

	api.ProfessionalsGetAppointmentsByprofessionalHandler = getProfessionalAppointments(db)

	api.PatientsGetPatientbyIDHandler = getPatient(db)

	if api.ProfessionalsGetProfesionalScheduleBySpecialtyHandler == nil {
		api.ProfessionalsGetProfesionalScheduleBySpecialtyHandler = professionals.GetProfesionalScheduleBySpecialtyHandlerFunc(func(params professionals.GetProfesionalScheduleBySpecialtyParams) middleware.Responder {
			return middleware.NotImplemented("operation professionals.GetProfesionalScheduleBySpecialty has not yet been implemented")
		})
	}

	api.ProfessionalsGetProfessionalbyIDHandler = getProfessional(db)

	api.ProfessionalsGetSpecialtiesByProfessionalHandler = getProfessionalSpecialties(db)

	api.LoginLoginHandler = authenticate(db)

	api.AppointmentsSearchAppointmentHandler = searchAppointments(db)

	api.SpecialtiesSearchSpecialtyHandler = searchSpecialties(db)

	if api.ProfessionalsSetProfesionalScheduleBySpecialtyHandler == nil {
		api.ProfessionalsSetProfesionalScheduleBySpecialtyHandler = professionals.SetProfesionalScheduleBySpecialtyHandlerFunc(func(params professionals.SetProfesionalScheduleBySpecialtyParams) middleware.Responder {
			return middleware.NotImplemented("operation professionals.SetProfesionalScheduleBySpecialty has not yet been implemented")
		})
	}

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix"
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
