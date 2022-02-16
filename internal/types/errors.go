package types

type CommandError struct {
	Message string
}

func (ce *CommandError) Error() string {
	return ce.Message
}
