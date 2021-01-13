package user

import (
	"github.com/0B1t322/distanceLearningWebSite/pkg/db"
	"gorm.io/gorm"
)

// User is a model of user
type User struct {
	ID			int64
	Username 	string
	Password 	string
	Role 		string
}

func copyUser(u User) *User {
	return &User{
		ID: u.ID, 
		Username: u.Username, 
		Password: u.Password, 
		Role: u.Role,
	}
}

// NewUser Return a new User
func NewUser(username, password, role string) *User {
	return &User{
			Username: username,
			Password: password,
			Role: role,
		}
}

// GetUserByID return a user for bd
// 	errors:
// 		ErrUserNotFound
// 		gorms errors
func GetUserByID(id string) (*User, error) {
	db, err := db.GormOpen()
	if err != nil {
		return nil, err
	}
	
	var user User
	err = db.First(&user, "id = ?", id).Error
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
func GetUserByUserName(username string) (*User, error) {
	db, err := db.GormOpen()
	if err != nil {
		return nil, err
	}

	var user User
	err = db.First(&user,"username = ?", username).Error
	if err == gorm.ErrRecordNotFound {
		return nil, ErrUserNotFound
	} else if err != nil {
		return nil, err
	}

	return &user, nil
}

// AddUser to database
func (u *User) AddUser() error {
	db, err := db.GormOpen()
	if err != nil {
		return err
	}
	// check if user with this username is exsist
	var user User
	err = db.First(&user, "username = ?", u.Username).Error
	if err == nil {
		return ErrUserExsist
	}
	
	err = db.Create(u).Error
	if err != nil {
		return err
	}

	return nil
}

// DeleteUser from db
func (u *User) DeleteUser() error {
	db, err := db.GormOpen()
	if err != nil {
		return err
	}
	// TODO подумать как удалять все же по ID
	user, err := GetUserByUserName(u.Username)
	if err != nil {
		return err
	}

	err = db.Delete(u,"username = ?", user.Username).Error
	if err != nil {
		return err
	}

	return nil
}

// TODO преобразовывать модель к виду для pb