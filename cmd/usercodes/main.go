package main

import (
	"github.com/dkadio/dysr/internal/routes"

	"log"
	"net/http"
)

func main() {

	log.Println("Starting..")
	f, err := routes.NewCodesRouter()
	if err != nil {
		log.Fatal(err)
	}
	log.Println(http.ListenAndServe(":8000", f))
	log.Println("End..")

}
