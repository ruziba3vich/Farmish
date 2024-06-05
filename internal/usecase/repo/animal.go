package repo

import (
	"Farmish/internal/entity"
	"Farmish/pkg/postgres"
	"context"
	"github.com/k0kubun/pp"

	"github.com/Masterminds/squirrel"
)

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

	sql, args, err := r.Builder.Insert("animals").
		SetMap(data).
		Suffix("RETURNING id, name, weight").
		ToSql()
	if err != nil {
		return nil, err
	}

	row := r.Pool.QueryRow(ctx, sql, args...)

	if err := row.Scan(&response.ID, &response.Name, &response.Weight); err != nil {
		return nil, err
	}

	return &response, nil
}

func (r *AnimalRepo) GetAnimalByID(ctx context.Context, request string) (*entity.Animal, error) {
	var (
		response entity.Animal
	)

	sql, args, err := r.Builder.Select("id", "name", "weight", "is_hungry").
		From("animals").
		Where(squirrel.Eq{"id": request}).
		ToSql()
	if err != nil {
		return nil, err
	}

	row := r.Pool.QueryRow(ctx, sql, args...)

	if err := row.Scan(&response.ID, &response.Name, &response.Weight); err != nil {
		return nil, err
	}

	return &response, nil
}

func (r *AnimalRepo) UpdateAnimal(ctx context.Context, request *entity.Animal) (*entity.Animal, error) {
	var (
		response entity.Animal
	)

	data := map[string]interface{}{
		"id":        request.ID,
		"name":      request.Name,
		"weight":    request.Weight,
		"is_hungry": request.IsHungry,
	}

	sql, args, err := r.Builder.Update("animals").
		SetMap(data).
		Suffix("RETURNING id, name, weight, is_hungry").
		Where(squirrel.Eq{"id": request.ID}).
		ToSql()

	if err != nil {
		return nil, err
	}

	row := r.Pool.QueryRow(ctx, sql, args...)

	if err := row.Scan(&response.ID, &response.Name, &response.Weight, &response.IsHungry); err != nil {
		return nil, err
	}

	return &response, nil
}

func (r *AnimalRepo) DeleteAnimal(ctx context.Context, id string) (*entity.Animal, error) {
	var (
		response entity.Animal
	)

	sql, _, err := r.Builder.Delete("animals").
		Prefix(" RETURNING id, name, weight, is_hungry").
		Where(squirrel.Eq{"id": id}).
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

func (r *AnimalRepo) GetAllAnimalsByFields(ctx context.Context, request map[string]interface{}) ([]entity.Animal, error) {
	var (
		response []entity.Animal
		builder  squirrel.SelectBuilder
	)

	if len(request) == 0 {
		builder = r.Builder.Select("animals")
	} else {
		builder := r.Builder.Select("animals").
			Where(squirrel.And{})

		for key, value := range request {
			builder = builder.Where(squirrel.Eq{
				key: value,
			})
		}
	}

	sql, _, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := r.Pool.Query(ctx, sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var animal entity.Animal
		if err := rows.Scan(&animal.ID, &animal.Name, &animal.Weight, &animal.IsHungry); err != nil {
			return nil, err
		}

		response = append(response, animal)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return response, nil
}

func (r *AnimalRepo) CheckHungryStatusOfAnimals(ctx context.Context) ([]entity.AnimalHungryReponse, error) {
	var (
		animal   entity.AnimalHungryReponse
		animals  []entity.AnimalHungryReponse
		IsHungry bool
	)
	sql, args, err := r.Builder.Select("id, name, is_hungry").
		From("animals").
		ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := r.Pool.Query(ctx, sql, args...)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		if err := rows.Scan(&animal.ID, &animal.Name, &IsHungry); err != nil {
			return nil, err
		}
		pp.Println(animal)
		if IsHungry {
			animals = append(animals, animal)
		}
	}

	return animals, nil
}
