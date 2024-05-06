package hw09structvalidator

import (
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"unicode/utf8"
)

var (
	errorValueIsNotMatchMinimum = errors.New("value is not match minimum")
	errorValueIsNotMatchMaximum = errors.New("value is not match maximum")
	errorValeuIsNotInSet        = errors.New("value is not in set")
	errorStringLength           = errors.New("invalid string length")
	errorStringIsNotMatchRegexp = errors.New("string is not match regular expression")

	errorValueIsNotStruct      = errors.New("value is not struct")
	errorUnsupportedFieldType  = errors.New("unsupported field type for validation")
	errorUnknownValidationRule = errors.New("unknown validation query")
)

type ValidationError struct {
	Field string
	Err   error
}

func (v ValidationError) Error() string {
	return fmt.Sprintf("%s: %s\n", v.Field, v.Err)
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
		return fmt.Errorf("expected struct, got %T. %w", v, errorValueIsNotStruct)
	}

	t := value.Type()
	validationErrors := make(ValidationErrors, 0)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		validationRules := field.Tag.Get("validate")
		if validationRules == "" {
			continue
		}

		fieldVal := value.Field(i)
		var err error
		switch fieldVal.Kind() {
		case reflect.Int:
			err = validateInt(fieldVal.Int(), field.Name, validationRules)
		case reflect.String:
			err = validateString(fieldVal.String(), field.Name, validationRules)
		case reflect.Slice:
			switch fieldVal.Type().Elem().Kind() {
			case reflect.Int:
				arr, ok := fieldVal.Interface().([]int)
				if !ok {
					return fmt.Errorf("field %s. %w", field.Name, errorUnsupportedFieldType)
				}
				err = validateIntSlice(arr, field.Name, validationRules)
			case reflect.String:
				arr, ok := fieldVal.Interface().([]string)
				if !ok {
					return fmt.Errorf("field %s. %w", field.Name, errorUnsupportedFieldType)
				}
				err = validateStringSlice(arr, field.Name, validationRules)
			default:
				return fmt.Errorf("field %s. %w", field.Name, errorUnsupportedFieldType)
			}
		default:
			return fmt.Errorf("field %s. %w", field.Name, errorUnsupportedFieldType)
		}
		if err != nil {
			var validationErrs ValidationErrors
			if !errors.As(err, &validationErrs) {
				return err
			}
			validationErrors = append(validationErrors, validationErrs...)
		}
	}

	if len(validationErrors) > 0 {
		return validationErrors
	}

	return nil
}

func validateInt(fieldValue int64, fieldName string, validationRules string) error {
	rules := strings.Split(validationRules, "|")
	validationErrors := make(ValidationErrors, 0)
	for _, rule := range rules {
		ruleName := strings.Split(rule, ":")[0]
		ruleValue := strings.Split(rule, ":")[1]

		var validateFunc func(string, int64) error
		var validateErr error

		switch ruleName {
		case "min":
			validateFunc = validateMin
			validateErr = errorValueIsNotMatchMinimum
		case "max":
			validateFunc = validateMax
			validateErr = errorValueIsNotMatchMaximum
		case "in":
			validateFunc = validateInSetInt
			validateErr = errorValeuIsNotInSet
		default:
			return fmt.Errorf("unknown validation rule(s) %s for field %s. %w", validationRules, fieldName, errorUnknownValidationRule)
		}

		err := validateFunc(ruleValue, fieldValue)
		if err != nil {
			if !errors.Is(err, validateErr) {
				return err
			}
			validationErrors = append(validationErrors, ValidationError{Field: fieldName, Err: err})
		}
	}
	return validationErrors
}

func validateString(fieldValue string, fieldName string, validationRules string) error {
	rules := strings.Split(validationRules, "|")
	validationErrors := make(ValidationErrors, 0)
	for _, rule := range rules {
		ruleName := strings.Split(rule, ":")[0]
		ruleValue := strings.Split(rule, ":")[1]

		var validateFunc func(string, string) error
		var validateErr error

		switch ruleName {
		case "len":
			validateFunc = validateLength
			validateErr = errorStringLength
		case "in":
			validateFunc = validateInSetString
			validateErr = errorValeuIsNotInSet
		case "regexp":
			validateFunc = validateRegexp
			validateErr = errorStringIsNotMatchRegexp
		default:
			return fmt.Errorf("unknown validation rule(s) %s for field %s. %w", validationRules, fieldName, errorUnknownValidationRule)
		}

		err := validateFunc(ruleValue, fieldValue)
		if err != nil {
			if !errors.Is(err, validateErr) {
				return err
			}
			validationErrors = append(validationErrors, ValidationError{Field: fieldName, Err: err})
		}
	}
	return validationErrors
}

func validateLength(ruleValue string, fieldValue string) error {
	validLength, err := strconv.Atoi(ruleValue)
	if err != nil {
		return err
	}
	if utf8.RuneCountInString(fieldValue) != validLength {
		return fmt.Errorf("string length expected: %d, got %d. %w", validLength, len(fieldValue), errorStringLength)
	}
	return nil
}

func validateInSetInt(ruleValue string, value int64) error {
	validSet := strings.Split(ruleValue, ",")
	valueInSet := false
	for _, s := range validSet {
		if strconv.Itoa(int(value)) == s {
			valueInSet = true
			break
		}
	}
	if !valueInSet {
		return fmt.Errorf("%d must belong to set: %v. %w", value, validSet, errorValeuIsNotInSet)
	}

	return nil
}

func validateInSetString(ruleValue string, value string) error {
	validSet := strings.Split(ruleValue, ",")
	valueInSet := false
	for _, s := range validSet {
		if value == s {
			valueInSet = true
			break
		}
	}
	if !valueInSet {
		return fmt.Errorf("%s must belong to set: %v. %w", value, validSet, errorValeuIsNotInSet)
	}

	return nil
}

func validateRegexp(ruleValue string, value string) error {
	re, err := regexp.Compile(ruleValue)
	if err != nil {
		return err
	}
	if !re.MatchString(value) {
		return fmt.Errorf("%s is not match regex %s. %w", value, ruleValue, errorStringIsNotMatchRegexp)
	}
	return nil
}

func validateMin(ruleValue string, value int64) error {
	goalValue, err := strconv.Atoi(ruleValue)
	if err != nil {
		return err
	}
	if value < int64(goalValue) {
		return fmt.Errorf("%w, should be at least %d", errorValueIsNotMatchMinimum, goalValue)
	}

	return nil
}

func validateMax(ruleValue string, value int64) error {
	goalValue, err := strconv.Atoi(ruleValue)
	if err != nil {
		return err
	}
	if value > int64(goalValue) {
		return fmt.Errorf("%w, should be at most %d", errorValueIsNotMatchMaximum, goalValue)
	}

	return nil
}

func validateStringSlice(fieldValue []string, fieldName string, validationRules string) error {
	for _, item := range fieldValue {
		err := validateString(item, fieldName, validationRules)
		if err != nil {
			var validationErrors ValidationErrors
			if errors.As(err, &validationErrors) && len(validationErrors) == 0 {
				continue
			}
			return err
		}
	}
	return nil
}

func validateIntSlice(fieldValue []int, fieldName string, validationRules string) error {
	for _, item := range fieldValue {
		err := validateInt(int64(item), fieldName, validationRules)
		if err != nil {
			var validationErrors ValidationErrors
			if errors.As(err, &validationErrors) && len(validationErrors) == 0 {
				continue
			}
			return err
		}
	}
	return nil
}
