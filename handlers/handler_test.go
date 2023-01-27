package handlers

import (
	"bytes"
	"encoding/json"
	"errors"
	"example/layered-architecture/models"
	"example/layered-architecture/services"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestGetAllUsers(t *testing.T) {

	type testCase struct {
		name                   string
		returnUsersFromService *[]models.User
		returnErrorFromService error
		expectedStatusCode     int
	}
	testCases := []testCase{{name: "error", returnUsersFromService: nil,
		returnErrorFromService: errors.New("some error"),
		expectedStatusCode:     http.StatusBadRequest},
		{name: "success", returnUsersFromService: nil,
			returnErrorFromService: nil,
			expectedStatusCode:     http.StatusOK}}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			req, _ := http.NewRequest(http.MethodGet, "/api/user", http.NoBody)
			res := httptest.NewRecorder()

			mockService := services.NewMockServiceInterface(gomock.NewController(t))
			mockService.
				EXPECT().
				GetAllUsers().
				Return(test.returnUsersFromService, test.returnErrorFromService).
				Times(1)

			mh := NewHandler(mockService)

			mh.GetAllUsers(res, req)

			assert.Equal(t, test.expectedStatusCode, res.Code)
		})
	}

}

func TestAddUser(t *testing.T) {

	type testCase struct {
		name                   string
		returnErrorFromService error
		expectedStatusCode     int
		requestBody            *models.User
	}
	testCases := []testCase{{name: "error",
		returnErrorFromService: errors.New("some error"),
		expectedStatusCode:     http.StatusBadRequest,
		requestBody: &models.User{
			Model:    gorm.Model{},
			Name:     "abc",
			Password: "ffdd"}},
		{name: "success",
			returnErrorFromService: nil,
			expectedStatusCode:     http.StatusOK,
			requestBody: &models.User{
				Model:    gorm.Model{},
				Name:     "abc",
				Password: "password"}}}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			body, _ := json.Marshal(test.requestBody)
			req, _ := http.NewRequest(http.MethodPost, "/api/user", bytes.NewBuffer(body))
			res := httptest.NewRecorder()
			mockService := services.NewMockServiceInterface(gomock.NewController(t))
			mockService.
				EXPECT().
				AddUser(test.requestBody).
				Return(test.returnErrorFromService).
				Times(1)

			mh := NewHandler(mockService)

			mh.AddUser(res, req)

			assert.Equal(t, test.expectedStatusCode, res.Code)
		})
	}

}

func TestSignIn(t *testing.T) {
	type testCase struct {
		name                   string
		returnErrorFromService error
		expectedStatusCode     int
		requestBody            *models.User
	}
	testCases := []testCase{{name: "error",
		returnErrorFromService: errors.New("some error"),
		expectedStatusCode:     http.StatusUnauthorized,
		requestBody: &models.User{
			Model:    gorm.Model{},
			Name:     "abc",
			Password: "ffdd"}},
		{name: "success",
			returnErrorFromService: nil,
			expectedStatusCode:     http.StatusOK,
			requestBody: &models.User{
				Model:    gorm.Model{},
				Name:     "abc",
				Password: "password"}}}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			body, _ := json.Marshal(test.requestBody)
			req, _ := http.NewRequest(http.MethodPost, "/api/user", bytes.NewBuffer(body))
			res := httptest.NewRecorder()
			mockService := services.NewMockServiceInterface(gomock.NewController(t))
			mockService.
				EXPECT().
				SignIn(test.requestBody).
				Return(test.returnErrorFromService).
				Times(1)

			mh := NewHandler(mockService)

			mh.SignIn(res, req)

			assert.Equal(t, test.expectedStatusCode, res.Code)
		})
	}

}

func TestAddTweet(t *testing.T) {

	type testCase struct {
		name                   string
		returnErrorFromService error
		expectedStatusCode     int
		requestBody            *models.Tweet
	}
	testCases := []testCase{{name: "error",
		returnErrorFromService: errors.New("some error"),
		expectedStatusCode:     http.StatusBadRequest,
		requestBody: &models.Tweet{
			Model:    gorm.Model{},
			UserName: "abc",
			Content:  ""}},
		{name: "success",
			returnErrorFromService: nil,
			expectedStatusCode:     http.StatusOK,
			requestBody: &models.Tweet{
				Model:    gorm.Model{},
				UserName: "abc",
				Content:  "ffdd"}}}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			body, _ := json.Marshal(test.requestBody)
			req, _ := http.NewRequest(http.MethodPost, "/api/user", bytes.NewBuffer(body))
			res := httptest.NewRecorder()
			mockService := services.NewMockServiceInterface(gomock.NewController(t))
			mockService.
				EXPECT().
				AddTweet(test.requestBody).
				Return(test.returnErrorFromService).
				Times(1)

			mh := NewHandler(mockService)

			mh.AddTweet(res, req)

			assert.Equal(t, test.expectedStatusCode, res.Code)
		})
	}

}

