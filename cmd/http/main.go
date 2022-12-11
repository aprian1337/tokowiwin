package main

import (
	"tokowiwin/config"
	"tokowiwin/servers/http"
	"tokowiwin/utils/contexts"
)

func init() {
	config.InitAppConfig()
}

func main() {
	cfg := config.GetConfig()
	ctx := contexts.BuildContextInit()
	http.InitFactoryHTTP(ctx, cfg)
}
