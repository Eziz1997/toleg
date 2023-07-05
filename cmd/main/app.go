package main

import (
	repeatable "toleg/pkg/utils"

	"context"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/rs/cors"
	"net"
	"net/http"
	"time"
	"toleg/internal/config"
	handlermanager "toleg/internal/handlers/manager"
	"toleg/pkg/client/postgresql"
	"toleg/pkg/logging"
)

func main() {
	cfg := config.GetConfig()
	logger := logging.GetLogger()
	postgresSQLClient, err := postgresql.NewClient(context.TODO(), 3, cfg.Storage)

	if err != nil {
		logger.Fatalf("%v", err)
	}

	err = repeatable.CrateDir()

	start(handlermanager.Manager(postgresSQLClient, logger), cfg, postgresSQLClient)

}

func start(router *mux.Router, cfg *config.Config, pGPool *pgxpool.Pool) {
	logger := logging.GetLogger()
	logger.Info("start application")

	//go autorun.AutoListen(pGPool)
	var listener net.Listener
	var listenErr error

	logger.Info("listen tcp")
	listener, listenErr = net.Listen("tcp", fmt.Sprintf("%s:%s", cfg.Listen.BindIP, cfg.Listen.Port))
	logger.Infof("server is listening port %s:%s", cfg.Listen.BindIP, cfg.Listen.Port)

	if listenErr != nil {
		logger.Fatal(listenErr)
	}

	fileServer := http.FileServer(http.Dir("../../uploads/"))
	router.PathPrefix("/api/v1/uploads/").Handler(http.StripPrefix("/api/v1/uploads/", fileServer))

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		AllowedHeaders: []string{
			"*",
		},
	})
	handler := c.Handler(router)

	server := &http.Server{
		Handler:      handler,
		WriteTimeout: 5000 * time.Second,
		ReadTimeout:  5000 * time.Second,
	}
	logger.Fatal(server.Serve(listener))
}
