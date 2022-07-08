package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"test_task/storage/api"
	"test_task/storage/internal"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func main() {
	psqlUrl := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		"localhost",
		5432,
		"postgres",
		"postgres",
		"rabbitmq_task",
	)

	psqlConn, err := sqlx.Connect("postgres", psqlUrl)
	if err != nil {
		log.Println("Could not connect to psql database", err)
	}
	defer psqlConn.Close()
	strg := internal.NewStoragePg(psqlConn)

	router := api.New(&api.RouterOptions{
		Storage: strg,
	})

	apiServer := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}
	go func() {
		if err := apiServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Println("could not start api server", err)
		}
	}()
	shutdownChan := make(chan os.Signal, 1)
	defer close(shutdownChan)
	signal.Notify(shutdownChan, syscall.SIGTERM, syscall.SIGINT)
	sig := <-shutdownChan

	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), time.Second*10)
	defer shutdownCancel()

	log.Println("received os signal", sig)
	if err := apiServer.Shutdown(shutdownCtx); err != nil {
		log.Println("could not shutdown http server", err)
		return
	}

	log.Println("server shutdown successfully")
}
