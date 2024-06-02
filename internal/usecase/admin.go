package usecase

import (
	"context"
	"fmt"
	"github.com/evrone/go-clean-template/config"
	tokens "github.com/evrone/go-clean-template/pkg/token"
	"github.com/go-redis/redis/v8"
	"github.com/k0kubun/pp"
	"time"

	"github.com/evrone/go-clean-template/internal/entity"
)

// AdminUseCase -.
type AdminUseCase struct {
	repo        AdminRepo
	cfg         *config.Config
	RedisClient *redis.Client
	//webAPI TranslationWebAPI
}

// NewAdminUseCase -.
func NewAdminUseCase(r AdminRepo, cfg *config.Config, RedisClient *redis.Client) *AdminUseCase {
	return &AdminUseCase{
		repo:        r,
		cfg:         cfg,
		RedisClient: RedisClient,

		//webAPI: w,
	}
}

// Login - login for admins.
func (uc *AdminUseCase) Login(ctx context.Context, request *entity.LoginRequest) (*entity.LoginResponse, error) {
	data := map[string]interface{}{
		"email": request.Email,
	}
	exists, err := uc.repo.CheckField(ctx, data)
	if err != nil {
		return nil, fmt.Errorf("AdminUseCase - Login - s.repo.GetAdmin: %w", err)
	}

	if !exists {
		return nil, fmt.Errorf("admin with given email does not exists")
	}

	byteDataForRedis := []byte(request.Password)

	status := uc.RedisClient.Set(ctx, request.Email, byteDataForRedis, time.Minute*5)
	if status.Err() != nil {
		return nil, fmt.Errorf("AdminUseCase - Login - s.repo.GetAdmin: %w", status.Err())
	}

	jwtHandler := tokens.JWTHandler{
		Sub:       "a",
		Iss:       time.Now().String(),
		Exp:       time.Now().Add(time.Hour * 6).String(),
		Role:      "owner",
		SigninKey: uc.cfg.Casbin.SigningKey,
		Timeout:   uc.cfg.Casbin.AccessTokenTimeOut,
	}

	access, refresh, err := jwtHandler.GenerateAuthJWT()
	if err != nil {
		return nil, fmt.Errorf("AdminUseCase - Login - s.repo.GetAdmin: %w", err)
	}

	pp.Println(access)
	pp.Println(refresh)

	return &entity.LoginResponse{AccessToken: access}, nil

}

// Translate -.
//func (uc *TranslationUseCase) Translate(ctx context.Context, t entity.Translation) (entity.Translation, error) {
//	translation, err := uc.webAPI.Translate(t)
//	if err != nil {
//		return entity.Translation{}, fmt.Errorf("TranslationUseCase - Translate - s.webAPI.Translate: %w", err)
//	}
//
//	err = uc.repo.Store(context.Background(), translation)
//	if err != nil {
//		return entity.Translation{}, fmt.Errorf("TranslationUseCase - Translate - s.repo.Store: %w", err)
//	}
//
//	return translation, nil
//}
