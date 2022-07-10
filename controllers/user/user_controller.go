package user_controller
//import
import (
	"bytes"
	"math/rand"
	"os"
	"time"

	"github.com/fops9311/mvc_server_app/app"
	"github.com/fops9311/mvc_server_app/model/controller"
	"github.com/fops9311/mvc_server_app/model/resource"
	"github.com/fops9311/mvc_server_app/model/server"
	"github.com/fops9311/mvc_server_app/utils/emailer"
	view "github.com/fops9311/mvc_server_app/views/user"
)                                           //import
var Resource resource.Resurce

var Index controller.Action = func(params map[string]interface{}) (result string, err error) {
	buf := bytes.NewBuffer([]byte{})
	err = view.Index(params, buf)
	return buf.String(), err
}
var Edit controller.Action = func(params map[string]interface{}) (result string, err error) {
	buf := bytes.NewBuffer([]byte{})
	err = view.Edit(params, buf)
	return buf.String(), err
}
var New controller.Action = func(params map[string]interface{}) (result string, err error) {
	buf := bytes.NewBuffer([]byte{})
	err = view.New(params, buf)
	return buf.String(), err
}
var Show controller.Action = func(params map[string]interface{}) (result string, err error) {
	buf := bytes.NewBuffer([]byte{})
	err = view.Show(params, buf)
	return buf.String(), err
}
var Create controller.Action = func(params map[string]interface{}) (result string, err error) {
	buf := bytes.NewBuffer([]byte{})
	err = view.Create(params, buf)
	return buf.String(), err
}
var Update controller.Action = func(params map[string]interface{}) (result string, err error) {
	buf := bytes.NewBuffer([]byte{})
	err = view.Update(params, buf)
	return buf.String(), err
}
var Delete controller.Action = func(params map[string]interface{}) (result string, err error) {
	buf := bytes.NewBuffer([]byte{})
	err = view.Delete(params, buf)
	return buf.String(), err
}

func Init() {
	view.Init()
	init_begin()
	Resource = resource.NewResource()
	Resource.Key = "/users"
	Resource.Actions["Index"] = &resource.ActionPath{ 	
		Verb:       "GET",
		Path:       "",
		Middleware: make([]string, 0),
		Action:     Index,
	}
	Resource.Actions["Edit"] = &resource.ActionPath{ 	
		Verb:       "GET",
		Path:       "/"+server.URIParam("user_id")+"/edit",
		Middleware: make([]string, 0),
		Action:     Edit,
	}
	Resource.Actions["New"] = &resource.ActionPath{ 	
		Verb:       "GET",
		Path:       "/new",
		Middleware: make([]string, 0),
		Action:     New,
	}
	Resource.Actions["Show"] = &resource.ActionPath{ 	
		Verb:       "GET",
		Path:       "/"+server.URIParam("user_id"),
		Middleware: make([]string, 0),
		Action:     Show,
	}
	Resource.Actions["Create"] = &resource.ActionPath{ 	
		Verb:       "POST",
		Path:       "",
		Middleware: make([]string, 0),
		Action:     Create,
	}
	Resource.Actions["Update"] = &resource.ActionPath{ 	
		Verb:       "PUT",
		Path:       "/"+server.URIParam("user_id"),
		Middleware: make([]string, 0),
		Action:     Update,
	}
	Resource.Actions["Delete"] = &resource.ActionPath{ 	
		Verb:       "DELETE",
		Path:       "/"+server.URIParam("user_id"),
		Middleware: make([]string, 0),
		Action:     Delete,
	}

	init_continue()
}

//!!define init_begin func(){}
//!!define init_continue func(){}
//DO NOT CHANGE ABOVE --GENERATED--
func init_begin() {
	Create = func(params map[string]interface{}) (result string, err error) {
		var email string = ""
		var password string = ""
		responce := bytes.NewBuffer([]byte{})
		emailpayload := bytes.NewBuffer([]byte{})
		switch v := params["email"].(type) {
		case []string:
			if len(v) < 1 {
				err = view.RegisterError(params, responce)
				return responce.String(), err
			}
			email = v[0]
		default:
			err = view.RegisterError(params, responce)
			return responce.String(), err
		}

		params["email_confirm_route"] = os.Getenv("EMAIL_CONFIRM_ROUTE") //"http://localhost:8000/v1/users"
		params["user_id"] = email
		switch v := params["password"].(type) {
		case []string:
			if len(v) < 1 {
				err = view.RegisterError(params, responce)
				return responce.String(), err
			}
			password = v[0]
		default:
			err = view.RegisterError(params, responce)
			return responce.String(), err
		}

		var email_confirm_param = RandStringRunes(60)
		params["email_confirm_param"] = email_confirm_param
		err = app.Users.NewUser(app.User{
			Id:                email,
			Email:             email,
			Password:          password,
			EmailConfirmed:    false,
			EmailConfirmParam: email_confirm_param,
		})
		if err != nil {
			err = view.InternalServerError(params, responce)
			return responce.String(), err
		}
		err = view.RegisterEmailPayload(params, emailpayload)
		if err != nil {
			err = view.InternalServerError(params, responce)
			return responce.String(), err
		}
		err = emailer.Send([]string{email}, "Registration", emailpayload.String())
		if err != nil {
			err = view.RegisterSendError(params, responce)
			return responce.String(), err
		}

		err = view.RegisterNeedConfirm(params, responce)
		return responce.String(), err
	}
}

var EmailConfirm controller.Action = func(params map[string]interface{}) (result string, err error) {
	var user_id string
	var email_confirm_param string
	switch v := params["user_id"].(type) {
	case string:
		user_id = v
	default:
		return "NOT confirmed", err
	}
	switch v := params["email_confirm_param"].(type) {
	case string:
		email_confirm_param = v
	default:
		return "NOT confirmed", err
	}
	err = app.Users.ConfirmEmail(user_id, email_confirm_param)
	if err != nil {
		return "NOT confirmed", err
	}
	return "Confirmed", err
}

func init_continue() {
	view.Init()
	Resource.Actions["EmailConfirm"] = &resource.ActionPath{
		Verb:       "GET",
		Path:       "/" + server.URIParam("user_id") + "/" + server.URIParam("email_confirm_param"),
		Middleware: make([]string, 0),
		Action:     EmailConfirm,
	}
	Resource.Actions["Edit"].Middleware = append(Resource.Actions["Edit"].Middleware, "BasicUserAuth")
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes(n int) string {
	rand.Seed(time.Now().UnixNano())
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
