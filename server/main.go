package main

import (
	"log"
	"net/http"

	"gin-eco/server/routes"
)

func main() {
	log.Fatal(http.ListenAndServe(":5000", routes.Engine()))
}
