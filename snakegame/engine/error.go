package engine

type DieError struct {
	msg string
}

func (e *DieError) Error() string {
	return e.msg
}
