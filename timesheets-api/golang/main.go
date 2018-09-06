package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, // "http://localhost:3000", "http://localhost:4200"
		AllowCredentials: true,
		AllowedHeaders:   []string{"*"}, //Authorization
	})

	r := mux.NewRouter()

	r.Handle("/timesheets", SecureWithScope("read:timesheets", TimesheetHandler))
	r.Handle("/open/timesheets", http.HandlerFunc(TimesheetHandler))

	handler := c.Handler(r)
	http.Handle("/", r)

	fmt.Println("Listening on http://localhost:8080")
	http.ListenAndServe("0.0.0.0:8080", handler)
}
