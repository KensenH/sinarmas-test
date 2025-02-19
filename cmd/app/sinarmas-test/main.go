package main

import (
	"context"
	"sinarmas-test/internal/boot"

	log "github.com/sirupsen/logrus"
)

func main() {
	ctx := context.Background()

	log.SetFormatter(&log.TextFormatter{
		DisableColors: true,
		FullTimestamp: true,
	})

	err := boot.Start(ctx)
	if err != nil {
		log.Fatalln(err)
	}
}
