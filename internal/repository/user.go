package repository

import (
	"context"
	"database/sql"

	sq "github.com/Masterminds/squirrel"
	"github.com/dickykmrlh/user/internal/core/domain"
	"github.com/google/uuid"
)

var userColumns = []string{"id", "first_name", "last_name", "role", "phone_number", "email"}

const usersTableName = "users"

type UserRepo struct {
	db *sql.DB
}

func New(db *sql.DB) *UserRepo {
	return &UserRepo{db: db}
}

func (u *UserRepo) GetUser(ctx context.Context, id uuid.UUID) (user *domain.User, err error) {
	user = &domain.User{}

	builder := sq.StatementBuilder.PlaceholderFormat(sq.Dollar).RunWith(u.db)
	err = builder.Select(userColumns...).
		From(usersTableName).
		Where(sq.Eq{"id": id}).
		QueryRowContext(ctx).
		Scan(scanColumns(user)...)
	if err == sql.ErrNoRows {
		return nil, nil
	}

	return
}

func (u *UserRepo) Save(user *domain.User) (err error) {
	return nil
}

func scanColumns(record *domain.User) []any {
	return []any{
		&record.ID,
		&record.FirstName,
		&record.LastName,
		&record.Role,
		&record.PhoneNumber,
		&record.Email,
	}
}
