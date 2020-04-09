package http

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

type svr struct {
	s      *http.Server
	router *gin.Engine
}

func New() *svr {
	port := os.Getenv("PORT")

	if port == "" {
		port = "3000"
	}
	router := gin.Default()
	router.GET("/", homePageHandlder)
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: router,
	}
	return &svr{s: srv, router: router}
}
func (svr *svr) Run() error {
	if err := svr.s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		return err
	}
	return nil
}
func (svr *svr) Close() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := svr.s.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}
}