func TestGetTweetsOfuser(t *testing.T) {
	type testCase struct {
		name                      string
		expectedStatusCode        int
		paramUsername             string
		returnedTweetsFromService *[]models.Tweet
		returnedErrorFromService  error
	}
	testCases := []testCase{{name: "error",
		expectedStatusCode:        http.StatusBadRequest,
		paramUsername:             "abcd",
		returnedTweetsFromService: &[]models.Tweet{},
		returnedErrorFromService:  errors.New("some error")},
		{name: "success",
			expectedStatusCode:        http.StatusOK,
			paramUsername:             "abc",
			returnedTweetsFromService: &[]models.Tweet{},
			returnedErrorFromService:  nil}}

	for _, test := range testCases {

		t.Run(test.name, func(t *testing.T) {

			req, _ := http.NewRequest(http.MethodGet, "/api/user/tweets/"+test.paramUsername, http.NoBody)

			res := httptest.NewRecorder()
			//Hack to try to fake gorilla/mux vars
			vars := map[string]string{
				"username": test.paramUsername,
			}

			// CHANGE THIS LINE!!!
			req = mux.SetURLVars(req, vars)
			mockService := services.NewMockServiceInterface(gomock.NewController(t))
			mockService.
				EXPECT().
				GetTweetsOfUser(test.paramUsername).
				Return(test.returnedTweetsFromService, test.returnedErrorFromService).
				Times(1)

			mh := NewHandler(mockService)

			mh.GetTweetsOfUser(res, req)

			assert.Equal(t, test.expectedStatusCode, res.Code)
		})
	}
}
func TestGetFolloweesOfUser(t *testing.T) {
	type testCase struct {
		name                       string
		expectedStatusCode         int
		paramUsername              string
		returnedFollowsFromService *[]models.Follows
		returnedErrorFromService   error
	}
	testCases := []testCase{{name: "error",
		expectedStatusCode:         http.StatusBadRequest,
		paramUsername:              "abcd",
		returnedFollowsFromService: &[]models.Follows{},
		returnedErrorFromService:   errors.New("some error")},
		{name: "success",
			expectedStatusCode:         http.StatusOK,
			paramUsername:              "abc",
			returnedFollowsFromService: &[]models.Follows{},
			returnedErrorFromService:   nil}}

	for _, test := range testCases {

		t.Run(test.name, func(t *testing.T) {

			req, _ := http.NewRequest(http.MethodGet, "/api/user/followees/", http.NoBody)

			res := httptest.NewRecorder()
			vars := map[string]string{
				"username": test.paramUsername,
			}
			req = mux.SetURLVars(req, vars)
			mockService := services.NewMockServiceInterface(gomock.NewController(t))
			mockService.
				EXPECT().
				GetFolloweesOfUser(test.paramUsername).
				Return(test.returnedFollowsFromService, test.returnedErrorFromService).
				Times(1)

			mh := NewHandler(mockService)

			mh.GetFolloweesOfUser(res, req)

			assert.Equal(t, test.expectedStatusCode, res.Code)
		})
	}
}

func TestCheckFollowing(t *testing.T) {
	type testCase struct {
		name                     string
		expectedStatusCode       int
		paramUsername            string
		paramFolloweename        string
		returnedErrorFromService error
	}
	testCases := []testCase{
		{name: "already following",
			expectedStatusCode:       http.StatusFound,
			paramUsername:            "abcd",
			paramFolloweename:        "fdfd",
			returnedErrorFromService: errors.New("some error")},
		{name: "success",
			expectedStatusCode:       http.StatusNotFound,
			paramUsername:            "abc",
			paramFolloweename:        "dfdfdf",
			returnedErrorFromService: nil}}

	for _, test := range testCases {

		t.Run(test.name, func(t *testing.T) {

			req, _ := http.NewRequest(http.MethodGet, "/api/user/followees/", http.NoBody)

			res := httptest.NewRecorder()
			vars := map[string]string{
				"username":     test.paramUsername,
				"followeename": test.paramFolloweename}
			req = mux.SetURLVars(req, vars)
			mockService := services.NewMockServiceInterface(gomock.NewController(t))
			mockService.
				EXPECT().
				CheckFollowing(test.paramUsername, test.paramFolloweename).
				Return(test.returnedErrorFromService).
				Times(1)

			mh := NewHandler(mockService)

			mh.CheckFollowing(res, req)

			assert.Equal(t, test.expectedStatusCode, res.Code)
		})
	}
}

