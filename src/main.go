package main

import (
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"configuration"
	"context"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

func main() {
	if err := configuration.Init(); err != nil {
		logrus.Fatalf("configuration init error: %v", err)
	}

	r := mux.NewRouter()

	// TODO: Middleware setup: metrics and other stuff

	handler := http.Handler(r)

	server := &http.Server{
		Addr:         configuration.Bind(),
		Handler:      makeCORS(handler),
		ReadTimeout:  configuration.ReadTimeout(),
		WriteTimeout: configuration.WriteTimeout(),
		IdleTimeout:  configuration.IdleTimeout(),
	}

	server.SetKeepAlivesEnabled(false)

	printRoutes(r)

	go func() {
		logrus.Fatal(server.ListenAndServe())
	}()

	sigc := make(chan os.Signal)
	signal.Notify(sigc, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGKILL, syscall.SIGSTOP)
	<-sigc
	ctx, cancel := context.WithTimeout(
		context.Background(),
		configuration.GracefulTimeout(),
	)
	defer cancel()
	logrus.Info(server.Shutdown(ctx))
}

func makeCORS(router http.Handler) http.Handler {
	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With"})
	originsOk := handlers.AllowedOrigins(configuration.CORSAllowedOrigins())
	methodsOk := handlers.AllowedMethods([]string{"GET", "OPTIONS", "HEAD"})
	return handlers.CORS(originsOk, headersOk, methodsOk)(router)
}

func printRoutes(router *mux.Router) {
	if configuration.ShowConfig() && configuration.IsDevelopment() {
		dumpRoutes(router)
	}
}

func dumpRoutes(router *mux.Router) {
	err := router.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		routeTemplate, err := route.GetPathTemplate()

		if err != nil {
			return err
		}

		logrus.Info(routeTemplate)
		return nil
	})

	if err != nil {
		logrus.Error("dumpRoutes: ", err)
	}
}
