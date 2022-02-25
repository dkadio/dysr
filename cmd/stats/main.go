package main

import (
	"github.com/dkadio/dysr/internal/controllers"
	//	"github.com/ip2location/ip2location-go"
	//"log"
	"net/http"
)

func main() {

	sc := controllers.NewStatsController()
	sc.RegisterForEvents()

	//db, err := ip2location.OpenDB("./IP2LOCATION-LITE-DB11.BIN")
	//
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//
	//	ip := "8.8.8.8"
	//	r, err := db.Get_all(ip)
	//
	//	log.Printf("%+v", r)
	//
	http.ListenAndServe(":8002", nil)
}
