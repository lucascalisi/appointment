// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"log"
	"net/http"

	errors "github.com/go-openapi/errors"
	runtime "github.com/go-openapi/runtime"

	"github.com/appointment/config"
	"github.com/appointment/http/restapi/operations"
	"github.com/appointment/mysql"
	"github.com/appointment/utils"
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

	emailSender, err := utils.NewEmailSender(cfg.EmailSenderCfg.User, cfg.EmailSenderCfg.MySecretPassword,
		cfg.EmailSenderCfg.SMTPServer, cfg.EmailSenderCfg.Port)
	if err != nil {
		log.Fatalf("could not instance new email sender: %v", err)
	}

	mysql.SetConnectionParameters(db)

	api.PatientsCancelAppoinmentForPatientHandler = patientCancelAppointment(db)

	api.QueueGetQueueHandler = getQueue(db)
	api.QueueAddPatientToQueueHandler = addItem(db)

	api.ProfessionalsCancelAppointmentProfessionalHandler = professionalCancelAppointment(db, emailSender)

	api.PatientsConfirmAppointmentForPatientHandler = patientConfirmAppointment(db)

	api.PatientsRequestAppointmentForPatientHandler = patientRequestAppointment(db, cfg.AppointmentDuration)

	api.ProfessionalsGetAppointmentByProfessionalAppointmentIDHandler = getProfessionalAppointment(db)

	api.PatientsGetAppointmentsByPatientHandler = getPatientAppointments(db)

	api.ProfessionalsGetAppointmentsByprofessionalHandler = getProfessionalAppointments(db)

	api.PatientsGetPatientbyIDHandler = getPatient(db)

	api.ProfessionalsGetProfesionalScheduleBySpecialtyHandler = getProfessionalSchedules(db)

	api.ProfessionalsGetProfessionalbyIDHandler = getProfessional(db)

	api.ProfessionalsGetSpecialtiesByProfessionalHandler = getProfessionalSpecialties(db)

	api.LoginLoginHandler = authenticate(db)

	api.AppointmentsSearchAppointmentHandler = searchAppointments(db)

	api.SpecialtiesSearchSpecialtyHandler = searchSpecialties(db)

	api.ProfessionalsSetProfesionalScheduleBySpecialtyHandler = setProfessionalSchedules(db, cfg.AppointmentDuration)

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
