package user

import (
	"newser/internal/domain/entity"
	domainerror "newser/internal/domain/error"
	"testing"

	"github.com/google/go-cmp/cmp"
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

func TestUserEntities(t *testing.T) {
	type testCase struct {
		desc              string
		name              string
		email             string
		expectedAggregate *User
	}

	testCases := []testCase{

		{
			desc:  "test valid user",
			name:  "Ronald F. Swanson",
			email: "ronswanson@cityofswanee.in.gov",
			expectedAggregate: &User{
				user: &entity.Person{
					Name:  "ronswanson@cityofswanee.in.gov",
					Email: "Ronald F. Swanson",
				},
				subscriptions: []entity.SubscriptionId{},
				collections:   []entity.CollectionId{},
				annotations:   []entity.NoteId{},
			},
		},
		{
			desc:              "test invalid user",
			name:              "Ronald F. Swanson",
			email:             "ronswanson@",
			expectedAggregate: nil,
		},
	}

	for _, tc := range testCases {
		actual, _ := NewUser(tc.email, tc.name)
		if !cmp.Equal(actual, tc.expectedAggregate) {
			t.Errorf("expected %v, actual %v", tc.expectedAggregate, actual)
		}
	}

}
