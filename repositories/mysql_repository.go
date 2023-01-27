package repositories

import (
	"errors"
	"example/layered-architecture/models"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type MySQLRepository struct {
	db *gorm.DB
}

func NewMySqlRepository(dsn string) *MySQLRepository {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})

	if err != nil {
		panic("cannot connect to DB!!")
	}
	err = db.AutoMigrate(&models.User{})
	if err != nil {
		panic("cannot initiate user table")
	}
	err = db.AutoMigrate(&models.Follows{})
	if err != nil {
		panic("cannot initiate followers table")
	}
	err = db.AutoMigrate(&models.Tweet{})
	if err != nil {
		panic("cannot initiate tweets table")
	}
	fmt.Println("connected to DB")
	return &MySQLRepository{db: db}

}

func (repository *MySQLRepository) AddUser(user *models.User) error {
	//create record in table
	err := repository.db.Create(user).Error
	return err
}

func (repository *MySQLRepository) SignIn(user *models.User) error {

	var signedinuser models.User
	rows := repository.db.Where("BINARY name = ? and password = ?", user.Name, user.Password).Find(&signedinuser).RowsAffected
	if rows != 1 {
		return errors.New("bad request")
	}
	return nil
}

func (repository *MySQLRepository) GetAllUsers() (*[]models.User, error) {

	var users []models.User
	//select all records from users
	err := repository.db.Find(&users).Error
	return &users, err
}

func (repository *MySQLRepository) AddTweet(tweet *models.Tweet) error {

	//check if user exists
	var user models.User
	rows := repository.db.Where("BINARY name = ?", tweet.UserName).Find(&user).RowsAffected
	if rows != 1 {
		return errors.New("bad request")
	}
	//tweet validation
	if len(tweet.Content) < 1 {
		return errors.New("bad request")
	}
	//craete the tweet and return json
	repository.db.Create(&tweet)
	return nil
}

func (repository *MySQLRepository) GetTweetsOfUser(username string) (*[]models.Tweet, error) {

	var tweets []models.Tweet
	err := repository.db.Where("BINARY user_name = ?", username).Find(&tweets).Error
	return &tweets, err
}

func (repository *MySQLRepository) GetFolloweesOfUser(username string) (*[]models.Follows, error) {

	var followees []models.Follows
	err := repository.db.Where("BINARY source_user = ?", username).Find(&followees).Error
	return &followees, err
}

func (repository *MySQLRepository) AddFollowee(follow *models.Follows) error {

	// check if user exists
	var user models.User
	rows := repository.db.Where("BINARY name = ?", follow.SourceUser).Find(&user).RowsAffected
	if rows != 1 {
		return errors.New("bad request")
	}

	var existing models.Follows

	//check if the user is already following
	rows = repository.db.Where("BINARY source_user = ? and target_user = ?", follow.SourceUser, follow.TargetUser).Find(&existing).RowsAffected
	if rows == 1 {
		return errors.New("bad request")
	}
	repository.db.Create(&follow)
	return nil
}

func (repository *MySQLRepository) DeleteTweet(tweetid int) error {
	var tweet models.Tweet
	err := repository.db.Delete(&tweet, tweetid).Error

	return err
}
func (repository *MySQLRepository) DeleteFollowee(username string, followeename string) error {
	var followee models.Follows
	err := repository.db.Delete(&followee, "BINARY source_user = ? and target_user = ?", username, followeename).Error
	return err

}

func (repository *MySQLRepository) CheckFollowing(username string, followeename string) error {
	// check if user exists
	var user models.User
	var existing models.Follows

	rows := repository.db.Where("BINARY name = ?", username).Find(&user).RowsAffected
	if rows != 1 {
		return errors.New("bad request")
	}
	//check number of rows returned for the 2 users
	rows = repository.db.Where("BINARY source_user = ? and target_user = ?", username, followeename).Find(&existing).RowsAffected
	if rows == 1 {
		return errors.New("bad request")
	}
	return nil
}
