package command

type Command interface {
	Execute() string
}

type ToggleOnCommand struct {
	receiver *Receiver
}

func (t *ToggleOnCommand) Execute() string {
	return t.receiver.ToggleOn()
}

type ToggleOffCommand struct {
	receiver *Receiver
}

func (t *ToggleOffCommand) Execute() string {
	return t.receiver.ToggleOff()
}

type Receiver struct{}

func (r *Receiver) ToggleOn() string {
	return "Toggle on!"
}

func (r *Receiver) ToggleOff() string {
	return "Toggle off!"
}

type Invoker struct {
	commands []Command
}

func (i *Invoker) AddCommand(command Command) {
	i.commands = append(i.commands, command)
}

func (i *Invoker) RemoveCommand() {
	if len(i.commands) != 0 {
		i.commands = i.commands[:len(i.commands)-1]
	}
}

func (i *Invoker) Execute() string {
	var result string
	for _, command := range i.commands {
		result += command.Execute() + "\n"
	}

	return result
}
