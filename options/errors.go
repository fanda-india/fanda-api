package options

// NewNotFoundError method
func NewNotFoundError(name string) error {
	return &NotFoundError{Name: name}
}

// NotFoundError type
type NotFoundError struct {
	Name string
}

func (e *NotFoundError) Error() string { return e.Name + " not found" }

// ErrBadRequest type
// var ErrBadRequest = errors.New("Bad request")
