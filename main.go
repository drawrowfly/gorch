package main

import (
	"context"
	"log"
	"math/rand"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
)

const (
	serverPort = "8080"
)

var router = mux.NewRouter()

func main() {

	rand.Seed(time.Now().UnixNano())

	router.HandleFunc("/wallet/create/{home}", func(w http.ResponseWriter, r *http.Request) {
		createWalletHandler(w, r)
	}).Methods("GET")

	router.HandleFunc("/wallet/list/{home}", func(w http.ResponseWriter, r *http.Request) {
		getWalletListHandler(w, r)
	}).Methods("GET")

	srv := &http.Server{
		Handler:      router,
		Addr:         ":" + serverPort,
		WriteTimeout: 45 * time.Second,
		ReadTimeout:  45 * time.Second,
	}

	// Start Server
	go func() {
		log.Println("Starting Server: ", serverPort)
		if err := srv.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()

	// Graceful Shutdown
	waitForShutdown(srv)
}

func waitForShutdown(srv *http.Server) {
	interruptChan := make(chan os.Signal, 1)
	signal.Notify(interruptChan, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	<-interruptChan

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	srv.Shutdown(ctx)

	log.Println("Shutting down")
	os.Exit(0)
}
