package storage

import (
	"sync"

	"fitness/internal/model"
)

type UserStorage interface {
	GetAll() []model.User
	Get(email string) (model.User, bool)
	Create(user model.User)
	Update(email string, user model.User)
	Delete(email string)
	Exists(email string) bool
}

type MemoryStorage struct {
	users map[string]model.User
	mutex sync.RWMutex
}

func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{
		users: make(map[string]model.User),
	}
}

func (s *MemoryStorage) GetAll() []model.User {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	users := make([]model.User, 0, len(s.users))
	for _, user := range s.users {
		users = append(users, user)
	}
	return users
}

func (s *MemoryStorage) Get(email string) (model.User, bool) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	user, exists := s.users[email]
	return user, exists
}

func (s *MemoryStorage) Create(user model.User) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.users[user.Email] = user
}

func (s *MemoryStorage) Update(email string, user model.User) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.users[email] = user
}

func (s *MemoryStorage) Delete(email string) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	delete(s.users, email)
}

func (s *MemoryStorage) Exists(email string) bool {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	_, exists := s.users[email]
	return exists
}