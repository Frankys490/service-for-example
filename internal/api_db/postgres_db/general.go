package postgres_db

import (
	"service-for-example/internal/model"
)

type GeneralApiImpl struct {
	//db БД постгры
}

func NewGeneralApi() *GeneralApiImpl {
	return &GeneralApiImpl{
		//db: БД которое передаешь в функцию
	}
}

func (a *GeneralApiImpl) GetRecordByUIDDB(id int64) (*model.GeneralItem, error) {
	return nil, nil
}
