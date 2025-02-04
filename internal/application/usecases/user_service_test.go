package usecases_test

import (
	"ANDERSON1808/hexagonal-go/internal/application/usecases"
	"ANDERSON1808/hexagonal-go/internal/application/usecases/mocks"
	"ANDERSON1808/hexagonal-go/internal/domain"
	"errors"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
	_ "github.com/stretchr/testify/mock"
)

func TestUserService(t *testing.T) {
	mockRepo := new(mocks.UserRepository)
	userService := usecases.NewUserService(mockRepo)

	testUser := &domain.User{
		ID:    1,
		Name:  "Test User",
		Email: "test@example.com",
	}

	concurrentUsers := []*domain.User{
		{ID: 2, Name: "User A", Email: "a@example.com"},
		{ID: 3, Name: "User B", Email: "b@example.com"},
		{ID: 4, Name: "User C", Email: "c@example.com"},
	}

	t.Run("CreateUser - Success", func(t *testing.T) {
		mockRepo.On("Save", testUser).Return(nil).Once()

		err := userService.CreateUser(testUser)
		assert.NoError(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("CreateUser - Failure", func(t *testing.T) {
		mockRepo.On("Save", testUser).Return(errors.New("database error")).Once()

		err := userService.CreateUser(testUser)
		assert.Error(t, err)
		assert.Equal(t, "database error", err.Error())
		mockRepo.AssertExpectations(t)
	})

	t.Run("GetUser - Success", func(t *testing.T) {
		mockRepo.On("FindByID", uint(1)).Return(testUser, nil).Once()

		user, err := userService.GetUser(1)
		assert.NoError(t, err)
		assert.Equal(t, testUser, user)
		mockRepo.AssertExpectations(t)
	})

	t.Run("GetUser - Not Found", func(t *testing.T) {
		mockRepo.On("FindByID", uint(1)).Return(nil, errors.New("user not found")).Once()

		user, err := userService.GetUser(1)
		assert.Error(t, err)
		assert.Nil(t, user)
		assert.Equal(t, "user not found", err.Error())
		mockRepo.AssertExpectations(t)
	})

	t.Run("GetAllUsers - Success", func(t *testing.T) {
		mockRepo.On("FindAll").Return([]domain.User{*testUser}, nil).Once()

		users, err := userService.GetAllUsers()
		assert.NoError(t, err)
		assert.Len(t, users, 1)
		mockRepo.AssertExpectations(t)
	})

	t.Run("DeleteUser - Success", func(t *testing.T) {
		mockRepo.On("Delete", uint(1)).Return(nil).Once()

		err := userService.DeleteUser(1)
		assert.NoError(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("DeleteUser - Failure", func(t *testing.T) {
		mockRepo.On("Delete", uint(1)).Return(errors.New("delete error")).Once()

		err := userService.DeleteUser(1)
		assert.Error(t, err)
		assert.Equal(t, "delete error", err.Error())
		mockRepo.AssertExpectations(t)
	})

	t.Run("CreateUsersConcurrently - Success", func(t *testing.T) {
		var wg sync.WaitGroup

		for _, u := range concurrentUsers {
			mockRepo.On("Save", u).Return(nil).Once()
		}

		wg.Add(1)
		go func() {
			defer wg.Done()
			userService.CreateUsersConcurrently(concurrentUsers)
		}()
		wg.Wait()

		mockRepo.AssertExpectations(t)
	})
}
