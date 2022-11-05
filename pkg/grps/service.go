package grps

import "chatss/pkg/msgs"

type IGroupService interface {
	CreateGroup(grpName, adminID string) string
	AddUserToGroup(userID, grpID string)
	SendGroupMessage(grpID, senderID string, msg msgs.Message)
	GetGroupMessages(grpID string) []msgs.Message
}

type groupService struct {
	dao            map[string]Group
	messageService msgs.IMessageService
}

func NewGroupService(mS msgs.IMessageService) IGroupService {
	return &groupService{
		dao:            map[string]Group{},
		messageService: mS,
	}
}

func (gs *groupService) CreateGroup(groupName, adminID string) string {

	grp := NewGroup(groupName, adminID)
	grp.setChatID()
	gs.dao[grp.id] = grp
	return grp.id

}
func (gs *groupService) AddUserToGroup(userID, grpID string) {

	grp, ok := gs.dao[grpID]

	if ok {
		grp.addUser(userID)
	}
	gs.dao[grpID] = grp

}

func (gs *groupService) SendGroupMessage(grpID, senderID string, msg msgs.Message) {
	grp, ok := gs.dao[grpID]

	if ok {

		for _, id := range grp.Users {
			if id == senderID {
				gs.messageService.AddMessage(grp.getChatID(), senderID, msg)
			}

		}
	}
}
func (gs *groupService) GetGroupMessages(grpID string) []msgs.Message {
	grp, ok := gs.dao[grpID]

	if ok {
		return gs.messageService.GetMessages(grp.getChatID())
	}
	return nil
}
