package users

import (
	"chatss/pkg/msgs"
)

type IUserService interface {
	CreateUser(name string) string
	UpdateUser(id, name string)
	DeleteUser(id string)
	SendFriendRequest(user1ID, user2ID string)
	ApproveFriendReq(user1ID, user2ID string)
	RejectFriendReq(user1ID, user2ID string)
	Send1to1Message(userID1, userID2 string, msg msgs.Message)
	Get1to1Messages(userID1, userID2 string) []msgs.Message
}

type userService struct {
	dao        map[string]User
	msgService msgs.IMessageService
}

func NewUserService(msgservice msgs.IMessageService) IUserService {
	return &userService{
		dao:        map[string]User{},
		msgService: msgservice,
	}
}

func (us *userService) CreateUser(name string) string {

	usr := NewUser(name)
	us.dao[usr.id] = usr

	return usr.id
}

func (us *userService) UpdateUser(id, name string) {

	user := us.dao[id]
	user.setName(name)
	us.dao[id] = user

}

func (us *userService) DeleteUser(id string) {

	delete(us.dao, id)
}

func (us *userService) SendFriendRequest(userID1, userID2 string) {

	user1, ok1 := us.dao[userID1]
	user2, ok2 := us.dao[userID2]

	if ok1 && ok2 {
		user2.addFriendRequest(user1.id)

	}
	us.dao[userID2] = user2
}

func (us *userService) ApproveFriendReq(userID1, userID2 string) {
	user1, ok1 := us.dao[userID1]
	user2, ok2 := us.dao[userID2]

	if ok1 && ok2 {

		chatID := user1.approveFriendRequest(user2.id)
		user1.addFriend(user2.id, chatID)
		user2.addFriend(user1.id, chatID)
	}
	us.dao[userID1] = user1
	us.dao[userID2] = user2

}

func (us *userService) RejectFriendReq(userID1, userID2 string) {
	user1, ok1 := us.dao[userID1]
	user2, ok2 := us.dao[userID2]

	if ok1 && ok2 {

		user1.deleteFriendRequest(user2.id)
	}
	us.dao[userID1] = user1

}

func (us *userService) Send1to1Message(userID1, userID2 string, msg msgs.Message) {

	user1, ok1 := us.dao[userID1]
	user2, ok2 := us.dao[userID2]

	chatID := user1.getChatID(user2.id)

	if ok1 && ok2 && chatID != "" {

		us.msgService.AddMessage(chatID, user1.id, msg)
	}
}

func (us *userService) Get1to1Messages(userID1, userID2 string) []msgs.Message {

	user1, ok1 := us.dao[userID1]
	user2, ok2 := us.dao[userID2]

	chatID := user1.getChatID(user2.id)

	if ok1 && ok2 {

		return us.msgService.GetMessages(chatID)
	}
	return nil
}
