package log5ws_test

import (
	"testing"

	"github.com/Espigah/log5ws"
	"github.com/Espigah/log5ws/adapter"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"go.uber.org/zap"
)

func TestCtxDisabled(t *testing.T) {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	log5ws.New(&adapter.Native{}).ID("Native").Who("me").What("something").When("today").Where("here").Why("because").Info()

	log5ws.New(&adapter.Zerolog{Logger: log.Logger}).ID("Zerolog").Who("me").What("something").When("today").Where("here").Why("because").Info()

	log5ws.New(&adapter.Zap{Logger: *logger}).ID("Zap").Who("me").What("something").When("today").Where("here").Why("because").Info()

}
