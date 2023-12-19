package service

import (
	"chat-service/internal/model"
	"service-for-example/internal/api_db/postgres_db"
)

type ChatService interface {
	FullTextSearchService(req model.FullTextSearchHandlerReq) (*model.FullTextSearchHandlerRes, error)
	GetChatByUsernameService(req model.GetChatByUsernameHandlerReq) (*model.GetChatByUsernameHandlerRes, error)
	GetAllUserChatsService(req model.GetAllUserChatsHandlerReq) ([]*model.GetAllUserChatsHandlerRes, error)
	GetChatService(req model.GetChatHandlerReq) (*model.GetChatHandlerRes, error)
	SendMessageService(req model.SendMessageHandlerReq) (*model.SendMessageHandlerRes, error)
	SendFileService(req model.SendFilesHandlerReq) (*model.SendFilesHandlerRes, error)
	SetOnlineService(req model.SetOnlineHandlerReq) error
	FullTextMessageSearchService(req model.FullTextMessageSearchHandlerReq) (*model.FullTextMessageSearchHandlerRes, error)
}

type Service struct {
	ChatService
}

func NewService(userApiDB *postgres_db.UserApiDB, chatApiDB *postgres_db.ChatApiDB, messageApiDB *postgres_db.MessageApiDB, activeApiDB *postgres_db.ActiveApiDB, generalApiDB *postgres_db.GeneralApiDB) *Service {
	return &Service{
		ChatService: NewChatService(userApiDB.UserApi, chatApiDB.ChatApi, messageApiDB.MessageApi, activeApiDB.ActiveApi, generalApiDB.GeneralApi),
	}
}
