package runner

import (
	"fmt"
	"github.com/theghostmac/web3trakka/internal/housekeeper"
	"net/http"
	"os"
	"time"
)

type StartRunner struct {
	ListenAddr             string
	PostgresSQLDatabaseURL string
	WriteTimeout           time.Duration
	ReadTimeout            time.Duration
	IdleTimeout            time.Duration
	HandlerTimeout         time.Duration
}

func NewStartRunner() *StartRunner {
	// TODO: create a Docker Postgres instance and put the following:
	dbConnString := os.Getenv("DB_CONN_URL")

	return &StartRunner{
		ListenAddr:             ":7080",
		PostgresSQLDatabaseURL: dbConnString,
		WriteTimeout:           time.Second * 30,
		ReadTimeout:            time.Second * 30,
		IdleTimeout:            time.Second * 30,
		HandlerTimeout:         time.Second * 30,
	}
}

func (sr *StartRunner) StartServer() {
	// New logger for housekeeping.
	logger := housekeeper.NewCustomLogger()

	// Set up the HTTP server.
	server := &http.Server{
		Addr:         sr.ListenAddr,
		WriteTimeout: sr.WriteTimeout,
		ReadTimeout:  sr.ReadTimeout,
		IdleTimeout:  sr.IdleTimeout,
		Handler:      nil,
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to the web3trakka server!")
	})

	// Logger for starting the server.
	infoMsg := fmt.Sprintf("Starting the web3trakka server on %s", sr.ListenAddr)
	logger.Info(infoMsg)

	// Start the server.
	if err := server.ListenAndServe(); err != nil {
		errMsg := fmt.Sprintf("Error starting the web3trakka server: %s", err.Error())
		logger.Fatal(errMsg)
	}
}
