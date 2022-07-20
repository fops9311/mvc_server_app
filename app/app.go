package app

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
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
	GetJWT         func(login string, password string) (string, error)
	VerifyJWT      func(*jwt.Token) bool
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

	GetJWT: func(login string, password string) (string, error) {
		return "{}", nil
	},
	VerifyJWT: func(*jwt.Token) bool {
		return false
	},
}

type User struct {
	Id                string
	Email             string
	Password          string
	EmailConfirmed    bool
	EmailConfirmParam string
	BasicAuthLogout   bool
}

var ObjectNotExist error
var ObjectAlreadyExist error
var ObjectComponentNotDefined error

var Objects = struct {
	AddObject  func(obj Object) (err error)
	GetObjects func(object_id string) []Object
	GetObject  func(object_id string, starttime time.Time, endtime time.Time, sampleCount int) (object Object, err error)
	AddSample  func(object_id string, value float32) (err error)
}{
	GetObject: func(object_id string, starttime time.Time, endtime time.Time, sampleCount int) (object Object, err error) {
		return Object{}, ObjectComponentNotDefined
	},
	AddSample: func(object_id string, value float32) (err error) {
		return ObjectComponentNotDefined
	},
	AddObject: func(obj Object) (err error) {
		return ObjectComponentNotDefined
	},
	GetObjects: func(object_id string) []Object {
		return []Object{}
	},
}

type Object struct {
	Id         string
	Samples    []Sample
	LastSample Sample
}
type Sample struct {
	Timestamp string
	Value     float32
}

var BasicAuthUserValidator func(login, pass string) (bool, error) = func(login, pass string) (bool, error) {
	return Users.Authentication(login, pass), nil
}
