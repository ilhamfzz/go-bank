package repository

import (
	"context"
	"database/sql"

	"go-bank/domain"

	"github.com/doug-martin/goqu/v9"
)

type userRepository struct {
	db *goqu.Database
}

func NewUser(con *sql.DB) domain.UserRepository {
	return &userRepository{
		db: goqu.New("default", con),
	}
}

func (u userRepository) Create(ctx context.Context, user *domain.User) (domain.User, error) {
	dataset := u.db.Insert("users").
		Cols("full_name", "phone", "email", "username", "password").
		Vals(goqu.Vals{
			user.FullName,
			user.Phone,
			user.Email,
			user.Username,
			user.Password,
		}).Executor()

	_, err := dataset.ExecContext(ctx)
	if err != nil {
		return domain.User{}, err
	}

	return *user, nil
}

func (u userRepository) Update(ctx context.Context, user *domain.User) error {
	user.EmailVerifiedAtDB = sql.NullTime{
		Time:  user.EmailVerifiedAt,
		Valid: true,
	}

	dataset := u.db.Update("users").Where(goqu.Ex{
		"id": user.ID,
	}).Set(goqu.Record{
		"full_name":         user.FullName,
		"phone":             user.Phone,
		"email":             user.Email,
		"username":          user.Username,
		"password":          user.Password,
		"email_verified_at": user.EmailVerifiedAtDB,
	}).Executor()

	_, err := dataset.ExecContext(ctx)
	return err
}

func (u userRepository) FindByID(ctx context.Context, id int64) (user domain.User, err error) {
	dataset := u.db.From("users").Where(goqu.Ex{
		"id": id,
	})

	_, err = dataset.ScanStructContext(ctx, &user)
	return
}

func (u userRepository) FindByUsername(ctx context.Context, username string) (user domain.User, err error) {
	dataset := u.db.From("users").Where(goqu.Ex{
		"username": username,
	})

	_, err = dataset.ScanStructContext(ctx, &user)
	return
}

func (u userRepository) FindByEmail(ctx context.Context, email string) (user domain.User, err error) {
	dataset := u.db.From("users").Where(goqu.Ex{
		"email": email,
	})

	_, err = dataset.ScanStructContext(ctx, &user)
	return
}

func (u userRepository) GetLastID(ctx context.Context) (int64, error) {
	var result struct {
		ID int64 `db:"id"`
	}

	dataset := u.db.From("users").Order(goqu.C("id").Desc()).Limit(1)

	_, err := dataset.ScanStructContext(ctx, &result)
	if err != nil {
		return 0, err
	}

	return result.ID, nil
}
