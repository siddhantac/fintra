package main

import (
	// "log"
	"context"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/siddhantac/fintra/http/rest"
	"github.com/siddhantac/fintra/infra/store"
	"github.com/siddhantac/fintra/repository"
	"github.com/siddhantac/fintra/service"
)

func main() {
	if err := run(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}

func run() error {
	var (
		port string
	)

	flag.StringVar(&port, "port", "8090", "port on which server will run")
	flag.Parse()

	txnRepo := repository.NewTransactionRepository(store.NewMemStore())
	accRepo := repository.NewAccountRepository(store.NewMemStore())
	txnSvc := service.NewTransactionService(txnRepo, accRepo)
	accSvc := service.NewAccountService(accRepo)
	txnHandler := rest.NewTransactionHandler(txnSvc)
	accHandler := rest.NewAccountHandler(accSvc)

	log.Println("starting server...")

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.SetHeader("content-type", "application/json"))
	r.Get("/healthcheck", txnHandler.HealthCheck)
	r.Route("/transactions", func(r chi.Router) {
		r.Post("/", txnHandler.CreateTransaction)
		r.Get("/{txnID}", txnHandler.GetTransactionByID)
		r.Get("/", txnHandler.GetAllTransactions)
	})
	r.Route("/accounts", func(r chi.Router) {
		r.Get("/", accHandler.GetAllAccounts)
		r.Post("/", accHandler.CreateAccount)
	})

	var wg sync.WaitGroup
	wg.Add(1)
	srv := startServer(&wg, r, port)

	// Setting up signal capturing
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, os.Kill)
	<-stop // waiting for SIGINT (kill -2)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Printf("server shutdown: %v\n", err)
	}

	log.Println("stopped")
	return nil
}

func startServer(wg *sync.WaitGroup, r http.Handler, port string) *http.Server {
	srv := &http.Server{Addr: ":" + port, Handler: r}

	go func() {
		defer wg.Done()

		if err := srv.ListenAndServe(); err != http.ErrServerClosed {
			log.Printf("ListenAndServe(): %v\n", err)
		}
	}()

	return srv
}
