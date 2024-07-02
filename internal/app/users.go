package app

type usersServiceStorage interface {
}

type UsersService struct {
	storage usersServiceStorage
}

func NewUsersService(storage usersServiceStorage) *UsersService {
	return &UsersService{storage: storage}
}
