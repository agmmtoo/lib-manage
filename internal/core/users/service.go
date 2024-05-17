package users

import "github.com/agmmtoo/lib-manage/pkg/libraryapp"

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

type Repository interface {
	// GetAllUsers returns all users.
	GetAllUsers() ([]*libraryapp.User, error)

	// GetByID returns the user with the specified ID.
	// GetByID(id int) (*libraryapp.User, error)

	// GetByUsername returns the user with the specified username.
	// GetByUsername(username string) (*libraryapp.User, error)

	// Create creates a new user.
	// Create(user *libraryapp.User) (int, error)

	// Update updates an existing user.
	// Update(user *libraryapp.User) error

	// Delete deletes the user with the specified ID.
	// Delete(id int) error
}

// GetAll returns all users.
func (s *Service) GetAllUsers() ([]*libraryapp.User, error) {
	return s.repo.GetAllUsers()
}
