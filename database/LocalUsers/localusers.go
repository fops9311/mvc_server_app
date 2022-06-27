package localusers

import (
	"strings"
	"sync"

	"github.com/fops9311/mvc_server_app/app"
)

func DefineComponent() {
	app.Users.NewUser = NewUser
	app.Users.ConfirmEmail = ConfirmEmail
}

var Users map[string]*app.User = map[string]*app.User{}
var m sync.Mutex

var NewUser = func(user app.User) (err error) {
	m.Lock()
	defer m.Unlock()
	if _, ok := Users[user.Id]; !ok || !user.EmailConfirmed {
		Users[user.Id] = &user
		return nil
	}
	if !user.EmailConfirmed {
		Users[user.Id] = &user
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
		return nil
	}
	err = app.ErrUserEmailConfirmationFail
	return err
}
