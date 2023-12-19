package postgres_db

import (
	"service-for-example/internal/model"
)

type UserApiImpl struct {
	//db БД постгры
}

func NewUserApi() *UserApiImpl {
	return &UserApiImpl{
		//db: БД которое передаешь в функцию
	}
}

func (a *UserApiImpl) FindUsersWithLoginDB(login string) ([]*model.UserItem, error) {
	return nil, nil
}

func (a *UserApiImpl) GetUserByUsernameDB(username string) (*model.UserItem, error) {
	return nil, nil
}

func (a *UserApiImpl) GetUserByIDDB(id int64) (*model.UserItem, error) {
	return nil, nil
}
