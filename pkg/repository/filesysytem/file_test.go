package filesysytem

import (
	"context"
	"exchange/pkg/domain"
	"fmt"
	"os"
	"reflect"
	"testing"

	"github.com/bxcodec/faker/v3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const filePath = "test.txt"

func TestFileSave(t *testing.T) {
	ctx := context.Background()

	repo, err := NewFileSystemRepository(filePath)
	require.NoError(t, err)

	defer os.Remove(filePath)

	var testEmail = faker.Email()

	err = repo.SaveEmail(ctx, domain.NewEmailUser(testEmail))
	require.NoError(t, err)

	fileContent, err := os.ReadFile(filePath)
	require.NoError(t, err)

	assert.Equal(t, testEmail, string(fileContent))
}

func TestSave(t *testing.T) {
	ctx := context.Background()

	repo, err := NewFileSystemRepository(filePath)
	require.NoError(t, err)

	defer os.Remove(filePath)

	var testEmail = faker.Email()

	err = repo.SaveEmail(ctx, domain.NewEmailUser(testEmail))
	require.NoError(t, err)

	email, err := repo.GetByEmail(ctx, testEmail)
	require.NoError(t, err)

	assert.Equal(t, testEmail, email.Email)
}

func TestEmailExist(t *testing.T) {
	ctx := context.Background()
	batch := 20

	repo, err := NewFileSystemRepository(filePath)
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(filePath)

	for i := 0; i < batch; i++ {
		mail := faker.Email()
		err = repo.SaveEmail(ctx, domain.NewEmailUser(mail))
		require.NoError(t, err)

		ok, err := repo.EmailExist(ctx, mail)
		require.NoError(t, err)
		require.True(t, ok)
	}

	ok, err := repo.EmailExist(ctx, faker.Email())
	require.NoError(t, err)
	require.False(t, ok)
}

func TestGetAll(t *testing.T) {
	ctx := context.Background()
	batch := 20

	repo, err := NewFileSystemRepository(filePath)
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(filePath)

	emails := make([]string, batch)
	for i := 0; i < batch; i++ {
		mail := faker.Email()
		err = repo.SaveEmail(ctx, domain.NewEmailUser(mail))
		require.NoError(t, err)

		emails[i] = mail
	}

	getEmails, err := repo.GetAllEmails(ctx)
	require.NoError(t, err)
	fmt.Println(emails, len(getEmails))
	require.True(t, reflect.DeepEqual(emails, getEmails), "slices elements are not equal")
}
