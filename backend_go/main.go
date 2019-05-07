package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/rs/cors"

	. "github.com/erick-adl/globo-bbb-wall/backend_go/aws/sqs"
	p "github.com/erick-adl/globo-bbb-wall/backend_go/router/participant"
	"github.com/gorilla/mux"
)

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	return port
}

func init() {
	fmt.Println("Configuring...")
	_, err := Configure()

	if err != nil {
		fmt.Println(err)
		os.Exit(3)
	}
}

func main() {
	r := mux.NewRouter()
	api := r.PathPrefix("/api/v1").Subrouter()
	p.AddHandler(api)

	var port = ":" + getPort()
	fmt.Println("Server running on port:", port)
	handler := cors.Default().Handler(r)
	log.Fatal(http.ListenAndServe(port, handler))
}
