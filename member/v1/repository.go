package v1

import (
	"gorm.io/gorm"
)

type RepositoryRegister interface {
	GetEventByID(id int) (*Member, error)
}

type repositoryRegister struct {
	database *gorm.DB
}

// GetEventByID implements RepositoryRegister
func (r repositoryRegister) GetEventByID(id int) (*Member, error) {
	var event *Member
	tx := r.database.First(&event, id)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return event, nil
}

func NewRepository(database *gorm.DB) RepositoryRegister {
	return &repositoryRegister{database: database}
}
