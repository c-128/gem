package gem

var (
	NotFoundErr = NewError(StatusNotFound, "not found")
)

func NewError(status int, err string) Error {
	newErr := &defaultError{
		status: status,
		err:    err,
	}
	return newErr
}

type Error interface {
	Status() int
	Error() string
}

type defaultError struct {
	status int
	err    string
}

func (e *defaultError) Status() int {
	return e.status
}

func (e *defaultError) Error() string {
	return e.err
}
