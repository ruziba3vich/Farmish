package usecase

import (
	"Farmish/config"
	"Farmish/internal/entity"
	"context"

	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
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
	request.ID = uuid.NewString()
	return uc.repo.CreateAnimal(ctx, request)
}

func (uc *AnimalUseCase) GetAnimalByID(ctx context.Context, id string) (*entity.Animal, error) {
	return uc.repo.GetAnimalByID(ctx, id)
}

func (uc *AnimalUseCase) UpdateAnimal(ctx context.Context, request *entity.Animal) (*entity.Animal, error) {
	return uc.repo.UpdateAnimal(ctx, request)
}

func (uc *AnimalUseCase) DeleteAnimal(ctx context.Context, id string) (*entity.Animal, error) {
	return uc.repo.DeleteAnimal(ctx, id)
}

func (uc *AnimalUseCase) GetAllAnimalsByFields(ctx context.Context, filter map[string]interface{}) ([]entity.Animal, error) {
	return uc.repo.GetAllAnimalsByFields(ctx, filter)
}

func (uc *AnimalUseCase) NotifyAnimalStatus(ctx context.Context) ([]entity.AnimalHungryReponse, error) {
	return uc.repo.CheckHungryStatusOfAnimals(ctx)
}
