package services

import (
	"errors"
	"example/layered-architecture/models"
	"example/layered-architecture/repositories"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestGetAllUsers(t *testing.T) {

	type testCase struct {
		name                      string
		returnUsersFromRepository *[]models.User
		returnErrorFromRepository error
		expectedError             error
	}
	testCases := []testCase{{name: "error", returnUsersFromRepository: nil,
		returnErrorFromRepository: errors.New("some error"),
		expectedError:             errors.New("some error")},
		{name: "success", returnUsersFromRepository: nil,
			returnErrorFromRepository: nil,
			expectedError:             nil}}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {

			mockRepository := repositories.NewMockRepositoryInterface(gomock.NewController(t))
			mockRepository.
				EXPECT().
				GetAllUsers().
				Return(test.returnUsersFromRepository, test.returnErrorFromRepository).
				Times(1)

			ms := NewUserService(mockRepository)

			_, err := ms.GetAllUsers()

			assert.Equal(t, err, test.expectedError)
		})
	}

}
