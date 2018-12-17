package deck

import "github.com/pkg/errors"

type ErrorMessage string

const (
	ErrInvalidDeckString  ErrorMessage = "Invalid Deck String"
	ErrUnsupportedVersion ErrorMessage = "Unsupported Deck String Version"
)

func (e ErrorMessage) String() string {
	return string(e)
}

func (e ErrorMessage) Error() string {
	return e.String()
}

func (e ErrorMessage) Wrap(err error) error {
	return errors.Wrap(err, e.String())
}
