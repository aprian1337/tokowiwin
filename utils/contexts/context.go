package contexts

import (
	"context"
	"time"
	"tokowiwin/config"
)

func BuildContextInit() context.Context {
	cfg := config.GetConfig()
	ctx, _ := context.WithTimeout(context.Background(), time.Duration(cfg.Context.TimeoutInit)*time.Second)
	return ctx
}

func BuildContextApp() context.Context {
	cfg := config.GetConfig()
	ctx, _ := context.WithTimeout(context.Background(), time.Duration(cfg.Context.TimeoutApp)*time.Second)
	return ctx
}
