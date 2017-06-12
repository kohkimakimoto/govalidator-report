package report

import (
	"errors"

	"github.com/asaskevich/govalidator"
)

type Report struct {
	Errors []govalidator.Error
}

func NewReport() *Report {
	r := &Report{
		Errors: []govalidator.Error{},
	}

	return r
}

func (r *Report) HasErrors() bool {
	return len(r.Errors) > 0
}

func (r *Report) AppendError(err govalidator.Error) *Report {
	r.Errors = append(r.Errors, err)

	return r
}

func (r *Report) AppendErrorMessage(name, message string) *Report {
	err := govalidator.Error{
		Name: name,
		Err:  errors.New(message),
		CustomErrorMessageExists: true,
	}

	return r.AppendError(err)
}

func ValidateStruct(i interface{}, r *Report) error {
	if _, err := govalidator.ValidateStruct(i); err != nil {
		if err := appendErrors(r, err); err != nil {
			return err
		}
	}

	return nil
}

func appendErrors(r *Report, err error) error {
	switch e := err.(type) {
	case govalidator.Errors:
		for _, e2 := range e.Errors() {
			if err := appendErrors(r, e2); err != nil {
				return err
			}
		}
	case govalidator.Error:
		r.AppendError(e)
	default:
		return err
	}

	return nil
}
