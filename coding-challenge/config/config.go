package config

import (
	"context"
	"time"

	s "github.com/Route42/ct-lib/service"
)

type config struct {
	CTX		context.Context
	CANCEL	context.CancelFunc
	SERVICE	*s.Service
}

func NewConfig() *config {
	return &config{}
}

func LoadConfig(cnfg *config) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)

	cnfg.CTX = ctx
	cnfg.CANCEL = cancel
	cnfg.SERVICE = s.New("Ivan Marinchev")

	return nil
}