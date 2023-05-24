package filesysytem

import (
	"context"
	"exchange/pkg/domain"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"sync"
)

type fileSystemRepository struct {
	filePath string
	index    map[string]struct{}
	fm       sync.RWMutex
	im       sync.RWMutex
}

func NewFileSystemRepository(filePath string) (domain.EmailRepository, error) {
	f := &fileSystemRepository{
		filePath: filePath,
		fm:       sync.RWMutex{},
		im:       sync.RWMutex{},
	}

	if err := f.loadIndex(); err != nil {
		return nil, err
	}

	return f, nil

}
func (f *fileSystemRepository) SaveEmail(ctx context.Context, eu *domain.EmailUser) error {
	f.fm.Lock()
	defer f.fm.Unlock()

	file, err := os.OpenFile(f.filePath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		return fmt.Errorf("failed to open data file: %w", err)
	}
	defer file.Close()

	if _, err = file.WriteString(eu.Email + "\n"); err != nil {
		return fmt.Errorf("failed write to file")
	}

	f.im.Lock()
	defer f.im.Unlock()

	f.index[eu.Email] = struct{}{}

	return nil
}

func (f *fileSystemRepository) GetByEmail(
	ctx context.Context,
	email string,
) (*domain.EmailUser, error) {
	f.im.RLock()
	defer f.im.RUnlock()

	_, ok := f.index[email]
	if !ok {
		return nil, domain.ErrNotFound
	}

	return domain.NewEmailUser(email), nil
}

func (f *fileSystemRepository) GetAllEmails(
	ctx context.Context,
) ([]string, error) {
	f.fm.RLock()
	defer f.fm.RUnlock()

	if _, err := os.Stat(f.filePath); os.IsNotExist(err) {
		return []string{}, nil
	}

	data, err := ioutil.ReadFile(f.filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read file by path: %s", f.filePath)
	}

	rows := strings.Split(string(data), "\n")

	emails := make([]string, 0, len(rows))

	for _, row := range rows {
		if row == "" {
			continue
		}

		emails = append(emails, row)
	}

	return emails, nil
}

func (f *fileSystemRepository) EmailExist(ctx context.Context, email string) (bool, error) {
	f.im.RLock()
	defer f.im.RUnlock()

	_, ok := f.index[email]

	return ok, nil
}
