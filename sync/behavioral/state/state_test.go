package state

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAlert_Alert_ShouldAlertThreeTimes(t *testing.T) {
	// Arrange
	sms := &SmsAlert{}
	email := &EmailAlert{}
	call := &CallAlert{}

	alert := &Alert{}

	// Act
	alert.SetState(sms)
	res1 := alert.Alert()

	alert.SetState(email)
	res2 := alert.Alert()

	alert.SetState(call)
	res3 := alert.Alert()

	// Assert
	assert.Equal(t, "SMS!", res1)
	assert.Equal(t, "Email!", res2)
	assert.Equal(t, "Call!", res3)
}
