package tracker

type InvalidEventError struct{}

func (e InvalidEventError) Error() string {
	return "invalid event"
}

func NewInvalidEventError() InvalidEventError {
	return InvalidEventError{}
}
