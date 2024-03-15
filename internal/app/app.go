package app

import (
	"context"
	"currency-operations/config"
	handler "currency-operations/internal/handler/http"
	"currency-operations/internal/repo"
	"currency-operations/internal/service"
	"currency-operations/pkg/httpserver"
	"currency-operations/pkg/postgres"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	_ "github.com/lib/pq"

	"github.com/sirupsen/logrus"
)

func Run(configPath string) {
	logrus.SetFormatter(&logrus.JSONFormatter{})

	cnf, err := config.New(configPath)
	if err != nil {
		logrus.Fatalf("error initializing configs: %s", err.Error())
	}

	db, err := postgres.NewPostgresDB(postgres.ConfigPG{
		Host:         cnf.PG.Host,
		Port:         cnf.PG.Port,
		Username:     cnf.PG.Username,
		Password:     cnf.PG.Password,
		DBName:       cnf.PG.DBname,
		SSLMode:      cnf.PG.SslMode,
		ConnAttempts: cnf.PG.ConnAttempts,
	})

	if err != nil {
		logrus.Fatalf("failed to initialize db: %s", err.Error())
	}

	repos := repo.NewRepository(db)
	service := service.NewService(repos)
	handlers := handler.NewHandler(service, cnf.API)

	srv := new(httpserver.Server)

	go func() {
		if err := srv.Run(cnf.HTTP.Port, handlers.InitRoutes()); err != nil && err != http.ErrServerClosed {
			logrus.Fatalf("error occured while running http server: %s", err.Error())
		}
	}()

	logrus.Printf("Currency operations app is running on port: %s", cnf.HTTP.Port)

	quit := make(chan os.Signal, 1)

	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Println("Currency operations app shutting down")

	if err := srv.ShutDown(context.Background()); err != nil {
		logrus.Errorf("error while shuttung down Currency operations app: %s", err.Error())
	}

	if err := db.Close(); err != nil {
		logrus.Errorf("error while closing database connection: %s", err.Error())
	}

}
