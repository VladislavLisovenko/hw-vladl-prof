package hw09structvalidator

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
)

var (
	errorFieldType = errors.New("field type is not valid")
)

type ValidationError struct {
	Field string
	Err   error
}

type ValidationErrors []ValidationError

func (v ValidationErrors) Error() string {
	s := make([]string, len(v))
	for i, e := range v {
		s[i] = fmt.Sprintf("field %s: %s", e.Field, e.Err)
	}

	return strings.Join(s, "\n")
}

func Validate(v interface{}) error {
	value := reflect.ValueOf(v)
	if value.Kind() != reflect.Struct {
		return errors.New("value is not a struct")
	}

	validationErrors := make(ValidationErrors, 0)

	t := value.Type()
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		validator := field.Tag.Get("validate")
		if validator == "" {
			continue
		}
		fieldValue := value.Field(i)

		switch fieldValue.Kind() {
		case reflect.Int:
			if err := validateIntField(fieldValue.String(), field.Name, validator); err != nil {
				addValidationError(&validationErrors, field.Name, errorFieldType)
			}
		case reflect.String:
			if err := validateStringField(fieldValue.String(), field.Name, validator); err != nil {
				addValidationError(&validationErrors, field.Name, errorFieldType)
			}
		case reflect.Slice:
			switch fieldValue.Type().Elem().Kind() {
			case reflect.Int:
				arr, ok := fieldValue.Interface().([]int)
				if !ok {
					return fmt.Errorf("%w, field %s", errorFieldType, field.Name)
				}
				err = validateIntSliceField(arr, field.Name, validator)
			case reflect.String:
				arr, ok := fieldValue.Interface().([]string)
				if !ok {
					return fmt.Errorf("%w, field %s", errorFieldType, field.Name)
				}
				err = validateStringSliceField(arr, field.Name, validator)
			default:
				return fmt.Errorf("%w, field %s", errorFieldType, field.Name)
			}
		default:
			return fmt.Errorf("%w, field %s", errorFieldType, field.Name)
		}
	}

	if len(validationErrors) > 0 {
		return validationErrors
	}

	return nil
}

func addValidationError(validationErrors *ValidationErrors, fieldName string, err error) {
	*validationErrors = append(*validationErrors, ValidationError{
		Field: fieldName,
		Err:   err,
	})
}

func validateIntField(v int64, fn string, validator string) error {
	validatorTerms := strings.Split(validator, "|")
	validationErrors := make(ValidationErrors, 0)
	for _, condition := range validatorTerms {
		condKeyWord := strings.Split(condition, ":")[0]
		cond := strings.Split(condition, ":")[1]
		var validateFunc func(string, int64) error
		var validateErr error
		switch condKeyWord {
		case "min":
			validateFunc = validateMin
			validateErr = errMinThreshold
		case "max":
			validateFunc = validateMax
			validateErr = errMaxThreshold
		case "in":
			validateFunc = validateInSetInt
			validateErr = errNotInSet
		default:
			return fmt.Errorf("%w: %s for field %s", errUnknownValidationQuery, validator, fn)
		}
		err := validateFunc(cond, v)
		if err != nil {
			if !errors.Is(err, validateErr) {
				return err
			}
			validationErrors = append(validationErrors, ValidationError{Field: fn, Err: err})
		}
	}
	return validationErrors
}

// func validateInt(f reflect.Type) error {
// 	return nil
// }

// func validateString(f reflect.Type) error {
// 	return nil
// }
