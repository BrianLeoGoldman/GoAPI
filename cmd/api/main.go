package main

import (
	"fmt"
	"net/http"

	"github.com/brianleogoldman/goapi/internal/handlers"
	"github.com/go-chi/chi"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetReportCaller(true)        // Set the logger on to print file name and line number when logging
	var r *chi.Mux = chi.NewRouter() // Returns a pointer to a Mux type (struct used to set up the API)
	handlers.Handler(r)

	fmt.Println("Starting GO API service...")
	fmt.Println(`
	 ____    _____       ______  ____    ______
	/\  _ \ /\  __ \    /\  _  \/\  _ \ /\__  _\
	\ \ \L\_\ \ \/\ \   \ \ \L\ \ \ \L\ \/_/\ \/
	 \ \ \L_L\ \ \ \ \   \ \  __ \ \ ,__/  \ \ \
	  \ \ \/, \ \ \_\ \   \ \ \/\ \ \ \/    \_\ \__
	   \ \____/\ \_____\   \ \_\ \_\ \_\    /\_____\
	    \/___/  \/_____/    \/_/\/_/\/_/    \/_____/`)

	err := http.ListenAndServe("localhost:8080", r)
	if err != nil {
		log.Error(err)
	}
}
