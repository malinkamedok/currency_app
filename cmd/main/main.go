package main

import (
	"devops_course_app/internal/app"
	"devops_course_app/internal/config"
	"log"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Error in parsing config: %s\n", err)
		return
	}

	app.Run(cfg)
}
