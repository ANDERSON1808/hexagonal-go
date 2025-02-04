package usecases

import (
	"ANDERSON1808/hexagonal-go/internal/domain"
	"log"
	"sync"
	"time"
)

type UserService struct {
	repo domain.UserRepository
}

func NewUserService(repo domain.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) CreateUser(user *domain.User) error {
	return s.repo.Save(user)
}

func (s *UserService) GetUser(id uint) (*domain.User, error) {
	return s.repo.FindByID(id)
}

func (s *UserService) GetAllUsers() ([]domain.User, error) {
	return s.repo.FindAll()
}

func (s *UserService) DeleteUser(id uint) error {
	return s.repo.Delete(id)
}

func (s *UserService) CreateUsersConcurrently(users []*domain.User) {
	var wg sync.WaitGroup
	var mu sync.Mutex

	startTime := time.Now()

	for _, user := range users {
		wg.Add(1)
		go func(u domain.User) {
			defer wg.Done()

			time.Sleep(time.Millisecond * 500)

			mu.Lock()
			err := s.repo.Save(&u)
			mu.Unlock()

			if err != nil {
				log.Printf("[ERROR] No se pudo crear el usuario %s: %v", u.Name, err)
				return
			}
			log.Printf("[INFO] Usuario %s creado exitosamente", u.Name)
		}(*user)
	}

	wg.Wait()

	log.Printf("[INFO] Todos los usuarios creados en %v", time.Since(startTime))
}
