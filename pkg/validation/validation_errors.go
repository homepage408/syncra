package validation

import (
	"errors"
	"fmt"

	"github.com/go-playground/validator/v10"
)

func FormatValidationError(err error) map[string]string {
	errs := make(map[string]string)

	var ve validator.ValidationErrors
	if errors.As(err, &ve) {
		for _, f := range ve {
			// f.Field() sekarang otomatis jadi "email", "password", "phone_number", dll.
			fieldName := f.Field()

			switch f.Tag() {
			case "required":
				errs[fieldName] = fmt.Sprintf("Field %s wajib diisi", fieldName)
			case "email":
				errs[fieldName] = fmt.Sprintf("Format %s tidak valid", fieldName)
			case "min":
				errs[fieldName] = fmt.Sprintf("Field %s minimal %s karakter", fieldName, f.Param())
			case "max":
				errs[fieldName] = fmt.Sprintf("Field %s maksimal %s karakter", fieldName, f.Param())
			case "numeric":
				errs[fieldName] = fmt.Sprintf("Field %s harus berupa angka", fieldName)
			default:
				errs[fieldName] = fmt.Sprintf("Field %s tidak valid (%s)", fieldName, f.Tag())
			}
		}
	}

	return errs
}
