package usecase

import (
	"Farmish/config"
	"Farmish/internal/entity"
	"context"
	"errors"

	"github.com/go-redis/redis/v8"
)

type AnimalUseCase struct {
	repo        AnimalRepo
	cfg         *config.Config
	RedisClient *redis.Client
}

func NewAnimalUseCase(r AnimalRepo, cfg *config.Config, RedisClient *redis.Client) *AnimalUseCase {
	return &AnimalUseCase{
		repo:        r,
		cfg:         cfg,
		RedisClient: RedisClient,

		//webAPI: w,
	}
}

func (uc *AnimalUseCase) CreateAnimal(ctx context.Context, request *entity.Animal) (*entity.Animal, error) {
	//iplement here
	return nil, errors.New("unimplemented usecase createanimal")
}

func (uc *AnimalUseCase) GetAnimalByID(ctx context.Context, id string) (*entity.Animal, error) {
	//iplement here
	return nil, errors.New("unimplemented usecase getanimalbyid")
}

func (uc *AnimalUseCase) UpdateAnimal(ctx context.Context, request *entity.Animal) (*entity.Animal, error) {
	//iplement here
	return nil, errors.New("unimplemented usecase updateanimal")
}

func (uc *AnimalUseCase) DeleteAnimal(ctx context.Context, id string) (*entity.Animal, error) {
	//iplement here
	return nil, errors.New("unimplemented usecase deleteanimal")
}

func (uc *AnimalUseCase) GetAllAnimalsByFields(ctx context.Context, filter map[string]interface{}) ([]entity.Animal, error) {
	//iplement here
	return nil, errors.New("unimplemented usecase getallanimalsbyfields")
}
