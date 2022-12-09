package contexts

import (
	"context"
	"time"
	"tokowiwin/config"
)

func BuildContextInit(cfg *config.AppConfig) context.Context {
	ctx, _ := context.WithTimeout(context.Background(), time.Duration(cfg.Context.Timeout)*time.Second)
	return ctx
}
