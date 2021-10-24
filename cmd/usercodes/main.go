package main

import (
	"github.com/dkadio/dysr/internal/routes"
	"github.com/dkadio/dysr/util"

	"log"
	"net/http"
)

func main() {

	config := util.LoadConfig()
	log.Println("Config Parsed:", config)

	log.Println("Starting..")
	f, err := routes.NewCodesRouter()
	if err != nil {
		log.Fatal(err)
	}
	log.Println(http.ListenAndServe(":8000", f))
	log.Println("End..")

}
