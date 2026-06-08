package main

import (
	"log"

	"github.com/ebenamoafo22/prolog/internal/server"
)

func main() {
	srv := server.NewHTTPServer(":8080")
	log.Fatal(srv.ListenAndServe())
}