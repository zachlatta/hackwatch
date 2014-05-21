package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/zachlatta/hackwatch/backend/database"
	"github.com/zachlatta/hackwatch/backend/handler"
)

const (
	Environment         = "BACKEND_ENV"
	Production          = "PRODUCTION"
	ProductionDatabase  = "DATABASE_URL"
	DevelopmentDatabase = "postgres://docker:docker@$DB_1_PORT_5432_TCP_ADDR/docker"
	databaseDriver      = "postgres"
)

func httpLog(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s", r.RemoteAddr, r.Method, r.URL)
		handler.ServeHTTP(w, r)
	})
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	production := os.Getenv(Environment) == Production

	if production {
		databaseURL := os.Getenv(ProductionDatabase)

		if databaseURL == "" {
			log.Fatal(ProductionDatabase + " is empty")
		}

		err := database.Init(databaseDriver, databaseURL)
		if err != nil {
			panic(err)
		}
	} else {
		err := database.Init(databaseDriver, os.ExpandEnv(DevelopmentDatabase))
		if err != nil {
			panic(err)
		}
	}
	defer database.Close()

	r := mux.NewRouter()

	r.Handle("/hackathons",
		handler.AppHandler(handler.NewHackathon)).Methods("POST")
	r.Handle("/hackathons",
		handler.AppHandler(handler.GetApprovedHackathons)).Methods("GET")

	http.Handle("/", r)
	http.ListenAndServe(":"+port, httpLog(http.DefaultServeMux))
}
