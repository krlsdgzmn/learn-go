package main

import (
	"flag"
	"fmt"
	"net/http"
	"time"

	"github.com/krlsdgzmn/learn-go/5_projectStructuring/internal/app"
	"github.com/krlsdgzmn/learn-go/5_projectStructuring/internal/routes"
)

func main() {
	var port int
	flag.IntVar(&port, "port", 8080, "Server port.")
	flag.Parse()

	app, err := app.NewApplication()
	if err != nil {
		panic(err)
	}
	defer app.Database.Close()

	r := routes.SetupRoutes(app)
	server := http.Server{
		Addr:         fmt.Sprintf(":%d", port),
		Handler:      r,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}
	app.Logger.Printf("Application is running on port %d.\n", port)

	err = server.ListenAndServe()
	if err != nil {
		app.Logger.Fatal(err)
	}
}
