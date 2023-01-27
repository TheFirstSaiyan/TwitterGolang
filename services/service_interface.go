package services

import (
	"example/layered-architecture/models"
)

//go:generate mockgen --destination=./mock_service_interface.go --package=services example/layered-architecture/services ServiceInterface
type ServiceInterface interface {
	AddUser(user *models.User) error
	SignIn(user *models.User) error
	GetAllUsers() (*[]models.User, error)
	AddTweet(tweet *models.Tweet) error
	GetTweetsOfUser(username string) (*[]models.Tweet, error)
	GetFolloweesOfUser(username string) (*[]models.Follows, error)
	AddFollowee(follow *models.Follows) error
	DeleteTweet(tweetid int) error
	DeleteFollowee(username string, followeename string) error
	CheckFollowing(username string, followeename string) error
}
