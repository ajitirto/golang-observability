package main

import (
	log "github.com/sirupsen/logrus"
)

func initLogger() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetLevel(log.InfoLevel)
}
