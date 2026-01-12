package local

import (
	"context"
	"errors"
	"github.com/ray-laboratories/saturn/application/authentication"
	"github.com/ray-laboratories/saturn/types"
	_ "modernc.org/sqlite"
)

var ErrorPasswordIncorrect = errors.New("password incorrect")
var ErrorUsernameTaken = errors.New("username taken")

// Compile-time check
var _ authentication.Service = &AuthenticationService{}

type AuthenticationService struct {
	userRepository    UserRepository
	hasher            Hasher
	tokenizer         Tokenizer
	sessionRepository SessionRepository
}

func NewAuthenticationService(userRepository UserRepository, hasher Hasher, tokenizer Tokenizer, sessionRepository SessionRepository) *AuthenticationService {
	return &AuthenticationService{
		userRepository:    userRepository,
		hasher:            hasher,
		tokenizer:         tokenizer,
		sessionRepository: sessionRepository,
	}
}

func (a *AuthenticationService) Validate(ctx context.Context, token string) (*types.User, error) {
	session, err := a.sessionRepository.Get(ctx, token)
	if err != nil {
		return nil, err
	}
	return session.User, nil
}

func (a *AuthenticationService) Login(ctx context.Context, username, password string) (string, error) {
	// Find user
	user, err := a.userRepository.GetByUsername(ctx, username)
	if err != nil {
		return "", err
	}

	// Compare passwords
	if a.hasher.Compare(password, user.HashedPassword) {
		// Generate session ID
		token := a.tokenizer.New()
		session := types.NewSession(token, user)

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

func (a *AuthenticationService) Logout(ctx context.Context, token string) error {
	// Remove session
	err := a.sessionRepository.Delete(ctx, token)
	if err != nil {
		return err
	}
	return nil
}

func (a *AuthenticationService) Register(ctx context.Context, req *types.UserRequest) (string, error) {
	hashedPassword, err := a.hasher.Hash(req.Password)
	if err != nil {
		return "", err
	}
	user, err := a.userRepository.GetByUsername(ctx, req.Username)
	if user != nil {
		return "", ErrorUsernameTaken
	}
	newUser := types.NewUser(req.Username, hashedPassword, req.GroupID)
	err = a.userRepository.Create(ctx, newUser)
	if err != nil {
		return "", err
	}
	token, err := a.Login(ctx, req.Username, req.Password)
	if err != nil {
		return "", err
	}
	return token, nil
}
