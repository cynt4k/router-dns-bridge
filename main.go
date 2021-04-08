package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/cynt4k/router-dns-bridge/cmd"
)

var (
	Version = "Unknown"
	Build   = "Unknown"
)

func main() {
	log.Printf("starting router-dns-bridge version: %s - build: %s", Version, Build)

	err := cmd.Execute()
	if err != nil {
		log.Fatal(err)
	}

	shutdownCtx, shutdown := context.WithCancel(context.Background())
	defer shutdown()

	err = waitForInterrupt(shutdownCtx)
	log.Printf("shuting down: %s", err)
}

func waitForInterrupt(ctx context.Context) error {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)

	select {
	case sig := <-c:
		return fmt.Errorf("received signal %s", sig)
	case <-ctx.Done():
		return errors.New("canceled")
	}
}
