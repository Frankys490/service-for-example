package postgres_db

import (
	"service-for-example/internal/model"
)

type ChatApiImpl struct {
	//db БД постгры
}

func NewChatApi() *ChatApiImpl {
	return &ChatApiImpl{
		//db: БД которое передаешь в функцию
	}
}

func (a *ChatApiImpl) GetAllUserChatsDB(id int64) ([]*model.ChatItem, error) {
	return nil, nil
}

func (a *ChatApiImpl) CreateNewChatDB(hostUID, UID int64) (*model.ChatItem, error) {
	return nil, nil
}

func (a *ChatApiImpl) CheckExistingChatDB(hostUID, UID int64) (bool, *model.ChatItem, error) {
	return false, nil, nil
}

func (a *ChatApiImpl) GetChatByHostUIDAndUIDDB(hostUID, UID int64) (*model.ChatItem, error) {
	return nil, nil
}

func (a *ChatApiImpl) AddLastMessageToChatDB(message string, id, time int64) error {
	return nil
}

func (a *ChatApiImpl) UserParticipantOfChatDB(uid int64) (bool, error) {
	return false, nil
}
