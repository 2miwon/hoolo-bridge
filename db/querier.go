// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package db

import (
	"context"
)

type Querier interface {
	CreateUser(ctx context.Context, arg CreateUserParams) (User, error)
	GetUserByEmailAndPassword(ctx context.Context, arg GetUserByEmailAndPasswordParams) (User, error)
	GetUserByID(ctx context.Context, id string) (User, error)
}

var _ Querier = (*Queries)(nil)
