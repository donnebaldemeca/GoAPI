package main

import (
	"fmt"
	"net/http"

	"github.com/donnebaldemeca/GoAPI/internal/handlers"
	"github.com/go-chi/chi"
	"github.com/gofiber/fiber/v2/log"
	log "github.com/sirupsen/logrus"
)

func main() {

	log.SetReportCaller(true)        // Enable reporting of the caller function in logs
	var r *chi.Mux = chi.NewRouter() // Create a new router
	handlers.Handler(r)

	fmt.Printlin("Starting GO API")

	err := http.ListenAndServe(":8080", r) // Start the server on port 8080, using the router
	if err != nil {                        // Error logging if server fails to start
		log.Error("Error starting server: ", err)
	}
}
