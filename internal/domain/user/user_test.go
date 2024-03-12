package user

import (
	domainerror "newser/internal/domain/error"
	"testing"
)

func TestUserValidation(t *testing.T) {
	type testCase struct {
		desc          string
		name          string
		email         string
		expectedError error
	}

	testCases := []testCase{
		{
			desc:          "Test valid name and email",
			name:          "Bilbo Baggins",
			email:         "bilbobaggins@theshire.net",
			expectedError: nil,
		},
		{
			desc:          "Test valid name and invalid email",
			name:          "Bilbo Baggins",
			email:         "bilbobaggins@",
			expectedError: domainerror.ErrInvalidValue,
		},
		{
			desc:          "Test invalid name and valid email",
			name:          "",
			email:         "bilbobaggins@theshite.net",
			expectedError: domainerror.ErrInvalidValue,
		},
	}

	for _, tc := range testCases {
		_, err := NewUser(tc.email, tc.name)
		if err != tc.expectedError {
			t.Errorf("expected error: %v, got %v", tc.expectedError, err)
		}
	}
}
