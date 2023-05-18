package messages

type Spinner interface {
	Start() error
	Pause() error
	Unpause() error
	Stop() error
	Message(message string)
	Colors(colors ...string) error
	StopMessage(message string)
	StopColors(colors ...string) error
}

type NullSpinner struct{}

func (n NullSpinner) Start() error {
	return nil
}

func (n NullSpinner) Pause() error {
	return nil
}

func (n NullSpinner) Unpause() error {
	return nil
}

func (n NullSpinner) Stop() error {
	return nil
}

func (n NullSpinner) Message(message string) {
}

func (n NullSpinner) Colors(colors ...string) error {
	return nil
}

func (n NullSpinner) StopMessage(message string) {
}

func (n NullSpinner) StopColors(colors ...string) error {
	return nil
}
