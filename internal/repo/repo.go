package repo

import "github.com/gradelyng/testRepo/repo/internal/models"

//Repo for database operations
type Repo interface {
	CreateUsers(users []models.User, userids *[]uint) ([]models.User, error)
	UpdateUsers(users []models.User, userids *[]uint) ([]models.User, error)
	DeleteUsers(userids *[]uint) error
	GetAllUsers() ([]models.User, error)
}
