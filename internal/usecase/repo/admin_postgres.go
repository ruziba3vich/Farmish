package repo

import (
	"Farmish/pkg/postgres"
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"
)

const _defaultEntityCap = 64

// AdminRepo -.
type AdminRepo struct {
	*postgres.Postgres
}

// NewAdminRepo -.
func NewAdminRepo(pg *postgres.Postgres) *AdminRepo {
	return &AdminRepo{pg}
}

// CheckField -.
func (r *AdminRepo) CheckField(ctx context.Context, filter map[string]interface{}) (bool, error) {
	var (
		email string
	)
	queryBuilder := r.Builder.
		Select("email").
		From("admins")

	for field, value := range filter {
		queryBuilder = queryBuilder.Where(squirrel.Eq{field: value})
	}

	sql, args, err := queryBuilder.ToSql()
	if err != nil {
		return false, fmt.Errorf("AdminRepo - CheckField - r.Builder: %w", err)
	}

	row := r.Pool.QueryRow(ctx, sql, args...)

	err = row.Scan(&email)
	if err != nil {
		return false, fmt.Errorf("email %s not found in database", email)
	}

	return true, nil
}
