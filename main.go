package main

import (
	"context"
	"github.com/brunchybit/muscle_ops/server"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	log.Println("booting simulation...")
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	router := chi.NewRouter()

	router.Use(middleware.Heartbeat("/ping"))

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		response := map[string]string{
			"muscle_ops": "v0.1",
		}
		server.EncodeAndWrite(w, r, http.StatusOK, response)
	})

	s := &http.Server{
		Addr: ":8088",
		Handler: router,
		ReadTimeout: 15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout: 15 * time.Second,
	}

	shutDown := func() {
		deadline := time.Now().Add(5 * time.Second)
		ctx, cancel := context.WithDeadline(context.Background(), deadline)
		err := s.Shutdown(ctx)
		if err != nil {
			log.Fatalf("problem shutting down %v", err)
		}
		cancel()
	}

	log.Println("starting server...")
	go func() {
		err := s.ListenAndServe()
		log.Fatalf("signal shutdown: %v", err)
	}()
	<- signals
	shutDown()
}
