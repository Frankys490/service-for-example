package postgres_db

type ActiveApiImpl struct {
	//db БД постгры
}

func NewActiveApi() *ActiveApiImpl {
	return &ActiveApiImpl{
		//db: БД которое передаешь в функцию
	}
}

func (a *ActiveApiImpl) GetOnlineDB(uid int64) (bool, error) {
	return false, nil
}

func (a *ActiveApiImpl) SetOnlineDB(uid int64) error {
	return nil
}
