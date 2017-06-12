package example

import (
	"fmt"
	"github.com/kohkimakimoto/govalidator-report"
)

type Person struct {
	FirstName string `valid:"required~Require first name"`
	LastName  string
}

func Example() {
	person := &Person{
		FirstName: "",
		LastName:  "bar",
	}

	r := report.NewReport()

	// validate a struct by using tags.
	if err := report.ValidateStruct(person, r); err != nil {
		panic(err)
	}

	// you can write additional validation code.
	if person.LastName == "foo" {
		r.AppendErrorMessage("LastName", "LastName must not be foo")
	}

	// report validation errors.
	if r.HasErrors() {
		for _, err := range r.Errors {
			fmt.Printf("%s: %s\n", err.Name, err.Err.Error())
		}
	}
	// Output: FirstName: Require first name
}
