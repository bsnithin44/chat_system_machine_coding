package msgs

type IMessageService interface {
	AddMessage(chatId, sender string, msg Message)
	GetMessages(chatId string) []Message
}

type messageService struct {
	dao map[string][]Message
}

func NewMessageService() IMessageService {
	return messageService{
		dao: make(map[string][]Message),
	}
}

func (ms messageService) AddMessage(chatId, sender string, msg Message) {

	ms.dao[chatId] = append(ms.dao[chatId], msg)

}

func (ms messageService) GetMessages(chatId string) []Message {

	if _, ok := ms.dao[chatId]; ok {

		return ms.dao[chatId]
	}

	return nil
}
