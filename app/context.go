package app

import (
	"net/http"

	"github.com/jiverson/go-rest-demo/db"
	"github.com/jiverson/go-rest-demo/model"
	"github.com/sirupsen/logrus"
)

type Context struct {
	Logger        logrus.FieldLogger
	RemoteAddress string
	Database      *db.Database
	User          *model.User
}

func (ctx *Context) WithLogger(logger logrus.FieldLogger) *Context {
	ret := *ctx
	ret.Logger = logger
	return &ret
}

func (ctx *Context) WithRemoteAddress(address string) *Context {
	ret := *ctx
	ret.RemoteAddress = address
	return &ret
}

func (ctx *Context) WithUser(user *model.User) *Context {
	ret := *ctx
	ret.User = user
	return &ret
}

func (ctx *Context) AuthorizationError() *UserError {
	return &UserError{Message: "unauthorized", StatusCode: http.StatusForbidden}
}

// WithBearerToken authenticates via JWT id token.
// func (ctx *Context) WithBearerToken(token string) (*Context, error) {
//     emailAddress, failureReason, err := ctx.App.ValidateBearerToken(token)
//     if err != nil {
//         return nil, err
//     } else if emailAddress == "" {
//         ctx.Logger.WithField("reason", failureReason).Info("bearer token validation failure")
//         return nil, nil
//     }
//     user, err := ctx.Store.GetUserByEmailAddress(emailAddress)
//     if err != nil {
//         return nil, err
//     }
//     return ctx.WithUser(user)
// }
