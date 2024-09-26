// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package db

import (
	"context"

	"github.com/google/uuid"
)

type Querier interface {
	CreateHolog(ctx context.Context, arg CreateHologParams) (CreateHologRow, error)
	CreateSchedule(ctx context.Context, arg CreateScheduleParams) (CreateScheduleRow, error)
	CreateScheduleDetail(ctx context.Context, arg CreateScheduleDetailParams) (CreateScheduleDetailRow, error)
	CreateUser(ctx context.Context, arg CreateUserParams) (CreateUserRow, error)
	DeleteBookmarkByHologId(ctx context.Context, arg DeleteBookmarkByHologIdParams) (DeleteBookmarkByHologIdRow, error)
	DeleteHologByID(ctx context.Context, id uuid.UUID) (DeleteHologByIDRow, error)
	DeleteScheduleDetail(ctx context.Context, arg DeleteScheduleDetailParams) (DeleteScheduleDetailRow, error)
	GetBookmarkByUserIDAndPlaceID(ctx context.Context, arg GetBookmarkByUserIDAndPlaceIDParams) (GetBookmarkByUserIDAndPlaceIDRow, error)
	GetHologByID(ctx context.Context, id uuid.UUID) (GetHologByIDRow, error)
	GetMyScheduleDetailsByScheduleId(ctx context.Context, scheduleID uuid.UUID) ([]GetMyScheduleDetailsByScheduleIdRow, error)
	GetScheduleByUserID(ctx context.Context, userID string) (GetScheduleByUserIDRow, error)
	GetScheduleDetailByScheduleIdAndPlaceId(ctx context.Context, arg GetScheduleDetailByScheduleIdAndPlaceIdParams) ([]GetScheduleDetailByScheduleIdAndPlaceIdRow, error)
	GetUserByEmailAndPassword(ctx context.Context, arg GetUserByEmailAndPasswordParams) (GetUserByEmailAndPasswordRow, error)
	GetUserByID(ctx context.Context, id string) (GetUserByIDRow, error)
	HardDeleteUserByID(ctx context.Context, id string) (HardDeleteUserByIDRow, error)
	ListAnnounces(ctx context.Context) ([]ListAnnouncesRow, error)
	ListHologsByPlaceId(ctx context.Context, arg ListHologsByPlaceIdParams) ([]ListHologsByPlaceIdRow, error)
	ListHologsByUserID(ctx context.Context, creatorID string) ([]ListHologsByUserIDRow, error)
	ListHologsMostByWeek(ctx context.Context) ([]ListHologsMostByWeekRow, error)
	SetBookmarkByHologId(ctx context.Context, arg SetBookmarkByHologIdParams) (SetBookmarkByHologIdRow, error)
	SoftDeleteUserByID(ctx context.Context, id string) (SoftDeleteUserByIDRow, error)
	UpdateSchedule(ctx context.Context, arg UpdateScheduleParams) (UpdateScheduleRow, error)
}

var _ Querier = (*Queries)(nil)
