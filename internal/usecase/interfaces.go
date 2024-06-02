// Package usecase implements application business logic. Each logic group in own file.
package usecase

import (
	"context"

	"Farmish/internal/entity"
)

//go:generate mockgen -source=interfaces.go -destination=./mocks_test.go -package=usecase_test

type (
	// Admin -.
	Admin interface {
		Login(context.Context, *entity.LoginRequest) (*entity.LoginResponse, error)
	}

	// AdminRepo -.
	AdminRepo interface {
		CheckField(context.Context, map[string]interface{}) (bool, error)
	}

	//TranslationWebAPI -.
	//TranslationWebAPI interface {
	//	Translate(entity.Translation) (entity.Translation, error)
	//}
)
