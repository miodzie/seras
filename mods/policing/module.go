package policing

import (
	"fmt"
	"time"

	"github.com/miodzie/seras"
)

type PolicingMod struct {
	seras.BaseModule
}

func NewPolicingMod() *PolicingMod {
	mod := &PolicingMod{}
	mod.LoopCheck = func() {
		for mod.Running {
			msg := <-mod.Stream
			if IsSpam(msg) {
				mod.Actions.Send(seras.Message{Content: "bruh, shut up", Channel: msg.Channel})
        err := mod.Actions.TimeoutUser(msg.Channel, msg.AuthorId, time.Now().Add(time.Minute * 1))
        if err != nil {
          fmt.Printf("Failed to TimeoutUser: \"%s\"\n", err)
        }
			}
		}
	}

	return mod
}
