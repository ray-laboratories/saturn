package local

import (
	"context"
	"database/sql"
	"errors"
	"github.com/ray-laboratories/saturn/application"
	"github.com/ray-laboratories/saturn/infra/cache"
	"github.com/ray-laboratories/saturn/infra/random"
	"github.com/ray-laboratories/saturn/infra/sqlite"
	"github.com/ray-laboratories/saturn/types"
	"log"
	_ "modernc.org/sqlite"
	"time"
)

var ErrorPasswordIncorrect = errors.New("password incorrect")
var ErrorUsernameTaken = errors.New("username taken")

// Compile-time check
var _ application.AuthService = &AuthService{}

type AuthService struct {
	accessorRepository AccessorRepository
	hasher             Hasher
	tokenizer          Tokenizer
	sessionRepository  SessionRepository
}

func NewAuthService(userRepository AccessorRepository, hasher Hasher, tokenizer Tokenizer, sessionRepository SessionRepository) *AuthService {
	return &AuthService{
		accessorRepository: userRepository,
		hasher:             hasher,
		tokenizer:          tokenizer,
		sessionRepository:  sessionRepository,
	}
}

var defAuthService *AuthService

func NewDefaultAuthService(path string) *AuthService {
	if defAuthService != nil {
		return defAuthService
	}
	db, err := sql.Open("sqlite", path)
	if err != nil {
		log.Fatal(err)
	}
	acccessorRepo := sqlite.NewAccessorRepository(db)
	sessionRepo := cache.NewSessionRepository(time.Hour)
	defAuthService = NewAuthService(acccessorRepo, random.Hasher{}, random.Tokenizer{}, sessionRepo)
	return defAuthService
}

func (a *AuthService) Validate(ctx context.Context, token string) (*types.Accessor, error) {
	session, err := a.sessionRepository.Get(ctx, token)
	if err != nil {
		return nil, err
	}
	return a.accessorRepository.Get(ctx, session.Username)
}

func (a *AuthService) Login(ctx context.Context, username, password string) (string, error) {
	// Find user
	user, err := a.accessorRepository.Get(ctx, username)
	if err != nil {
		return "", err
	}

	// Compare passwords
	if a.hasher.Compare(password, user.HashedPassword) {
		// Generate session ID
		token := a.tokenizer.New()
		session := types.NewSession(token, username)

		// Save session
		err = a.sessionRepository.Save(ctx, session)
		if err != nil {
			return "", err
		}

		// Return token
		return token, nil
	}

	// Password didn't match
	return "", ErrorPasswordIncorrect
}

func (a *AuthService) Logout(ctx context.Context, token string) error {
	// Remove session
	err := a.sessionRepository.Delete(ctx, token)
	if err != nil {
		return err
	}
	return nil
}

func (a *AuthService) Register(ctx context.Context, username, password string) (string, error) {
	hashedPassword, err := a.hasher.Hash(password)
	if err != nil {
		return "", err
	}
	user, err := a.accessorRepository.Get(ctx, username)
	if user != nil {
		return "", ErrorUsernameTaken
	}
	newUser := types.NewAccessor(username, hashedPassword)
	err = a.accessorRepository.Create(ctx, newUser)
	if err != nil {
		return "", err
	}
	token, err := a.Login(ctx, username, password)
	if err != nil {
		return "", err
	}
	return token, nil
}
