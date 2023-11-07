package command

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInvoker_Execute(t *testing.T) {
	// Arrange
	receiver := &Receiver{}

	command1 := &ToggleOnCommand{receiver: receiver}
	command2 := &ToggleOffCommand{receiver: receiver}

	invoker := &Invoker{}
	invoker.AddCommand(command1)
	invoker.AddCommand(command2)

	// Act
	res := invoker.Execute()

	// Assert
	assert.Equal(t, "Toggle on!\nToggle off!\n", res)
}
