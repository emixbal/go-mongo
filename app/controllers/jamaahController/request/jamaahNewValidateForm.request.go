package request

import "github.com/gookit/validate"

type JamaahNewValidateForm struct {
	Name string `json:"name" xml:"name" form:"name" validate:"required|min:2"`
	UUID string `json:"uuid" xml:"uuid" form:"uuid" validate:"required|min:36"`
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
		"UUID": "uuid",
	}
}
