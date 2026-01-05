package cache

import (
	"context"
	"errors"
	"github.com/ray-laboratories/saturn/types"
	"sync"
	"time"
)

type TimedSession struct {
	*types.Session
	Entered time.Time
}

type SessionRepository struct {
	cache map[string]TimedSession
	mutex sync.RWMutex
	ttl   time.Duration
}

func (s *SessionRepository) Save(ctx context.Context, session *types.Session) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.cache[session.Token] = TimedSession{session, time.Now()}
	return nil
}

func (s *SessionRepository) Get(ctx context.Context, token string) (*types.Session, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	session, ok := s.cache[token]
	if !ok {
		return nil, errors.New("session not found")
	}
	if time.Since(session.Entered) > s.ttl {
		return nil, errors.New("session expired")
	}
	return session.Session, nil
}

func (s *SessionRepository) Delete(ctx context.Context, token string) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	delete(s.cache, token)
	return nil
}

func NewSessionRepository(ttl time.Duration) *SessionRepository {
	return &SessionRepository{
		cache: make(map[string]TimedSession),
		ttl:   ttl,
	}
}
