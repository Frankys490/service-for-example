package service

import (
	"fmt"
	"service-for-example/internal/api_db/postgres_db"
	"service-for-example/internal/model"
	"service-for-example/pkg/custom_errors"
	"service-for-example/pkg/files"
	"strconv"
)

type ChatServiceImpl struct {
	userApiDB    postgres_db.UserApi
	chatApiDB    postgres_db.ChatApi
	messageApiDB postgres_db.MessageApi
	activeApiDB  postgres_db.ActiveApi
	generalApiDB postgres_db.GeneralApi
}

func NewChatService(userApiDB postgres_db.UserApi, chatApiDB postgres_db.ChatApi, messageApiDB postgres_db.MessageApi, activeApiDB postgres_db.ActiveApi, generalApiDB postgres_db.GeneralApi) *ChatServiceImpl {
	return &ChatServiceImpl{
		userApiDB:    userApiDB,
		chatApiDB:    chatApiDB,
		messageApiDB: messageApiDB,
		activeApiDB:  activeApiDB,
		generalApiDB: generalApiDB,
	}
}

func (s *ChatServiceImpl) FullTextSearchService(req model.FullTextSearchHandlerReq) (*model.FullTextSearchHandlerRes, error) {
	/*запрос на поис юзера по логину
	users, err := s.userApiDB.FindUsersWithLoginDB(req.Search)
	if err != nil {
		return nil, fmt.Errorf("find users with login DB: %w", err)
	}*/

	var res model.FullTextSearchHandlerRes

	for _, user := range users {
		/*запрос на поиск записи по id
		record, err := s.generalApiDB.GetRecordByUIDDB(user.ID)
		if err != nil {
			return nil, fmt.Errorf("get record by UID DB: %w", err)
		}*/

		/*запрос на получение онлайна по id
		online, err := s.activeApiDB.GetOnlineDB(user.ID)
		if err != nil {
			return nil, fmt.Errorf("get online DB : %w", err)
		}*/

		/*запрос на наличие чата
		found, chat, err := s.chatApiDB.CheckExistingChatDB(req.UID, user.ID)
		*/
		if !found {
			res.Users = append(res.Users, model.UserSearchItem{})
		} else {
			/*получение последнего сообщения
			message, err := s.messageApiDB.GetLastMessageDB(chat.ID)
			if err != nil {
				return nil, fmt.Errorf("get last message DB: %w", err)
			}
			*/

			res.Users = append(res.Users, model.UserSearchItem{})
		}
	}

	return &res, nil
}

func (s *ChatServiceImpl) GetChatByUsernameService(req model.GetChatByUsernameHandlerReq) (*model.GetChatByUsernameHandlerRes, error) {
	/*запрос на получение юзера по id
	user, err := s.userApiDB.GetUserByUsernameDB(req.Username)
	if err != nil {
		return nil, fmt.Errorf("get user by username DB: %w", err)
	}
	*/

	/*запрос на получение чата по id
	chat, err := s.chatApiDB.GetChatByHostUIDAndUIDDB(req.UID, user.ID)
	if err != nil {
		return nil, fmt.Errorf("get chat by host UID and UID DB: %w", err)
	}
	*/

	if chat == nil {
		return nil, nil
	}

	/*запрос на получение всех сообщений чата по id
		messages, err := s.messageApiDB.GetAllChatMessagesDB(chat.ID)
	if err != nil {
		return nil, fmt.Errorf("get all chat messages DB: %w", err)
	}

	*/

	return &model.GetChatByUsernameHandlerRes{Messages: nil, RequestedUserID: req.UID}, nil
}

func (s *ChatServiceImpl) GetAllUserChatsService(req model.GetAllUserChatsHandlerReq) ([]*model.GetAllUserChatsHandlerRes, error) {
	/*запрос на получение всех чатов по id
		chats, err := s.chatApiDB.GetAllUserChatsDB(req.ID)
	if err != nil {
		return nil, fmt.Errorf("get all user chats DB: %w", err)
	}

	*/

	var res []*model.GetAllUserChatsHandlerRes

	var user *model.UserItem

	for _, elem := range chats {
		var id int64
		if req.ID == elem.UID {
			id = elem.HostUID
		}
		id = elem.UID

		if req.ID == elem.UID {
			/*запрос на получение пользователя чата по id
						user, err = s.userApiDB.GetUserByIDDB(id)
			if err != nil {
				return nil, fmt.Errorf("get user by id DB: %w", err)
			}

			*/
		}
		/*запрос на получение пользователя чата по id
		user, err = s.userApiDB.GetUserByIDDB(id)
		if err != nil {
			return nil, fmt.Errorf("get user by id DB: %w", err)
		}

		*/
		/*запрос на получение general по id
		record, err := s.generalApiDB.GetRecordByUIDDB(id)
		if err != nil {
			return nil, fmt.Errorf("get record by uid DB: %w", err)
		}
		*/

		/*запрос на получение онлайна по id
		online, err := s.activeApiDB.GetOnlineDB(id)
		if err != nil {
			return nil, fmt.Errorf("get online DB : %w", err)
		}

		*/

		res = append(res, &model.GetAllUserChatsHandlerRes{Username: user.Username, Chat: elem, Online: online, ProfilePictureLink: record.IMGLink})
	}

	return res, nil
}

