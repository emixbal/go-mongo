package request

import "github.com/gookit/validate"

type JamaahNewValidateForm struct {
	Name string `json:"name" xml:"name" form:"name" validate:"required|min:2"`
}

// Messages you can custom validator error messages.
func (f JamaahNewValidateForm) Messages() map[string]string {
	return validate.MS{
		"required": "{field} is required.",
		"int":      "{field} must int",
	}
}

// Translates you can custom field translates.
func (f JamaahNewValidateForm) Translates() map[string]string {
	return validate.MS{
		"Name": "name",
	}
}
