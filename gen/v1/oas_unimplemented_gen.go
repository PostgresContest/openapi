// Code generated by ogen, DO NOT EDIT.

package v1

import (
	"context"

	ht "github.com/ogen-go/ogen/http"
)

// UnimplementedHandler is no-op Handler which returns http.ErrNotImplemented.
type UnimplementedHandler struct{}

var _ Handler = UnimplementedHandler{}

// AuthLoginPost implements POST /auth/login operation.
//
// POST /auth/login
func (UnimplementedHandler) AuthLoginPost(ctx context.Context, req *LoginBody) (r *Jwt, _ error) {
	return r, ht.ErrNotImplemented
}

// UserGet implements GET /user operation.
//
// GET /user
func (UnimplementedHandler) UserGet(ctx context.Context) (r *User, _ error) {
	return r, ht.ErrNotImplemented
}

// NewError creates *ErrorStatusCode from error returned by handler.
//
// Used for common default response.
func (UnimplementedHandler) NewError(ctx context.Context, err error) (r *ErrorStatusCode) {
	r = new(ErrorStatusCode)
	return r
}
