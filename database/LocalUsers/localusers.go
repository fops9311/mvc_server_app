package localusers

import (
	"crypto/subtle"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/fops9311/mvc_server_app/app"
	"github.com/golang-jwt/jwt"
)

func DefineComponent() {
	app.Users.NewUser = NewUser
	app.Users.ConfirmEmail = ConfirmEmail
	app.Users.Authentication = Authentication
	app.Users.GetJWT = GetJWT
}
func init() {
	fmt.Println("Using LocalUsers db...")
	readUsersFile()
}
func readUsersFile() {
	usersfile, err := os.ReadFile("./users.json")
	if err != nil {
		return
	}
	json.Unmarshal(usersfile, &Users)
}
func writeUsersFile() {
	b, err := json.MarshalIndent(Users, "", "	")
	if err != nil {
		return
	}
	os.WriteFile("users.json", b, 0666)
}

var Users map[string]*app.User = map[string]*app.User{}
var m sync.Mutex

var NewUser = func(user app.User) (err error) {
	m.Lock()
	defer m.Unlock()
	if _, ok := Users[user.Id]; !ok || !user.EmailConfirmed {
		Users[user.Id] = &user
		writeUsersFile()
		return nil
	}
	if !user.EmailConfirmed {
		Users[user.Id] = &user
		writeUsersFile()
		return nil
	}
	return app.ErrUserAlreadyExist
}
var ConfirmEmail = func(user_id string, confirm_email_param string) (err error) {
	m.Lock()
	defer m.Unlock()
	if _, ok := Users[user_id]; !ok {
		err = app.ErrUserNotExist
		return err
	}
	if strings.Compare(Users[user_id].EmailConfirmParam, confirm_email_param) == 0 {
		Users[user_id].EmailConfirmed = true
		writeUsersFile()
		return nil
	}
	err = app.ErrUserEmailConfirmationFail
	return err
}
var Authentication func(login string, password string) bool = func(login string, password string) bool {
	m.Lock()
	defer m.Unlock()
	for _, val := range Users {
		if subtle.ConstantTimeCompare([]byte(login), []byte(val.Id)) == 1 &&
			subtle.ConstantTimeCompare([]byte(password), []byte(val.Password)) == 1 &&
			val.EmailConfirmed {
			return true
		}

	}
	return false
}
var GetJWT func(login string, password string) (string, error) = func(login string, password string) (string, error) {
	var claims *jwtCustomClaims = &jwtCustomClaims{
		"Guest",
		false,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}

	if app.Users.Authentication(login, password) {
		claims = &jwtCustomClaims{
			"Jon Snow",
			false,
			jwt.StandardClaims{
				ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
			},
		}
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return "", err
	}
	return t, nil
}

type jwtCustomClaims struct {
	Name  string `json:"name"`
	Admin bool   `json:"admin"`
	jwt.StandardClaims
}
