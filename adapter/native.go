package adapter

import (
	"fmt"

	"github.com/Espigah/log5ws"
)

type Native struct {
}

func (a *Native) Info(l *log5ws.LogInfo) {
	fmt.Printf("ID:%s | Who:%s | What:%s | When:%s | Where:%s | Why:%s\n", l.ID, l.Who, l.What, l.When, l.Where, l.Why)
}
