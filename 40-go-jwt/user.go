package main

import (
	"context"
	"net/http"

	"github.com/pkg/errors"
)

// User holds user-related info
type User struct {
	Name     string `json:"name"`
	ID       string `json:"id"`
	IP       string `json:"ip,omitempty"`
	Admin    bool   `json:"admin"`
	Blocked  bool   `json:"block,omitempty"`
	Verified bool   `json:"verified,omitempty"`
}

type contextKey string

// MustGetUserInfo fails if can't extract user data from the request.
// should be called from authed controllers only
func MustGetUserInfo(r *http.Request) User {
	user, err := GetUserInfo(r)
	if err != nil {
		panic(err)
	}
	return user
}

// GetUserInfo returns user from request context
func GetUserInfo(r *http.Request) (user User, err error) {

	ctx := r.Context()
	if ctx == nil {
		return User{}, errors.New("no info about user")
	}
	if u, ok := ctx.Value(contextKey("user")).(User); ok {
		return u, nil
	}

	return User{}, errors.New("user can't be parsed")
}

// SetUserInfo sets user into request context
func SetUserInfo(r *http.Request, user User) *http.Request {
	ctx := r.Context()
	ctx = context.WithValue(ctx, contextKey("user"), user)
	return r.WithContext(ctx)
}
