package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"os"
	"os/exec"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
	ownhttp "github.com/zainul/txn/internal/delivery/http"
	"github.com/zainul/txn/internal/pkg/initial"
	"github.com/zainul/txn/internal/repository/store"
	"github.com/zainul/txn/internal/usecase"
)

func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
		return
	})
}

func main() {
	var wait time.Duration
	flag.DurationVar(&wait,
		"graceful-timeout", time.Second*15,
		"the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m")
	flag.Parse()

	db := initial.GetDB()
	defer db.Close()

	// update schema
	gopath := os.Getenv("GOPATH")
	cmd := exec.Command("gan", "-c", gopath+"/src/github.com/zainul/txn/migrations/config.json", "migrate")

	if out, err := cmd.CombinedOutput(); err != nil {
		log.Println(string(out), err)
	} else {
		log.Println(string(out))
	}

	r := mux.NewRouter()

	userStore := store.NewUserAccountStore(db)
	txLogStore := store.NewTransactionLogStore(db)
	txHistoryStore := store.NewTransactionHistoryStore(db)

	usecaseUser := usecase.NewUser(userStore)
	usecaeTransaction := usecase.NewTransaction(txHistoryStore, txLogStore, userStore)

	ownhttp.NewUserHanlder(r, usecaseUser)
	ownhttp.NewTxHanlder(r, usecaeTransaction)

	r.Use(commonMiddleware)

	http.Handle("/", r)
	// config the http server
	srv := &http.Server{
		Addr: "0.0.0.0:7000",
		// Good practice to set timeouts to avoid Slowloris attacks.
		// WriteTimeout: time.Second * 1,
		ReadTimeout: time.Second * 5, // -> it will be in config file
		// IdleTimeout: time.Second * 1,
		Handler: r, // Pass our instance of gorilla/mux in.
	}

	// Run our server in a goroutine so that it doesn't block.
	go func() {
		log.Println("Server up :7000")
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	c := make(chan os.Signal, 1)
	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(c, os.Interrupt)

	// Block until we receive our signal.
	<-c

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()
	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	srv.Shutdown(ctx)
	// Optionally, you could run srv.Shutdown in a goroutine and block on
	// <-ctx.Done() if your application should wait for other services
	// to finalize based on context cancellation.
	log.Println("shutting down")
	os.Exit(0)
}
