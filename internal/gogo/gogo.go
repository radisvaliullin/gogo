package gogo

import (
	"log/slog"
	"time"
)

type GoGo struct {
	logger *slog.Logger
}

func New(logger *slog.Logger) *GoGo {
	return &GoGo{logger: logger}
}

func (g *GoGo) Start() error {
	g.logger.Info("gogo service start")
	for {
		time.Sleep(time.Second * 1)
	}
}
