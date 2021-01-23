package user

import (
	_ "github.com/sirupsen/logrus"
	u "github.com/0B1t322/distanceLearningWebSite/pkg/models/user"
	"gorm.io/gorm"
)

type UserController struct {
	db *gorm.DB
}

func NewUserController(db *gorm.DB) *UserController {
	return &UserController{db}
}

// GetUserByID return a user for bd
// 	errors:
// 		ErrUserNotFound
// 		gorms errors
func (c *UserController) GetUserByID(id string) (*u.User, error) {
	var user u.User
	err := c.db.First(&user, "id = ?", id).Error
	if err == gorm.ErrRecordNotFound {
		return nil, ErrUserNotFound
	} else if err != nil {
		return nil, err
	}

	return &user, nil
}

//GetUserByUserName return a user from bd
// 	errors:
// 		ErrUserNotFound
// 		gorms errors
func (c *UserController) GetUserByUserName(username string) (*u.User, error) {
	var user u.User
	err := c.db.First(&user,"username = ?", username).Error
	if err == gorm.ErrRecordNotFound {
		return nil, ErrUserNotFound
	} else if err != nil {
		return nil, err
	}

	return &user, nil
}

// AddUser to database
func (c *UserController) AddUser(model *u.User) error {
	// // check if user with this username is exsist
	// defer func() error {
	// 	sqlDB, err := c.db.DB()
	// 	if err != nil {
	// 		log.Warn(err)
	// 		return err
	// 	}

	// 	return sqlDB.Close()
	// }()
	
	var user u.User
	err := c.db.First(&user, "username = ?", model.Username).Error
	if err == nil {
		return ErrUserExsist
	}
	
	err = c.db.Create(model).Error
	if err != nil {
		return err
	}

	return nil
}

// DeleteUser from db
func (c *UserController) DeleteUser(model *u.User) error {
	// TODO подумать как удалять все же по ID
	user, err := c.GetUserByUserName(model.Username)
	if err != nil {
		return err
	}

	err = c.db.Delete(model,"username = ?", user.Username).Error
	if err != nil {
		return err
	}

	return nil
}