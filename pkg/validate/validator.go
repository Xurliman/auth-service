package validate

import (
	"errors"
	"fmt"
	"github.com/Xurliman/auth-service/internal/database"
	"github.com/go-playground/validator"
	"reflect"
	"strings"
)

type ErrorResponse struct {
	Error       bool
	FailedField string
	Tag         string
	Value       interface{}
}

func ExtractValidationErrors(req interface{}) (validationErrors []ErrorResponse) {
	validate := validator.New()
	errs := validate.Struct(req)
	if errs != nil {
		for _, err := range errs.(validator.ValidationErrors) {
			var elem ErrorResponse

			elem.FailedField = err.Field()
			elem.Tag = err.Tag()
			elem.Value = err.Value()
			elem.Error = true

			validationErrors = append(validationErrors, elem)
		}
	}

	return validationErrors
}

func ExtractValidationError(req interface{}) error {
	var message string
	var v = validator.New()

	err := v.RegisterValidation("unique", validateColumnUnique)
	if err != nil {
		return fmt.Errorf("failed to register 'unique' validation %s", err)
	}

	v.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})

	err = v.Struct(req)
	if err != nil {
		for i, err := range err.(validator.ValidationErrors) {
			if i > 0 {
				message += " | "
			}

			if err.Tag() == "unique" {
				message += err.Field() + ": not unique"
			} else {
				message += err.Field() + ": " + err.Tag()
			}
		}

		return errors.New(message)
	}

	return nil
}

func validateColumnUnique(fl validator.FieldLevel) bool {
	param := fl.Param()
	parts := strings.Split(param, ":")
	if len(parts) != 2 {
		return false
	}
	columnName := parts[0]
	tableName := parts[1]
	fieldValue := fl.Field().Interface()

	// Prepare the query
	var count int64
	db := database.GetDB()
	var query = fmt.Sprintf(`SELECT COUNT(*) FROM %s WHERE %s = ?`, tableName, columnName)
	if err := db.Raw(query, fieldValue).Count(&count).Error; err != nil {
		return false
	}

	return count == 0
}
