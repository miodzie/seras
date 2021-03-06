package seras

var token string = "!"

func Token() string {
	return token
}

type Messenger interface {
	Send(Message) error
	Reply(Message, string) error
}

type Message struct {
	Content       string
	Arguments     []string
	Channel       string
	Author        Author
}

type Author struct {
	Id      string // Host in IRC, User ID in Discord.
	Nick    string
	Mention string // TODO: Refactor?, this is quick fix to get mentions working in Discord.
}

type MessageFormatter interface {
	Bold(string) string
	Italicize(string) string
}

func (msg *Message) Command(command string, call func(Message)) {
	if msg.IsCommand(command) {
		call(*msg)
	}
}

func (msg *Message) IsCommand(command string) bool {
	return token+command == msg.Arguments[0]
}

type NullMessenger struct{}

func (messenger *NullMessenger) Send(msg Message) error {
	return nil
}
