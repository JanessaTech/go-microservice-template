package server

import (
	"context"
	"fmt"
	"net"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

func addGroup(r *gin.Engine) {
	api := r.Group("/api")
	{
		admin := api.Group("/admin")
		{
			admin.GET("/hello", adminFuc)
		}

	}
}

func adminFuc(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"msg": "hello admin"})

}

func Server(lc fx.Lifecycle) *gin.Engine {
	r := gin.Default()
	addGroup(r)

	srv := &http.Server{Addr: ":8080", Handler: r} // define a web server
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			ln, err := net.Listen("tcp", srv.Addr)
			if err != nil {
				fmt.Println("Failed to start HTTP server at", srv.Addr)
				return err
			}
			go srv.Serve(ln)
			fmt.Println("Succeeded to start HTTP server at", srv.Addr)
			return nil
		},
		OnStop: func(ctx context.Context) error {
			srv.Shutdown(ctx)
			fmt.Println("HTTP server is stopped")
			return nil
		},
	})
	return r
}

func StartApplication(configFile string) {
	app := fx.New(
		fx.Provide(
			Server,
		),
		fx.Invoke(func(*gin.Engine) {}),
	)
	app.Run()
}
