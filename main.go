package main

import (
	"awesomeProject1/ingest"

	_ "github.com/lib/pq"
)

const (
	serverAddress = "0.0.0.0:8080"
)

func main() {

	ingest.SetupService()
}
