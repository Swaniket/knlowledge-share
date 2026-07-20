package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type KSapplication struct {
	Config AppConfig
	Logger *zap.SugaredLogger
}

func (app *KSapplication) mountRoutes() *gin.Engine {
	r := gin.New()

	// r.Use(gin.RequestID())
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	v1Routes := r.Group("/v1")
	{
		v1Routes.GET("/health")
	}

	return r
}

func (app *KSapplication) run(ginHandler *gin.Engine) error {
	srv := &http.Server{
		Addr:         app.Config.Addr,
		Handler:      ginHandler,
		WriteTimeout: time.Second * 30, // Timeout for writing the response - 30 seconds
		ReadTimeout:  time.Second * 10, // Timeout for reading the request - 10 seconds
		IdleTimeout:  time.Minute,      // Timeout for idle connections - 1 minute
	}

	app.Logger.Infow("Server has started at", "addr", app.Config.Addr)

	return srv.ListenAndServe()
}
