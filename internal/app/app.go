package app

import (
	"devops_course_app/internal/config"
	v1 "devops_course_app/internal/controller/http/v1"
	"devops_course_app/internal/usecase"
	"devops_course_app/internal/usecase/cbrf"
	"devops_course_app/pkg/httpserver"
	"github.com/go-chi/chi/v5"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func Run(cfg *config.Config) {

	i := usecase.NewInfoUseCase(cbrf.NewInfoReq())

	handler := chi.NewRouter()

	v1.NewRouter(handler, i)

	server := httpserver.New(handler, httpserver.Port(cfg.AppPort))
	interruption := make(chan os.Signal, 1)
	signal.Notify(interruption, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interruption:
		log.Printf("signal: " + s.String())
	case err := <-server.Notify():
		log.Printf("Notify from http server: %s\n", err)
	}

	err := server.Shutdown()
	if err != nil {
		log.Printf("Http server shutdown")
	}
}
