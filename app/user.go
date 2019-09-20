package app

import (
	"strings"

	"github.com/jiverson/go-rest-demo/model"
)

func (a *App) GetUserByEmail(email string) (*model.User, error) {
	return a.Database.GetUserByEmail(email)
}

func (ctx *Context) CreateUser(user *model.User, password string) error {
	if err := ctx.validateUser(user, password); err != nil {
		return err
	}

	if err := user.SetPassword(password); err != nil {
		return err
	}

	return ctx.Database.CreateUser(user)
}

func (ctx *Context) validateUser(user *model.User, password string) *ValidationError {
	if !strings.Contains(user.Email, "@") {
		return &ValidationError{"invalid email"}
	}

	if password == "" {
		return &ValidationError{"password is required"}
	}

	return nil
}
