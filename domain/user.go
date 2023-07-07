package domain

import (
	"context"
	"database/sql"
	"time"

	dto "go-bank/dto"
)

type User struct {
	ID                int64        `db:"id"`
	FullName          string       `db:"full_name"`
	Phone             string       `db:"phone"`
	Email             string       `db:"email"`
	Username          string       `db:"username"`
	Password          string       `db:"password"`
	EmailVerifiedAtDB sql.NullTime `db:"email_verified_at"`
	EmailVerifiedAt   time.Time    `db:"-"`
}

type UserRepository interface {
	Create(ctx context.Context, user *User) (User, error)
	Update(ctx context.Context, user *User) error
	FindByID(ctx context.Context, id int64) (User, error)
	FindByUsername(ctx context.Context, username string) (User, error)
	FindByEmail(ctx context.Context, email string) (User, error)
	GetLastID(ctx context.Context) (int64, error)
}

type UserService interface {
	Register(ctx context.Context, req dto.RegisterReq) (dto.RegisterRes, error)
	Authenticate(ctx context.Context, req dto.AuthReq) (dto.AuthRes, error)
	ValidateToken(ctx context.Context, token string) (dto.UserData, error)
	ValidateOTP(ctx context.Context, req dto.ValidateOTPReq) error
}
