package app

import (
	"errors"
	"time"
)

var ErrUserNotExist error = errors.New("UserNotExist")
var ErrUserAlreadyExist error = errors.New("UserAlreadyExist")
var ErrUserComponentNotDefined error = errors.New("UserComponentNotDefined")
var ErrUserEmailConfirmationFail error = errors.New("UserEmailConfirmationFail")

var Users = struct {
	GetUser        func(user_id string) (user User, err error)
	NewUser        func(user User) (err error)
	ConfirmEmail   func(user_id, confirm_email_param string) (err error)
	Authentication func(login string, password string) bool
}{
	GetUser: func(user_id string) (user User, err error) {
		return User{}, ErrUserComponentNotDefined
	},
	NewUser: func(user User) (err error) {
		return ErrUserComponentNotDefined
	},
	ConfirmEmail: func(user_id, confirm_email_param string) (err error) {
		return ErrUserComponentNotDefined
	},
	Authentication: func(login, password string) bool { return false },
}

type User struct {
	Id                string
	Email             string
	Password          string
	EmailConfirmed    bool
	EmailConfirmParam string
}

var ObjectNotExist error
var ObjectAlreadyExist error
var ObjectComponentNotDefined error

var Objects = struct {
	GetObject     func(object_id string, starttime time.Time, endtime time.Time, sampleCount int) (object Object, err error)
	AddSample     func(object_id string, value float32) (err error)
	Authorization func(user_id string, object_id string) bool
}{
	GetObject: func(object_id string, starttime time.Time, endtime time.Time, sampleCount int) (object Object, err error) {
		return Object{}, ObjectComponentNotDefined
	},
	AddSample: func(object_id string, value float32) (err error) {
		return ObjectComponentNotDefined
	},
	Authorization: func(user_id, object_id string) bool { return false },
}

type Object struct {
	Id       string
	Children []Object
	Samples  []Sample
}
type Sample struct {
	Timestamp time.Time
	Value     float32
}
