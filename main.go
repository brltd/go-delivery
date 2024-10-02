package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/brltd/delivery/db"
	_ "github.com/brltd/delivery/docs"
	"github.com/brltd/delivery/handlers"
	"github.com/brltd/delivery/logger"
	"github.com/brltd/delivery/middlewares"
	"github.com/brltd/delivery/repositories"
	"github.com/brltd/delivery/services"
	"github.com/joho/godotenv"
	"github.com/julienschmidt/httprouter"
	"github.com/justinas/alice"
	httpSwagger "github.com/swaggo/http-swagger"
)

func main() {
	loadEnv()

	router := httprouter.New()

	dbClient := db.GetSQLiteClient()

	userRepository := repositories.NewUserRepository(dbClient)

	userHandler := handlers.UserHandler{
		UserService: services.NewUserService(userRepository),
	}

	public := alice.New()

	// router.Handler(http.MethodGet, "/customers", dynamic.ThenFunc(customerHandler.GetAllCustomers))
	// router.Handler(http.MethodGet, "/customers/:customer_id", dynamic.ThenFunc(customerHandler.GetCustomer))
	router.Handler(http.MethodPost, "/api/user/register", public.ThenFunc(userHandler.CreateUser))
	router.Handler(http.MethodPost, "/api/user/login", public.ThenFunc(userHandler.CreateUser))

	router.GET("/swagger/*any", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		httpSwagger.WrapHandler.ServeHTTP(w, r)
	})

	// protected := public.Append(middlewares.Authenticate)

	standard := alice.New(middlewares.RecoverPanic, middlewares.LogRequest)

	address := os.Getenv("SERVER_ADDRESS")
	port := os.Getenv("SERVER_PORT")

	tlsConfig := &tls.Config{
		CurvePreferences: []tls.CurveID{tls.X25519, tls.CurveP256},
	}

	srv := &http.Server{
		ErrorLog:    logger.ErrLog,
		Addr:        fmt.Sprintf("%s:%s", address, port),
		Handler:     standard.Then(router),
		TLSConfig:   tlsConfig,
		IdleTimeout: time.Minute,
		// ReadTimeout:  5 * time.Second,
		// WriteTimeout: 10 * time.Second,
	}

	logger.Info(fmt.Sprintf("Starting server on port %s", port))
	// log.Fatal(srv.ListenAndServeTLS("./certificates/cert.pem", "./certificates/key.pem"))
	log.Fatal(srv.ListenAndServe())
}

func loadEnv() {
	godotenv.Load()

	address := os.Getenv("SERVER_ADDRESS")
	port := os.Getenv("SERVER_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASS")
	dbAddress := os.Getenv("DB_ADDR")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	jwtSecret := os.Getenv("JWT_SECRET")
	expHour := os.Getenv("EXP_HOUR")

	message := "Missing environment variable "

	if address == "" {
		logger.Error(message + "SERVER_ADDRESS")
	}

	if port == "" {
		logger.Error(message + "SERVER_PORT")
	}

	if dbUser == "" {
		logger.Error(message + "DB_USER")
	}

	if dbPassword == "" {
		logger.Error(message + "DB_PASS")
	}

	if dbAddress == "" {
		logger.Error(message + "DB_ADDR")
	}

	if dbPort == "" {
		logger.Error(message + "DB_PORT")
	}

	if dbName == "" {
		logger.Error(message + "DB_NAME")
	}

	if jwtSecret == "" {
		logger.Error(message + "JWT_SECRET")
	}

	if expHour == "" {
		logger.Error(message + "EXP_HOUR")
	}
}
