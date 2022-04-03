package adapter

import (
	"github.com/Espigah/log5ws"
	"go.uber.org/zap"
)

type Zap struct {
	Logger zap.Logger
}

func (a *Zap) Info(l *log5ws.LogInfo) {
	a.Logger.Info("",
		zap.String("id", l.ID),
		zap.String("who", l.Who),
		zap.String("what", l.What),
		zap.String("when", l.When),
		zap.String("where", l.Where),
		zap.String("why", l.Why))
}
