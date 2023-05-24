package mem

import (
	"context"
	"exchange/pkg/domain"
	"sync"
)

type memoryEmailRepository struct {
	db map[string]struct{}
	mu sync.RWMutex
}

func NewMemoryRepository() domain.EmailRepository {
	return &memoryEmailRepository{
		db: make(map[string]struct{}),
		mu: sync.RWMutex{},
	}
}

func (m *memoryEmailRepository) SaveEmail(ctx context.Context, eu *domain.EmailUser) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.db[eu.Email] = struct{}{}

	return nil
}

func (m *memoryEmailRepository) GetByEmail(
	ctx context.Context,
	email string,
) (*domain.EmailUser, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	_, ok := m.db[email]
	if !ok {
		return nil, domain.ErrNotFound
	}

	return domain.NewEmailUser(email), nil
}

func (m *memoryEmailRepository) GetAllEmails(
	ctx context.Context,
) ([]string, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	emails := make([]string, 0, len(m.db))

	for key := range m.db {
		emails = append(emails, key)
	}

	return emails, nil
}
func (m *memoryEmailRepository) EmailExist(ctx context.Context, email string) (bool, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	_, ok := m.db[email]

	return ok, nil
}
