package users

import (
	"github.com/google/uuid"
)

type User struct {
	id   string
	name string

	friendRequests []string
	friends        []string
	chatIDs        map[string]string
}

func NewUser(name string) User {
	return User{
		id:      uuid.New().String(),
		name:    name,
		chatIDs: map[string]string{},
	}
}
func (u *User) GetID() string {
	return u.id
}
func (u *User) getName() string {
	return u.name
}

func (u *User) setName(newName string) {
	u.name = newName
}

func (u *User) addFriendRequest(userID string) {
	u.friendRequests = append(u.friendRequests, userID)
}

func (u *User) approveFriendRequest(userID string) string {
	for i, ur := range u.friendRequests {
		if ur == userID {
			u.friendRequests = append(u.friendRequests[:i], u.friendRequests[i+1:]...)
			return uuid.New().String()
		}
	}
	return ""
}

func (u *User) deleteFriendRequest(userID string) {
	for i, ur := range u.friendRequests {
		if ur == userID {
			u.friendRequests = append(u.friendRequests[:i], u.friendRequests[i+1:]...)
		}
	}
}
func (u *User) getFriends() []string {
	return u.friends
}

func (u *User) addFriend(userID, chatID string) {

	u.setChatID(chatID, userID)
	u.friends = append(u.friends, userID)
}

func (u *User) getChatID(reciverID string) string {

	return u.chatIDs[reciverID]

}
func (u *User) setChatID(chatID, reciverID string) {
	u.chatIDs[reciverID] = chatID
}
