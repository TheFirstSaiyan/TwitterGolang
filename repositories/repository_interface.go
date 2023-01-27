package repositories

import (
	"example/layered-architecture/models"
)

//go:generate mockgen --destination=./mock_repository_interface.go --package=repositories example/layered-architecture/repositories RepositoryInterface
type RepositoryInterface interface {
	AddUser(user *models.User) error
	SignIn(user *models.User) error
	GetAllUsers() *[]models.User
	AddTweet(tweet *models.Tweet) error
	GetTweetsOfUser(username string) *[]models.Tweet
	GetFolloweesOfUser(username string) *[]models.Follows
	AddFollowee(follow *models.Follows) error
	DeleteTweet(tweetid int)
	DeleteFollowee(username string, followeename string)
	CheckFollowing(username string, followeename string) error
}