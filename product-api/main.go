package main

import (
	"net/http"
	"log"
	"os"
	"time"
	"context"
	"os/signal"
	"Building-Microservices-With-Go/product-api/handlers"
	gohandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/go-openapi/runtime/middleware"
)

func main() {
	l := log.New(os.Stdout, "product-api", log.LstdFlags)

	ph := handlers.NewProducts(l)

	sm := mux.NewRouter()

	getRouter := sm.Methods("GET").Subrouter()
	getRouter.HandleFunc("/products", ph.ListAll)
	getRouter.HandleFunc("/products/{id:[0-9]+}", ph.ListSingle)

	putRouter := sm.Methods("PUT").Subrouter()
	putRouter.HandleFunc("/products", ph.UpdateProducts)
	putRouter.Use(ph.MiddlewareProductValidation)

	postRouter := sm.Methods("POST").Subrouter()
	postRouter.HandleFunc("/products", ph.AddProduct)
	postRouter.Use(ph.MiddlewareProductValidation)

	deleteRouter := sm.Methods("DELETE").Subrouter()
	deleteRouter.HandleFunc("/products/{id:[0-9]+}", ph.DeleteProduct)

	opts := middleware.RedocOpts{SpecURL: "/swagger.yaml"}
	sh := middleware.Redoc(opts, nil)

	getRouter.Handle("/docs", sh)
	getRouter.Handle("/swagger.yaml", http.FileServer(http.Dir("./")))

	// CORS
	ch := gohandlers.CORS(gohandlers.AllowedOrigins([]string{"http://localhost:3000"}))

	s := &http.Server{
		Addr: ":9090",
		Handler: ch(sm),
		IdleTimeout: 120*time.Second,
		ReadTimeout: 1*time.Second,
		WriteTimeout: 1*time.Second,
	}

	go func() {
		err := s.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}()

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <- sigChan
	l.Println("Recieved terminate, graceful houtdown", sig) // ctrl + c で停止した時に表示される
	

	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(tc)
}