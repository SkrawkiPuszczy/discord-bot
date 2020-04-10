package http

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/skrawkipuszczy/discord-bot/pkg/config"
)

type svr struct {
	s      *http.Server
	router *gin.Engine
}

func New(cfg *config.EnvConfig) *svr {
	router := gin.Default()
	authorized := router.Group("/", gin.BasicAuth(gin.Accounts{
		cfg.HTTPAdmin: cfg.HTTPPassword,
	}))

	router.Use(static.Serve("/", static.LocalFile(cfg.HTMLStaticDir, false)))
	authorized.GET("/dd", homePageHandler)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: router,
	}
	return &svr{s: srv, router: router}
}
func (svr *svr) Run() error {
	if err := svr.s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Println(err)
		return err
	}
	return nil
}
func (svr *svr) Close() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := svr.s.Shutdown(ctx); err != nil {
		return err
	}
	return nil
}
