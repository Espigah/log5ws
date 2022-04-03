package adapter

import (
	"github.com/espigah/log5ws"
	"github.com/rs/zerolog"
)

type Zerolog struct {
	Logger zerolog.Logger
}

func (a *Zerolog) Info(l *log5ws.LogInfo) {
	a.Logger.Info().
		Str("id", l.ID).
		Str("who", l.Who).
		Str("what", l.What).
		Str("when", l.When).
		Str("where", l.Where).
		Str("why", l.Why).
		Msg("")
}
