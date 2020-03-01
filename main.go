package main

import (
	"log"
	"net/http"

	"gin-eco/server/routes"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load("")
	myEnv, err := godotenv.Read()
	if err != nil {
		log.Panicln("err:", err.Error())
	}
	log.Fatal(http.ListenAndServe(":"+myEnv["SERVER_PORT"], routes.Engine()))
}
