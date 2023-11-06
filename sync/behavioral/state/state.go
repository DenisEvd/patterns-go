package state

type AlertSystem interface {
	Alert() string
}

type Alert struct {
	state AlertSystem
}

func (a *Alert) SetState(alert AlertSystem) {
	a.state = alert
}

func (a *Alert) Alert() string {
	return a.state.Alert()
}

type SmsAlert struct{}

func (s *SmsAlert) Alert() string {
	return "SMS!"
}

type EmailAlert struct{}

func (e *EmailAlert) Alert() string {
	return "Email!"
}

type CallAlert struct{}

func (c *CallAlert) Alert() string {
	return "Call!"
}
