package seras

import (
	"sync"

	"gorm.io/gorm"
)

type Module interface {
	Name() string
	Loop(Stream, Actions) error
	Stop()
	HasDatabase
}

type Actions interface {
	Messenger
	Admin
}

type BaseModule struct {
	Actions Actions
	Stream  Stream
	Running bool
	// See loopCheckExample()
	LoopCheck func()
	db      *gorm.DB
	sync.Mutex
}

func (mod *BaseModule) Loop(stream Stream, actions Actions) error {
	mod.Lock()
	defer mod.Unlock()

	mod.Actions = actions
	mod.Stream = stream
	mod.Running = true
	go mod.LoopCheck()

	return nil
}

func (mod *BaseModule) Stop() {
	mod.Lock()
	defer mod.Unlock()

	mod.Running = false
}

func (mod *BaseModule) DB() *gorm.DB {
	return mod.db
}

func (mod *BaseModule) setDB(db *gorm.DB) {
	mod.db = db
}

func loopCheckExample() {
	mod := &BaseModule{}
	mod.LoopCheck = func() {
		for mod.Running {
			msg := <-mod.Stream
			if msg.Content == "hello" {
				mod.Actions.Send(Message{Content: "Hi", Channel: msg.Channel})
			}
		}
	}
}
