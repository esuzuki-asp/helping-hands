package cmd

import (
	"fmt"
	"helping-hands/service/item"
	"helping-hands/service/location"
	"helping-hands/service/user"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/sirupsen/logrus"
)

func runServer() {
	http.HandleFunc("/ping", func(res http.ResponseWriter, _ *http.Request) {
		res.Write([]byte(`pong`))
		logrus.Println("pong")
	})

	http.Handle(item.PathPrefix, item.NewHandler())
	http.Handle(user.PathPrefix, user.NewHandler())
	http.Handle(location.PathPrefix, location.NewHandler())

	server := &http.Server{
		Addr: fmt.Sprintf(":%s", os.Getenv("PORT")),
	}

	sig := make(chan os.Signal)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)

	go func() {
		err := server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			logrus.Fatalf("listen: %s\n", err)
		}
	}()

	<-sig
	logrus.Info("Starting shutdown")
	// shutdown procedure
	//db.RemoveTables(db.DB)
	logrus.Info("Shutdown complete")
}
