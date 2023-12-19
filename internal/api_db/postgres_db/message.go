package postgres_db

import (
	"service-for-example/internal/model"
)

type MessageApiImpl struct {
	//db БД постгры
}

func NewMessageApi() *MessageApiImpl {
	return &MessageApiImpl{
		//db: БД которое передаешь в функцию
	}
}

func (a *MessageApiImpl) GetAllChatMessagesDB(chatId int64) ([]*model.ChatMessageItem, error) {
	return nil, nil
}

func (a *MessageApiImpl) CreateNewMessageDB(chatID, uid int64, messageStr string, typeMessage string) (*model.ChatMessageItem, error) {
	return nil, nil
}

func (a *MessageApiImpl) GetLastMessageDB(chatID int64) (*model.ChatMessageItem, error) {
	return nil, nil
}

func (a *MessageApiImpl) FullTextMessageSearchDB(search string, chatID int64) ([]*model.ChatMessageItem, error) {
	return nil, nil
}
