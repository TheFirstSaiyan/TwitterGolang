package handlers

import (
	"testing"
)

func TestGetAllUsers(t *testing.T) {

	// t.Run("Error", func(t *testing.T) {
	// 	req, _ := http.NewRequest(http.MethodGet, "/api/user", http.NoBody)
	// 	rec := httptest.NewRecorder()

	// 	mockService := service.NewMockIMovieService(gomock.NewController(t))
	// 	mockService.
	// 		EXPECT().
	// 		GetMovies().
	// 		Return([]model.Movie{}, errors.New("oops!")).
	// 		Times(1)

	// 	mh := NewMovieHandler(mockService)

	// 	mh.GetMovies(rec, req, nil)

	// 	assert.Equal(t, http.StatusInternalServerError, rec.Code)
	// })

}
