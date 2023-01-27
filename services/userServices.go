package services

import (
	"example/layered-architecture/models"
	"example/layered-architecture/repositories"
)

type UserService struct {
	repository repositories.RepositoryInterface
}

func NewUserService(repository repositories.RepositoryInterface) *UserService {
	return &UserService{repository: repository}
}

func (service *UserService) AddUser(user *models.User) error {
	return service.repository.AddUser(user)
}

func (service *UserService) SignIn(user *models.User) error {
	return service.repository.SignIn(user)
}

func (service *UserService) GetAllUsers() *[]models.User {
	return service.repository.GetAllUsers()
}

func (service *UserService) AddTweet(tweet *models.Tweet) error {
	return service.repository.AddTweet(tweet)
}

func (service *UserService) GetTweetsOfUser(username string) *[]models.Tweet {
	return service.repository.GetTweetsOfUser(username)
}

func (service *UserService) GetFolloweesOfUser(username string) *[]models.Follows {
	return service.repository.GetFolloweesOfUser(username)
}

func (service *UserService) AddFollowee(follow *models.Follows) error {
	return service.repository.AddFollowee(follow)
}

func (service *UserService) DeleteTweet(tweetid int) {
	service.repository.DeleteTweet(tweetid)
}

func (service *UserService) DeleteFollowee(username string, followeename string) {
	service.repository.DeleteFollowee(username, followeename)
}

func (service *UserService) CheckFollowing(username string, followeename string) error {
	return service.repository.CheckFollowing(username, followeename)
}
