package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"tokowiwin/config"
	"tokowiwin/servers/http"
	"tokowiwin/utils/contexts"
)

func init() {
	config.InitAppConfig()
}

func main() {
	log.Printf("Starting tokowiwin-http service...")

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs)
	go func() {
		for sig := range sigs {
			if sig != syscall.SIGURG {
				log.Printf("Stopping tokowiwin-http service...")
				os.Exit(1)
			}
		}
	}()

	cfg := config.GetConfig()
	ctx := contexts.BuildContextInit()
	http.InitFactoryHTTP(ctx, cfg)
}