func (s *ChatServiceImpl) GetChatService(req model.GetChatHandlerReq) (*model.GetChatHandlerRes, error) {
	/*запрос на проверку пользователя чата по id
		found, err := s.chatApiDB.UserParticipantOfChatDB(req.UID)
	if err != nil {
		return nil, fmt.Errorf("user participant of chat DB: %w", err)
	}
	*/

	if !found {
		return nil, fmt.Errorf("user isn't participant of this chat: %w", custom_errors.ErrNotParticipant)
	}

	/*запрос на получение сообщений чата по id
	messages, err := s.messageApiDB.GetAllChatMessagesDB(req.ChatID)
	if err != nil {
		return nil, fmt.Errorf("get all chat messages: %w", err)
	}

	*/

	return &model.GetChatHandlerRes{Messages: messages, RequestedUserID: req.UID}, nil
}

func (s *ChatServiceImpl) SendMessageService(req model.SendMessageHandlerReq) (*model.SendMessageHandlerRes, error) {
	/*запрос на получение пользователя чата по юзернейму
		receiverUser, err := s.userApiDB.GetUserByUsernameDB(req.ReceiverUsername)
	if err != nil {
		return nil, fmt.Errorf("get user by username DB: %w", err)
	}
	*/

	/*запрос на существование чата по id
	found, chat, err := s.chatApiDB.CheckExistingChatDB(req.UID, receiverUser.ID)
	*/
	if !found {
		/*запрос на создание чата по id
		chat, err = s.chatApiDB.CreateNewChatDB(req.UID, receiverUser.ID)
		if err != nil {
			return nil, fmt.Errorf("create new chat DB: %w", err)
		}
		*/
	}
	/*запрос на получение создание сообщения чата по id
	message, err := s.messageApiDB.CreateNewMessageDB(chat.ID, req.UID, req.Message, "string")
	if err != nil {
		return nil, fmt.Errorf("create new message DB: %w", err)
	}
	*/

	/*запрос на получение последнего сообщения чата по id
	if err = s.chatApiDB.AddLastMessageToChatDB(message.Message, chat.ID, message.Time); err != nil {
		return nil, fmt.Errorf("add last message to chat DB: %w", err)
	}
	*/

	return &model.SendMessageHandlerRes{MessageID: message.ID, RequestedUserID: req.UID}, nil
}

func (s *ChatServiceImpl) SendFileService(req model.SendFilesHandlerReq) (*model.SendFilesHandlerRes, error) {
	/*запрос на получение пользователя по юзернейму
		receiverUser, err := s.userApiDB.GetUserByUsernameDB(req.ReceiverUsername)
	if err != nil {
		return nil, fmt.Errorf("get user by username DB: %w", err)
	}
	*/

	/*запрос на существование чата по id
	found, chat, err := s.chatApiDB.CheckExistingChatDB(req.UID, receiverUser.ID)
	*/

	var messages []int64
	if !found {
		/*запрос на создание чата по id
		chat, err = s.chatApiDB.CreateNewChatDB(req.UID, receiverUser.ID)
		if err != nil {
			return nil, fmt.Errorf("create new chat DB: %w", err)
		}
		*/
	}

	paths, err := files.DownloadFiles(chat.ID, req.Files)
	if err != nil {
		return nil, fmt.Errorf("download files: %w", err)
	}

	for key, elem := range paths {
		/*запрос на создание сообщения чата
		message, err := s.messageApiDB.CreateNewMessageDB(chat.ID, req.UID, strings.ReplaceAll(key, "src/chat/", "https://center.web-gen.ru:444/"), elem)
		if err != nil {
			return nil, fmt.Errorf("create new message DB: %w", err)
		}
		*/

		/*запрос на обновление последнего сообщения чата по id
		if err = s.chatApiDB.AddLastMessageToChatDB(message.Message, chat.ID, message.Time); err != nil {
			return nil, fmt.Errorf("add last message to chat DB: %w", err)
		}

		*/

		messages = append(messages, message.ID)
	}

	return &model.SendFilesHandlerRes{MessagesID: messages, RequestedUserID: req.UID}, nil
}

func (s *ChatServiceImpl) SetOnlineService(req model.SetOnlineHandlerReq) error {
	/*запрос на установление онлайна по id
		err := s.activeApiDB.SetOnlineDB(req.UID)
	if err != nil {
		return fmt.Errorf("set online DB: %w", err)
	}
	*/
	return nil
}

func (s *ChatServiceImpl) FullTextMessageSearchService(req model.FullTextMessageSearchHandlerReq) (*model.FullTextMessageSearchHandlerRes, error) {
	chatIDint, err := strconv.Atoi(req.ChatID)
	if err != nil {
		return nil, fmt.Errorf(fmt.Errorf("chat id atoi: %w", err).Error()+": %w", custom_errors.ErrWrongInputData)
	}

	/*запрос на поиск по сообщениям
	messages, err := s.messageApiDB.FullTextMessageSearchDB(req.Search, int64(chatIDint))
	if err != nil {
		return nil, fmt.Errorf("full text message search DB: %w", err)
	}
	*/

	return &model.FullTextMessageSearchHandlerRes{Messages: messages}, nil
}
