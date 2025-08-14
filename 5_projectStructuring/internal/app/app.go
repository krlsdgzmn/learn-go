package app

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/krlsdgzmn/learn-go/5_projectStructuring/internal/api"
	"github.com/krlsdgzmn/learn-go/5_projectStructuring/internal/store"
	"github.com/krlsdgzmn/learn-go/5_projectStructuring/migrations"
)

type Application struct {
	Logger         *log.Logger
	WorkoutHandler *api.WorkoutHandler
	Database       *sql.DB
}

func NewApplication() (*Application, error) {
	db, err := store.Open()
	if err != nil {
		return nil, err
	}

	err = store.MigrateFS(db, migrations.FS, ".")
	if err != nil {
		panic(err)
	}
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	workoutStore := store.NewPostgresWorkoutStore(db)
	workoutHandler := api.NewWorkoutHandler(workoutStore)
	app := &Application{
		Logger:         logger,
		WorkoutHandler: workoutHandler,
		Database:       db,
	}
	return app, nil
}

func (a *Application) HealthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Status is available.\n")
}
