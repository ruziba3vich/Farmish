// Package usecase implements application business logic. Each logic group in own file.
package usecase

import (
	"context"

	"Farmish/internal/entity"
)

//go:generate mockgen -source=interfaces.go -destination=./mocks_test.go -package=usecase_test

type (
	Animal interface {
		CreateAnimal(context.Context, *entity.Animal) (*entity.Animal, error)
		GetAnimalByID(context.Context, string) (*entity.Animal, error)
		UpdateAnimal(context.Context, *entity.Animal) (*entity.Animal, error)
		DeleteAnimal(context.Context, string) (*entity.Animal, error)
		GetAllAnimalsByFields(context.Context, map[string]interface{}) ([]entity.Animal, error)
	}

	AnimalRepo interface {
		CreateAnimal(context.Context, *entity.Animal) (*entity.Animal, error)
		GetAnimalByID(context.Context, string) (*entity.Animal, error)
		UpdateAnimal(context.Context, *entity.Animal) (*entity.Animal, error)
		DeleteAnimal(context.Context, string) (*entity.Animal, error)
		GetAllAnimalsByFields(context.Context, map[string]interface{}) ([]entity.Animal, error)
	}
)