func TestAddFollowee(t *testing.T) {
	type testCase struct {
		name                     string
		expectedStatusCode       int
		returnedErrorFromService error
		requestBody              *models.Follows
	}
	testCases := []testCase{{name: "error",
		expectedStatusCode:       http.StatusBadRequest,
		returnedErrorFromService: errors.New("some error"),
		requestBody: &models.Follows{
			Model:      gorm.Model{},
			SourceUser: "abc",
			TargetUser: "ffdd"}},
		{name: "success",
			expectedStatusCode:       http.StatusOK,
			returnedErrorFromService: nil,
			requestBody: &models.Follows{
				Model:      gorm.Model{},
				SourceUser: "abc",
				TargetUser: "ffdd"}}}

	for _, test := range testCases {

		t.Run(test.name, func(t *testing.T) {
			body, _ := json.Marshal(test.requestBody)

			req, _ := http.NewRequest(http.MethodPost, "/api/follow/", bytes.NewBuffer(body))

			res := httptest.NewRecorder()
			mockService := services.NewMockServiceInterface(gomock.NewController(t))
			mockService.
				EXPECT().
				AddFollowee(test.requestBody).
				Return(test.returnedErrorFromService).
				Times(1)

			mh := NewHandler(mockService)

			mh.AddFollowee(res, req)

			assert.Equal(t, test.expectedStatusCode, res.Code)
		})
	}
}

func TestDeleteTweet(t *testing.T) {
	type testCase struct {
		name                     string
		expectedStatusCode       int
		returnedErrorFromService error
		paramId                  string
	}
	testCases := []testCase{{name: "error",
		expectedStatusCode:       http.StatusBadRequest,
		returnedErrorFromService: errors.New("some error"),
		paramId:                  "1"},
		{name: "success",
			expectedStatusCode:       http.StatusOK,
			returnedErrorFromService: nil,
			paramId:                  "2"}}

	for _, test := range testCases {

		t.Run(test.name, func(t *testing.T) {

			req, _ := http.NewRequest(http.MethodDelete, "/api/tweet/", http.NoBody)

			res := httptest.NewRecorder()
			vars := map[string]string{
				"tweetid": test.paramId}
			req = mux.SetURLVars(req, vars)
			val, _ := strconv.Atoi(test.paramId)
			mockService := services.NewMockServiceInterface(gomock.NewController(t))
			mockService.
				EXPECT().
				DeleteTweet(val).
				Return(test.returnedErrorFromService).
				Times(1)

			mh := NewHandler(mockService)

			mh.DeleteTweet(res, req)

			assert.Equal(t, test.expectedStatusCode, res.Code)
		})
	}
}

func TestDeleteFollowee(t *testing.T) {
	type testCase struct {
		name                     string
		expectedStatusCode       int
		returnedErrorFromService error
		paramSourceUsername      string
		paramTargetUsername      string
	}
	testCases := []testCase{{name: "error",
		expectedStatusCode:       http.StatusBadRequest,
		returnedErrorFromService: errors.New("some error"),
		paramSourceUsername:      "adfd",
		paramTargetUsername:      "fdfdf"},
		{name: "success",
			expectedStatusCode:       http.StatusOK,
			returnedErrorFromService: nil,
			paramSourceUsername:      "adfd",
			paramTargetUsername:      "fdfdf"}}

	for _, test := range testCases {

		t.Run(test.name, func(t *testing.T) {

			req, _ := http.NewRequest(http.MethodDelete, "/api/user/followees", http.NoBody)

			res := httptest.NewRecorder()
			vars := map[string]string{
				"username":     test.paramSourceUsername,
				"followeename": test.paramTargetUsername}
			req = mux.SetURLVars(req, vars)
			mockService := services.NewMockServiceInterface(gomock.NewController(t))
			mockService.
				EXPECT().
				DeleteFollowee(test.paramSourceUsername, test.paramTargetUsername).
				Return(test.returnedErrorFromService).
				Times(1)

			mh := NewHandler(mockService)

			mh.DeleteFollowee(res, req)

			assert.Equal(t, test.expectedStatusCode, res.Code)
		})
	}
}
