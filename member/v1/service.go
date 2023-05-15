package v1

type ServiceRegister interface {
	GetEventByID(id int) (*Member, error)
}

type serviceRegister struct {
	repository RepositoryRegister
}

func (s serviceRegister) GetEventByID(id int) (*Member, error) {
	member, err := s.repository.GetEventByID(id)
	if err != nil {
		return nil, err
	}

	return member, nil
}

func NewService(repository RepositoryRegister) ServiceRegister {
	return &serviceRegister{repository: repository}
}
