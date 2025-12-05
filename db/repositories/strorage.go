package db

type Storage struct {
	UsersRepository UsersRepository
}

func NewStorage() *Storage {
	return &Storage{
		UsersRepository: &UserRepositoryImp{},
	}
}