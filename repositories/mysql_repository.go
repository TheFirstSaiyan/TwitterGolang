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

func (repository *MySQLRepository) GetAllUsers() *[]models.User {

	var users []models.User
	//select all records from users
	repository.db.Find(&users)
	return &users
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

func (repository *MySQLRepository) GetTweetsOfUser(username string) *[]models.Tweet {

	var tweets []models.Tweet
	repository.db.Where("BINARY user_name = ?", username).Find(&tweets)
	return &tweets
}

func (repository *MySQLRepository) GetFolloweesOfUser(username string) *[]models.Follows {

	var followees []models.Follows
	repository.db.Where("BINARY source_user = ?", username).Find(&followees)
	return &followees
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

func (repository *MySQLRepository) DeleteTweet(tweetid int) {
	var tweet models.Tweet
	repository.db.Delete(&tweet, tweetid)
}
func (repository *MySQLRepository) DeleteFollowee(username string, followeename string) {
	var followee models.Follows
	repository.db.Delete(&followee, "BINARY source_user = ? and target_user = ?", username, followeename)

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
