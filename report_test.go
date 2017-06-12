package report

import "testing"

func TestValidateStruct(t *testing.T) {
	a := struct {
		FirstName string `valid:"required~Require Firstname"`
		LastName  string
	}{
		FirstName: "",
		LastName:  "makimoto",
	}

	r := NewReport()
	if err := ValidateStruct(a, r); err != nil {
		t.Error(err)
	}

	if !r.HasErrors() {
		t.Error("should get a validation error")
	}

	if r.Errors[0].Err.Error() != "Require Firstname" {
		t.Errorf("invalid error message: %s", r.Errors[0].Err.Error())
	}

	if r.Errors[0].Name != "FirstName" {
		t.Errorf("invalid error name: %s", r.Errors[0].Name)
	}
}
