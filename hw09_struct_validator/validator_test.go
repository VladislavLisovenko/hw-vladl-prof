package hw09structvalidator

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type UserRole string

// Test the function on different structures and other types.
type (
	User struct {
		ID     string `json:"id" validate:"len:36"`
		Name   string
		Age    int             `validate:"min:18|max:50"`
		Email  string          `validate:"regexp:^\\w+@\\w+\\.\\w+$"`
		Role   UserRole        `validate:"in:admin,stuff"`
		Phones []string        `validate:"len:11"`
		meta   json.RawMessage //nolint:unused
	}

	App struct {
		Version string `validate:"len:5"`
	}

	Token struct {
		Header    []byte
		Payload   []byte
		Signature []byte
	}

	Response struct {
		Code int    `validate:"in:200,404,500"`
		Body string `json:"omitempty"`
	}
)

func TestValidate(t *testing.T) {
	tests := []struct {
		descr       string
		in          interface{}
		expectedErr error
	}{
		{
			descr: "No errors, struct is valid",
			in: User{
				ID:     "123456789_123456789_123456789_123456",
				Name:   "Bob",
				Age:    30,
				Email:  "asd@asd.ei",
				Role:   "admin",
				Phones: []string{"01234567891", "01234567892"},
			},
			expectedErr: nil,
		},
		{
			descr: "ID length is not 36",
			in: User{
				ID:     "123456789_123456789_123456789",
				Name:   "Bob",
				Age:    30,
				Email:  "asd@asd.ei",
				Role:   "admin",
				Phones: []string{"01234567891", "01234567892"},
			},
			expectedErr: errorStringLength,
		},
		{
			descr: "Age doesn't belong to the interval - less then min",
			in: User{
				ID:     "123456789_123456789_123456789_123456",
				Name:   "Bob",
				Age:    15,
				Email:  "asd@asd.ei",
				Role:   "admin",
				Phones: []string{"01234567891", "01234567892"},
			},
			expectedErr: errorValueIsNotMatchMinimum,
		},
		{
			descr: "Age doesn't belong to the interval - more then max",
			in: User{
				ID:     "123456789_123456789_123456789_123456",
				Name:   "Bob",
				Age:    115,
				Email:  "asd@asd.ei",
				Role:   "admin",
				Phones: []string{"01234567891", "01234567892"},
			},
			expectedErr: errorValueIsNotMatchMaximum,
		},
		{
			descr: "Email is not match regular expression",
			in: User{
				ID:     "123456789_123456789_123456789_123456",
				Name:   "Bob",
				Age:    30,
				Email:  "@asd@asd.ei",
				Role:   "admin",
				Phones: []string{"01234567891", "01234567892"},
			},
			expectedErr: errorStringIsNotMatchRegexp,
		},
		{
			descr: "User role is not belong to set",
			in: User{
				ID:     "123456789_123456789_123456789_123456",
				Name:   "Bob",
				Age:    30,
				Email:  "asd@asd.ei",
				Role:   "radmin",
				Phones: []string{"01234567891", "01234567892"},
			},
			expectedErr: errorValeuIsNotInSet,
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
			tt := tt
			t.Parallel()
			//var validErrors ValidationErrors

			err := Validate(tt.in)
			if tt.expectedErr != nil {
				assert.Error(t, err)
				assert.ErrorAs(t, err, &ValidationErrors{})
			} else {
				require.NoError(t, err)
			}
		})
	}
}
