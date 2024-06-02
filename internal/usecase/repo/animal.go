package repo

import (
	"Farmish/internal/entity"
	"Farmish/pkg/postgres"
	"context"

	"github.com/pkg/errors"
)

const _defaultEntityCap = 64

type AnimalRepo struct {
	*postgres.Postgres
}

func NewAnimalRepo(pg *postgres.Postgres) *AnimalRepo {
	return &AnimalRepo{pg}
}

func (r *AnimalRepo) CreateAnimal(ctx context.Context, request *entity.Animal) (*entity.Animal, error) {
	var (
		response entity.Animal
	)
	data := map[string]interface{}{
		"name":   request.Name,
		"weight": request.Weight,
		"id":     request.ID,
	}

	sql, _, err := r.Builder.Insert("animals").
		SetMap(data).
		Suffix("RETURNING id, name, weight").
		ToSql()
	if err != nil {
		return nil, err
	}

	row := r.Pool.QueryRow(ctx, sql)

	if err := row.Scan(&response.ID, &response.Name, &response.Weight); err != nil {
		return nil, err
	}

	return &response, nil
}

func (r *AnimalRepo) GetAnimalByID(ctx context.Context, request string) (*entity.Animal, error) {
	//implement here
	return nil, errors.New("unimplemented method - GetAnimal")
}

func (r *AnimalRepo) UpdateAnimal(ctx context.Context, request *entity.Animal) (*entity.Animal, error) {
	//implement here
	return nil, errors.New("unimplemented method - UpdateAnimal")
}

func (r *AnimalRepo) DeleteAnimal(ctx context.Context, id string) (*entity.Animal, error) {
	//implement here
	return nil, errors.New("unimplemented method - DeleteAnimal")
}

func (r *AnimalRepo) GetAllAnimalsByFields(ctx context.Context, request map[string]interface{}) ([]entity.Animal, error) {
	//implement here
	return nil, errors.New("unimplemented method - getall animals by fields")
}
