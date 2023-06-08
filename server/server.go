package server

import (
	"context"
	"fmt"
	"net"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hi-supergirl/go-microservice-template/config"
	"github.com/hi-supergirl/go-microservice-template/handlers"
	"github.com/hi-supergirl/go-microservice-template/handlers/services"
	"github.com/hi-supergirl/go-microservice-template/handlers/services/repositories"
	"github.com/hi-supergirl/go-microservice-template/handlers/services/repositories/database"
	"github.com/hi-supergirl/go-microservice-template/logging"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
)

func Server(lc fx.Lifecycle, logger *zap.Logger) *gin.Engine {
	r := gin.Default()

	srv := &http.Server{Addr: ":8080", Handler: r} // define a web server
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			ln, err := net.Listen("tcp", srv.Addr)
			if err != nil {
				logger.Sugar().Infoln("Failed to start HTTP server at", srv.Addr)
				return err
			}
			go srv.Serve(ln)
			logger.Sugar().Infoln("Succeeded to start HTTP server at", srv.Addr)
			return nil
		},
		OnStop: func(ctx context.Context) error {
			srv.Shutdown(ctx)
			logger.Sugar().Infoln("HTTP server is stopped")
			return nil
		},
	})
	return r
}

func printBanner(logger *zap.Logger) {
	logger.Sugar().Infoln(`
****************************************************************
****************************************************************
***************JanessaTech's micro-service template**************
****************************************************************
****************************************************************`)
}

func StartApplication(configFile string) {
	config, err := config.NewConfig(configFile)
	if err != nil {
		fmt.Println("Failed to read config file. Exit")
	}

	app := fx.New(
		fx.Supply(logging.GetLogger(config), config),
		fx.Invoke(printBanner),
		fx.WithLogger(func(logger *zap.Logger) fxevent.Logger {
			return &fxevent.ZapLogger{Logger: logger.Named("JanessaTech Template")}
		}),
		fx.Provide(
			Server,
			database.NewDataBase,
			repositories.NewAccountDB,
			services.NewAccountService,
			handlers.NewAccountHandler,
			handlers.NewProductController,
		),
		fx.Invoke(
			func(*gin.Engine) {},
			handlers.AccountRoute,
			handlers.ProductRoute,
		),
	)
	app.Run()
}
